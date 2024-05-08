package main

import (
	"fmt"
	"strings"

	"github.com/ArrayOfLilly/pokedexcli/internal/pokeapi"
)

func callbackExplore(cfg *config, param string) error {
	pokeapiClient := pokeapi.NewClient()
  	resp, err := pokeapiClient.ListPokemonsPerAreas(param)
	if err != nil {
		return err
	}

	
	location := strings.Title(strings.ReplaceAll(param, "-", " "))

	fmt.Println("-----------------------------------")
	fmt.Printf("Pokemon Encounters at %s: \n", location)
	fmt.Println("-----------------------------------")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
