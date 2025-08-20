package main

import (
	"errors"
	"fmt"
	"github.com/Tenlinks20/pokedex_cli/internal/pokeapi"
)


func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return  errors.New("require 1 name only")
	}
	name := args[0]
	info, exists := cfg.CaughtPokemon[name]

	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	if name != info.Name {
		return fmt.Errorf("bad data: user put: %s, found: %s", name, info.Name)
	}

	fmt.Printf("Name: %s\nWeight: %d\nHeight: %d\n", name,info.Weight, info.Height)
	fmt.Printf("Stats:\n")

	for key, val := range pokeapi.GetStats(info) {
		fmt.Printf("    -%s: %d\n", key, val)
	}

	fmt.Printf("Types:\n")
	for _, t := range pokeapi.GetTypes(info) {
		fmt.Printf("    -%s\n", t)
	}

	return nil
}