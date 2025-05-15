package main

import (
	"fmt"
	"os"
	"sort"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commRegistry = make(map[string]cliCommand)

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	// Get the keys
	keys := make([]string, 0, len(commRegistry))
	for k := range commRegistry {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %s\n", k, commRegistry[k].description)
	}

	return nil
}
