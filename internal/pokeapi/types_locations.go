package pokeapi

type LocationArea struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemons struct {
	PokemonEncounters []Pokemon `json:"pokemon_encounters"`
}

type Pokemon struct {
	PokemonObj PokemonDetails `json:"pokemon"`
}

type PokemonDetails struct {
	Name string `json:"name"`
}

type CaughtPokemon struct {
	Name           string `json:"name"`
	BaseExperience int64  `json:"base_experience"`
}
