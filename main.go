package main

import (
	"time"
	"github.com/Tenlinks20/pokedex_cli/internal/pokeapi"
	"github.com/Tenlinks20/pokedex_cli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(20 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache: pokeCache,
	
	}

	interactiveMode(cfg)
}
