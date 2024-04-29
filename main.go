package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Error: too many arguments")
		fmt.Println("Usage: pokedex")
		return
	}

	// greeting
	fmt.Printf("\nWelcome to the 🐙 Pokedex!\n\n")

	startREPL()
}
