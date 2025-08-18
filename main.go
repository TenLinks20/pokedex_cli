package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		words := cleanInput(userInput)

		if len(words) == 0 {
			fmt.Println("No input provided")
		} else {
			fmt.Printf("Your command was: %s\n", words[0])
		}

	}
}