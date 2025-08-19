package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"github.com/Tenlinks20/pokedex_cli/internal/pokeapi"
)

// Configuration
type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

// Command registry
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

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
		commands := getCommands()

		command, ok := commands[userCommand]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		}
		
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}