package main

import (
	"fmt"
	"os"

	"github.com/ArrayOfLilly/pokedexcli/internal/pokeapi"
)

type config struct{
	pokeapiClient pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Error: too many arguments")
		fmt.Println("Usage: pokedex")
		return
	}

	// config
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	// greeting
	fmt.Printf("\nWelcome to the 🐙 Pokedex!\n\n")
	
	// REPL
	startREPL(&cfg)
}
