package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	displayHelp := "Welcome to the Pokedex!\n"
	commands := getCommands()
	for _, command := range commands {
		displayHelp += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
	fmt.Println(displayHelp)
	return nil
}