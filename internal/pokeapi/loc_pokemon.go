package pokeapi

import (
	"fmt"
	"encoding/json"
	"errors"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	URL string  `json:"url"`
}

type PokemonEncounter struct {
    Pokemon Pokemon `json:"pokemon"`
    // other fields if needed
}

type LocPokemon struct {
    Name string `json:"name"`
    PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func (c *Client) GetLocPokemonFromReq(loc string) (LocPokemon, error) {

	if loc == "" {
		return LocPokemon{}, errors.New("no loc given")
	}
	url := baseURL + "/location-area/" + loc

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return  LocPokemon{}, err
	}

	var area LocPokemon
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&area); err != nil {
		return LocPokemon{}, nil
	}

	return area, nil
}

func PrintPokemonNames(encs []PokemonEncounter) {
	var names []string
	for _, enc := range encs {
		names = append(names, enc.Pokemon.Name)
	}
	for _, name := range names {
		fmt.Printf("- %s\n", name)
	}
}