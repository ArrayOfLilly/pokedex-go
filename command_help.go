package main

import (
	"fmt"
)

func callbackHelp(cfg *config, pkdx *map[string]Pokemon, param string) error {
	fmt.Println("-----------------------------------")
	fmt.Println("Welcome to the Pokedex help menu:")
	fmt.Println("-----------------------------------")
	
	commands := getCommands()
	cmd, ok := commands[param]
	fmt.Println(param)
	if ok {
		fmt.Printf("* %s: %s\n", cmd.name, cmd.description)
		return nil
	}

	for _, item := range commands {
		if item.name == "ash" {
			continue
		}
		fmt.Printf("* %s: %s\n", item.name, item.description)
	}
	fmt.Println("")
	
	return nil
}



