package main

import (
	"fmt"
)

func callbackHelp(cfg *config) error {
	commands := getCommands()
	fmt.Println("-----------------------------------")
	fmt.Println("Welcome to the Pokedex help menu:")
	fmt.Println("-----------------------------------")
	for _, item := range commands {
		if item.name == "ash" {
			continue
		}
		fmt.Printf("* %s: %s\n", item.name, item.description)
	}
	fmt.Println("")
	
	return nil
}



