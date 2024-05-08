package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func callbackCatch(cfg *config, param string) error {
	resp, err := cfg.pokeapiClient.DescribePokemon(param)
	if err != nil {
		return err
	}

	resp2, err := cfg.pokeapiClient.DescribePokemonSpecies(fmt.Sprintf("%d", resp.ID))
	if err != nil {
		return err
	}

	catchRate := resp2.CaptureRate
	shinyRate := 8192 / 2  // Shiny Charm
	

	fmt.Printf("Throwing a Pokeball at %s...", strings.Title(param))
	catch := rand.Intn(catchRate + 1)
	fmt.Println(catch)
	
	if catch == 0 {
		fmt.Printf("%s was caught!\n", strings.Title(param))
		isShiny := rand.Intn(shinyRate + 1)
		if isShiny == 0 {
			fmt.Printf("Your %s is shiny!\n", strings.Title(param))
			
		}
	} else {
		fmt.Printf("%s was escaped!\n", strings.Title(param))
	}
	return nil
}

