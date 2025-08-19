package main

import (
	"time"
	"github.com/Tenlinks20/pokedex_cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	interactiveMode(cfg)
}
