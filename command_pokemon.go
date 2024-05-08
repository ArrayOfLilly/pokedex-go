package main

import (
	"fmt"
	"strings"
)

func callbackPokemon(cfg *config, param string) error {
  	resp, err := cfg.pokeapiClient.DescribePokemon(param)
	if err != nil {
		return err
	}

	resp2, err := cfg.pokeapiClient.DescribePokemonSpecies(fmt.Sprintf("%d", resp.ID))
	if err != nil {
		return err
	}

	resp3, err := cfg.pokeapiClient.GetEvolutionChain(resp2.EvolutionChain.URL)
	if err != nil {
		return err
	}


	pokemon := strings.Title(strings.ReplaceAll(param, "-", " "))

	fmt.Println("--------------------------------------")
	fmt.Printf("All about %s: \n", pokemon)
	fmt.Println("--------------------------------------")
	fmt.Printf("Species: %s\n", strings.Title(resp.Species.Name))
	fmt.Printf("National Pokedex ID: %d\n", resp2.PokedexNumbers[0].EntryNumber)
	
	// Type
	fmt.Printf("Types: \n")
	for _, t := range resp.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	
	// Stats
	fmt.Printf("Base Stats: \n")
	for _, s := range resp.Stats {
		fmt.Printf("- %s: %d\n", s.Stat.Name, s.BaseStat)
	}

	// Evolutin Chain
	fmt.Printf("Evolution Chain: \n")
	if len(resp3.Chain.EvolvesTo) == 0 {
		fmt.Printf("%s does not evolve\n", strings.Title(resp3.Chain.Species.Name))
	} else {
		for _, e := range resp3.Chain.EvolvesTo {
			fmt.Printf("%s evolves to %s\n", strings.Title(resp3.Chain.Species.Name), strings.Title(e.Species.Name))
			if len(e.EvolvesTo) == 0 {
				fmt.Printf("%s does not evolve\n", strings.Title(e.Species.Name))
			} else {
				for _, ee := range e.EvolvesTo {
					fmt.Printf("%s evolves to %s\n", strings.Title(e.Species.Name), strings.Title(ee.Species.Name))
				}	
			}
		}
	}

	fmt.Printf("Egg Groups: \n")
	for _, t := range resp2.EggGroups {
		 fmt.Printf("- %s \n", t.Name)
	}

	if resp2.IsLegendary {
		fmt.Println("Legendary Pokemon")
	}
	if resp2.IsMythical {
		fmt.Println("Mythical Pokemon")
	}
	
	fmt.Println("")
	return nil
}