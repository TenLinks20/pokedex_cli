package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

// Command registry
type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
	
}

func cleanInput(text string) []string {
    lowerText := strings.ToLower(text)
    return strings.Fields(lowerText)
}

func interactiveMode() {
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
		
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}