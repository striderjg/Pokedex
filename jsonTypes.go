package main

type locAreaJSON struct {
	Name string
	Url  string
}

type mapJSON struct {
	Next     string
	Previous string
	Results  []locAreaJSON
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type PokemonList struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}
