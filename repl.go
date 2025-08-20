package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"github.com/Tenlinks20/pokedex_cli/internal/pokeapi"
	"github.com/Tenlinks20/pokedex_cli/internal/pokecache"
)

// Configuration
type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokeCache   *pokecache.Cache
	CaughtPokemon map[string]pokeapi.PokemonInfo
}

// Command registry
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

// 

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help", 
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "(map back) Displays the names of  the previous 20 location areas in the Pokemon world",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "(map back) Displays the names of  the previous 20 location areas in the Pokemon world",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "(map back) Displays the names of  the previous 20 location areas in the Pokemon world",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "(map back) Displays the names of  the previous 20 location areas in the Pokemon world",
			callback: commandInspect,
		},
	}
	
}

func cleanInput(text string) []string {
    lowerText := strings.ToLower(text)
    return strings.Fields(lowerText)
}

func interactiveMode(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := scanner.Text()
		if userInput == "" {
			continue
		}

		words := cleanInput(userInput)
		if len(words) == 0 {
			continue
		}

		userCommand := words[0]
		command, ok := getCommands()[userCommand]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		}

		if len(words) == 1 {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else if args := words[1:]; len(args) > 0 {		
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} 
	}
}