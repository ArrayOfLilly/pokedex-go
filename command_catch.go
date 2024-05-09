package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func callbackCatch(cfg *config, pkdx *map[string]Pokemon, param string) error {
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

	catchRate := resp2.CaptureRate / 50    // real catch rate unbearable in CLI 
	shinyRate := 8192 / 2000               // huge shiny charm
	rand.Seed(time.Now().UnixMilli())

	fmt.Printf("Throwing a Pokeball at %s...", strings.Title(param))
	catch := rand.Intn(catchRate + 1)   
	fmt.Printf(" 1, 2, 3, ... %d\n", catch)
	
	if catch != 0 {
		fmt.Printf("🐾 %s was escaped! \n", strings.Title(param))
		return nil
	} 
	
	fmt.Printf("🐭 %s was caught!\n", strings.Title(param))
	isShiny := rand.Intn(shinyRate + 1)
	if isShiny == 0 {
		fmt.Printf("Your %s is shiny! ✨\n", strings.Title(param))
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Add nickname your new pokemon (spaces are not allowed): \n")
	scanner.Scan()
	text := scanner.Text()
	cleaned := cleanupInput(text)
	nickname := ""
	if len(cleaned) <= 0 {
		nickname = strings.Title(param) + fmt.Sprintf("%d", rand.Intn(10000))
	} else {
		nickname = cleaned[0]
		nickname = strings.Title(nickname)
	}
	
	var stats []int
	for _, bs := range resp.Stats {
		stats = append(stats, bs.BaseStat)
	}

	var moves []string
	for _, mov := range resp.Moves {
		moves = append(moves, mov.Move.Name)
	}

	var abilities []string
	for _, ability := range resp.Abilities {
		abilities = append(abilities, ability.Ability.Name)
	}

	var types []string
	for _, t := range resp.Types {
		types = append(types, t.Type.Name)
	} 

	var eggs []string
	for _, egg := range resp2.EggGroups {
		eggs = append(eggs, egg.Name)
	}

	var chain [][]string
	var first []string
	var second []string
	var third []string 
	first = append(first, resp3.Chain.Species.Name)
	if len(resp3.Chain.EvolvesTo) != 0 {
		for _, e := range resp3.Chain.EvolvesTo {
			second = append(second, e.Species.Name)
			if len(e.EvolvesTo) != 0 {
				for _, ee := range e.EvolvesTo {
					third = append(third, ee.Species.Name)
				}	
			}
		}
	}
	chain = append(chain, first, second, third)
	
	pokedex := *pkdx
	_, ok := pokedex[param]
	if !ok {
		pokedex[param] = Pokemon{
			species: resp.Species.Name,
			isLegendary: resp2.IsLegendary,
			isMythical: resp2.IsMythical,
			isBaby: resp2.IsBaby,
			height: resp.Height,
			weight: resp.Weight,
			hp: stats[0],
			attack: stats[1],
			defense: stats[2],
			specialAttack: stats[3],
			specialDefense: stats[4],
			speed: stats[5],
			moves: moves,
			abilities: abilities,
			types: types, 
			eggGroups: eggs,
			evolutionChain: chain,
			caught: []ownedPokemon{
				{
				nickname: nickname,
				isShinny: (isShiny == 0),
				height: int(float32(resp.Height) * (rand.Float32()/10+1)),
				weight: int(float32(resp.Weight) * (rand.Float32()/10+1)),
				hp: stats[0] + rand.Intn(32),
				attack: stats[1] + rand.Intn(32),
				defense: stats[2] + rand.Intn(32),
				specialAttack: stats[3] + rand.Intn(32),
				specialDefense: stats[4] + rand.Intn(32),
				speed: stats[5] + rand.Intn(32),
				},
			},
		}
	
	} else {
		newPokemon := ownedPokemon{
				nickname: nickname,
				isShinny: (isShiny == 0),
				height: int(float32(resp.Height) * (rand.Float32()/10+1)),
				weight: int(float32(resp.Weight) * (rand.Float32()/10+1)),
				hp: stats[0] + rand.Intn(32),
				attack: stats[1] + rand.Intn(32),
				defense: stats[2] + rand.Intn(32),
				specialAttack: stats[3] + rand.Intn(32),
				specialDefense: stats[4] + rand.Intn(32),
				speed: stats[5] + rand.Intn(32),
				}

		pokemon := pokedex[param]
		pokemon.caught = append(pokemon.caught, newPokemon)
		pokedex[param] = pokemon		
		
	}
	return nil
}