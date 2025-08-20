package main

import (
	"errors"
	"fmt"
    "encoding/json"
	"github.com/Tenlinks20/pokedex_cli/internal/pokeapi"
)

func commandMap(cfg *config, args ...string) error {
    keyPtr := cfg.nextLocationsURL
    if keyPtr == nil {
        url := "https://pokeapi.co/api/v2/location-area"
        keyPtr = &url
    }
    entry, exists := cfg.pokeCache.Get(*keyPtr)

    if exists {
        var page pokeapi.Page
        err := json.Unmarshal(entry, &page)
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

    page, err := cfg.pokeapiClient.GetPageFromReq(keyPtr)
    if err != nil {
        return err
    }

    jsonData, err := json.Marshal(page)
    if err != nil {
        return  err
    }
    cfg.pokeCache.Add(*keyPtr, jsonData)

    locations := pokeapi.GetLocations(page)
    for _, location := range locations {
        fmt.Println(location)
    }
    
    // Update the URLs for next/previous
    cfg.nextLocationsURL = page.Next
    cfg.prevLocationsURL = page.Previous
    
    return nil
}

func commandMapb(cfg *config, args ...string) error {
    keyPtr := cfg.prevLocationsURL
	if keyPtr == nil {
		return errors.New("you're on the first page")
	}

    entry, exists := cfg.pokeCache.Get(*keyPtr)
    
    if exists {
        var page pokeapi.Page
        err := json.Unmarshal(entry, &page)
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

    page, err := cfg.pokeapiClient.GetPageFromReq(keyPtr)
    if err != nil {
        return err
    }

    jsonData, err := json.Marshal(page)
    if err != nil {
        return  err
    }
    cfg.pokeCache.Add(*keyPtr, jsonData)
    
    locations := pokeapi.GetLocations(page)
    for _, location := range locations {
        fmt.Println(location)
    }
    
    // Update the URLs for next/previous
    cfg.nextLocationsURL = page.Next
    cfg.prevLocationsURL = page.Previous
    
    return nil
}