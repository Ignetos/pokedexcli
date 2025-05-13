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
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text = scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		firstWord := words[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}
