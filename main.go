package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	fmt.Println("Hello, World!")
	cleanInput("text")
}
