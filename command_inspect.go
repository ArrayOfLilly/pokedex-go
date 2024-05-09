package main

import (
	"fmt"
	"strings"
)

// TODO: thorough description about the species and/or the specified Pokemon, later
func callbackInspect(cfg *config, pkdx *map[string]Pokemon, param string) error {
	pokedex := *pkdx
	if len(param) == 0 {
		fmt.Printf("\nYou registered:\n")
		for k, v := range pokedex {
			fmt.Printf("- %d %s\n", len(v.caught), k)
		}
	} else {
		_, ok := pokedex[param]
		if !ok {
			fmt.Printf("You haven't caught %s yet.\n", strings.Title(param))
			return nil
		}
		// pkmn := pokedex[param]
		fmt.Printf("You have caught %d %s so far\n", len(pokedex[param].caught), strings.Title(param))
		fmt.Printf("your buddy %ss are: \n", pokedex[param].species)
		for _, pokemon := range pokedex[param].caught {
			shine := ""
			if pokemon.isShinny {
				shine = "✨"
			}
			fmt.Printf("- %s%s, it has %d lbs weight and %d feet height\n", pokemon.nickname, shine, pokemon.weight, pokemon.height)
		}
	}
	fmt.Println("")
	return nil
}