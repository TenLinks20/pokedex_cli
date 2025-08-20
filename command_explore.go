package main

import (
	"fmt"
	"errors"
	"encoding/json"
	"github.com/Tenlinks20/pokedex_cli/internal/pokeapi"
	
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location name")
	}
	
	if len(args) > 1 {
		return errors.New("you must provide exactly one location name")
	}

	loc := args[0]
	fmt.Printf("Exploring %s...\n", loc)

	entry, exists := cfg.pokeCache.Get(loc)
	if exists {
		var encounters []pokeapi.PokemonEncounter
		if err := json.Unmarshal(entry, &encounters); err != nil {
			return err
		}

		fmt.Println("Found Pokemon:")
		pokeapi.PrintPokemonNames(encounters)
		return nil
	}

	area, err := cfg.pokeapiClient.GetLocPokemonFromReq(loc)
	if err != nil {
		return err
	} else if area.Name != loc {
		return fmt.Errorf("bad data found:\nuser put:%s\nloc found: %s", loc, area.Name)
	}

	jsonData, err := json.Marshal(area.PokemonEncounters)
	if err != nil {
		return err
	} 
	cfg.pokeCache.Add(loc, jsonData)
	
	fmt.Println("Found Pokemon:")
	pokeapi.PrintPokemonNames(area.PokemonEncounters)
	return nil
}
	