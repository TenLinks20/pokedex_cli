package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokemon:")
	for name := range cfg.CaughtPokemon {
		fmt.Printf("- %s\n", name)
	} 
	return nil
}