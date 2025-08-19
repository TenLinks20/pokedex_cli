package main

import (
	"fmt"
	"errors"
	"github.com/Tenlinks20/pokedex_cli/internal/pokeapi"
)

func commandMap(cfg *config) error {
    page, err := cfg.pokeapiClient.GetPageFromReq(cfg.nextLocationsURL)
    if err != nil {
        return err
    }
    
    locations := pokeapi.GetLocations(page)
    for _, location := range locations {
        fmt.Println(location)
    }
    
    // Update the URLs for next/previous
    cfg.nextLocationsURL = page.Next
    cfg.prevLocationsURL = page.Previous
    
    return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	page, err := cfg.pokeapiClient.GetPageFromReq(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = page.Next
	cfg.prevLocationsURL = page.Previous

	locations := pokeapi.GetLocations(page)
    for _, location := range locations {
        fmt.Println(location)
    }
	
	return nil
}