package main

type ownedPokemon struct {
	nickname string
	isShinny bool
	height int
	weight int
	hp int
	attack int
	defense int
	specialAttack int
	specialDefense int
	speed int
}

type Pokemon struct {
	species string
	isLegendary bool
	isMythical bool
	isBaby bool
	height int
	weight int
	hp int
	attack int
	defense int
	specialAttack int
	specialDefense int
	speed int
	moves []string
	abilities []string
	types []string
	eggGroups []string
	evolutionChain [][]string
	caught []ownedPokemon
}

func getPokedex() map[string]Pokemon {
	return map[string]Pokemon{}
}