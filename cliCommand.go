package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/Ignetos/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	next     string
	previous string
	param    string
	cache    *internal.Cache
}

var commRegistry = make(map[string]cliCommand)

func printMap(md internal.MapData) {
	data := md.Results
	for _, area := range data {
		fmt.Println(area.Name)
	}
}

func printExplore(ed internal.ExploreData, area string) {
	data := ed.PokemonEncounters
	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, encounter := range data {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
}

func commandExit(c *Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
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

func commandMap(c *Config) error {
	url := c.next
	mapData, err := internal.GetData[internal.MapData](url, c.cache)
	if err != nil {
		return fmt.Errorf("unable to get map data: %w", err)
	}

	c.next = mapData.Next
	c.previous = mapData.Previous

	printMap(mapData)

	return nil
}

func commandMapB(c *Config) error {
	url := c.previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	mapData, err := internal.GetData[internal.MapData](url, c.cache)
	if err != nil {
		return fmt.Errorf("unable to get map data: %w", err)
	}

	c.next = mapData.Next
	c.previous = mapData.Previous

	printMap(mapData)

	return nil
}

func commandExplore(c *Config) error {
	url := internal.BASEURL + "location-area/" + c.param
	exploreData, err := internal.GetData[internal.ExploreData](url, c.cache)
	if err != nil {
		return fmt.Errorf("unable to get map data: %w", err)
	}
	if exploreData.Location.Name == "" {
		fmt.Println("location not found")
		return nil
	}

	printExplore(exploreData, c.param)

	return nil
}
