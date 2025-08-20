package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("one pokemon name required only")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	entry, exists := cfg.CaughtPokemon[name]
	if exists {

		catchChance := float64(entry.BaseExperience) / 600.0
		randFloat := rand.Float64()  

		if randFloat > catchChance {
			fmt.Printf("%s was caught again!\n", name)
		} else {
			fmt.Printf("%s escaped!\n", name)
		}
		return nil
	}

	info, err := cfg.pokeapiClient.GetPokemonInfoFromReq(name)
	if err != nil {
		return err
	}
	catchChance := float64(info.BaseExperience) / 400.0
	randFloat := rand.Float64()  

	if randFloat > catchChance {
		fmt.Printf("%s was caught!\n", name)
		cfg.CaughtPokemon[name] = info
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}