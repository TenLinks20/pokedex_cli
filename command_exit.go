package main

import (
	"os"
	"fmt"
)

func commandExit(cfg *config) error {
	if _, err := fmt.Println("Closing the Pokedex... Goodbye!"); err != nil {
		return err
	}
	os.Exit(0)
	return nil
}