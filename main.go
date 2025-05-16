package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	return strings.Fields(input)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var text string

	config := Config{
		next:     "https://pokeapi.co/api/v2/location-area",
		previous: "",
	}

	commRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays location area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous location area",
			callback:    commandMapB,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text = scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		firstWord := words[0]

		if _, exist := commRegistry[firstWord]; exist {
			err := commRegistry[firstWord].callback(&config)
			if err != nil {
				fmt.Println("Unable to execute command")
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
