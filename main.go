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
			commRegistry[firstWord].callback()
		} else {
			fmt.Println("Unknown command")
		}

	}
}
