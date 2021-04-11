package storage

type Ability struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IsMainSeries bool   `json:"is_main_series"`
	Slot         int    `json:"slot"`
	Ref          struct {
		Url string `json:"url"`
	} `json:"ability"`
}

type Pokemon struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	BaseExperience int       `json:"base_experience"`
	Height         int       `json:"height"`
	Weight         int       `json:"weight"`
	IsDefault      bool      `json:"is_default"`
	Order          int64     `json:"order"`
	Abilities      []Ability `json:"abilities"`
	Url            string    `json:"url"`
}

type Repository interface {
	GetPokemons() ([]Pokemon, error)
}
