package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/caleeli/go_concurrent/errors"
	"github.com/caleeli/go_concurrent/storage"
)

const (
	pokemonsUrl = "https://pokeapi.co/api/v2/pokemon"
)

type repository struct {
	url string
}

type apiResponse struct {
	Results []storage.Pokemon `json:"results"`
}

func NewRepository() storage.Repository {
	return &repository{url: pokemonsUrl}
}

func (repo *repository) GetPokemons() (pokemons []storage.Pokemon, err error) {

	completed := make(chan []storage.Pokemon)
	failed := make(chan error)

	go repo.getPokemons(completed, failed)

	select {
	case pokemons = <-completed:
		return
	case err = <-failed:
		fmt.Println("Error:", err)
		return nil, err
	}
}

func getFromUrl(url string) (contents []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.WrapCanNotLoadList(err, "Error loading data from API %s", url)
	}
	contents, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapCanNotLoadList(err, "Error reading data from API %s", url)
	}
	return
}

func (repo *repository) getPokemons(completed chan []storage.Pokemon, failed chan error) {
	var response apiResponse
	waitGroup := sync.WaitGroup{}
	contents, err := getFromUrl(repo.url)
	if err != nil {
		failed <- err
	}
	err = json.Unmarshal(contents, &response)
	if err != nil {
		failed <- errors.WrapCanNotLoadList(err, "can't parsing response into pokemons")
	}

	// Load pokemon info
	lenght := len(response.Results)
	waitGroup.Add(lenght)
	for i := 0; i < lenght; i++ {
		go loadPokemon(&response.Results[i], &waitGroup, failed)
	}
	waitGroup.Wait()
	completed <- response.Results
}

func loadPokemon(pokemon *storage.Pokemon, waitGroup *sync.WaitGroup, failed chan error) {
	contents, err := getFromUrl(pokemon.Url)
	if err != nil {
		failed <- err
		return
	}
	err = json.Unmarshal(contents, pokemon)
	if err != nil {
		failed <- errors.WrapCanNotLoadList(err, "can't parsing response into %s \n %s", pokemon.Url, string(contents))
		return
	}

	// Load abilities
	waitAbilities := sync.WaitGroup{}
	lenght := len(pokemon.Abilities)
	waitAbilities.Add(lenght)
	for i := 0; i < lenght; i++ {
		go loadAbility(&pokemon.Abilities[i], &waitAbilities, failed)
	}
	waitAbilities.Wait()

	waitGroup.Done()
}

func loadAbility(ability *storage.Ability, waitGroup *sync.WaitGroup, failed chan error) {
	contents, err := getFromUrl(ability.Ref.Url)
	if err != nil {
		failed <- err
		return
	}
	err = json.Unmarshal(contents, ability)
	if err != nil {
		failed <- errors.WrapCanNotLoadList(err, "can't parsing response into %s \n %s", ability.Ref.Url, string(contents))
		return
	}
	waitGroup.Done()
}
