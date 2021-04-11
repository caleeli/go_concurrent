package cli

import (
	"fmt"

	"github.com/caleeli/go_concurrent/storage"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

// InitPokemonCmd is a command to load pokemons info
func InitPokemonCmd(repository storage.Repository) *cobra.Command {
	command := &cobra.Command{
		Use:   "pokemons",
		Short: "Print data about pokemons",
		Run:   runPokemonsFn(repository),
	}

	return command
}

func runPokemonsFn(repository storage.Repository) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		pokemons, _ := repository.GetPokemons()
		fmt.Println(pokemons)
	}
}
