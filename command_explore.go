package main

import (
	"fmt"
	"strings"
)

func callbackExplore(cfg *config, pkdx *map[string]Pokemon, param string) error {
  	resp, err := cfg.pokeapiClient.ListPokemonsPerAreas(param)
	if err != nil {
		return err
	}

	location := strings.Title(strings.ReplaceAll(param, "-", " "))

	fmt.Println("---------------------------------------------")
	fmt.Printf("Pokemon Encounters at %s: \n", location)
	fmt.Println("---------------------------------------------")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	fmt.Println("")
	return nil
}
