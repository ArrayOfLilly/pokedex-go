package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL(cfg *config) {

	// read stin
	scanner := bufio.NewScanner(os.Stdin)
	
	// REPL running
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		
		cleaned := cleanupInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		
		availableCommands := getCommands()
		
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue		
		}

		if len(cleaned) > 1 {
			param := cleaned[1]
			err := command.callback(cfg, param)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err:= command.callback(cfg, "")
			if err != nil {
				fmt.Println(err)
			}
		}		
	}
}


type cliCommand struct {
	name string
	description string
	callback func(*config, string) error
}
	

// list all available commands (help)
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message\n  usage: help or help [command name]",
			callback:    callbackHelp,
		},
		"clear": {
			name:        "clear",
			description: "Clear screen",
			callback:    callbackClear,
		}, 
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		}, 
		"ash": {
			name:        "ash",
			description: "Say something kind.",
			callback:    callbackAsh,
		},
		"map": {
			name:        "map",
			description: "Paginated list of location areas, 20 per page",
			callback:    callbackMap,
		}, 
		"mapb": {
			name:        "mapb",
			description: "Paginated list of location areas, 20 per page, backwards",
			callback:    callbackMapBack,
		}, 
		"explore": {
			name:        "explore",
			description: "list of pokemon encounters in a specific location area\n  usage: explore [location area name]",
			callback:    callbackExplore,
		}, 
		"pokemon": {
			name:        "pokemon",
			description: "short description of the specified pokemon\n  usage: pokemon [pokemon name]",
			callback:    callbackPokemon,
		}, 
		"catch": {
			name:        "catch",
			description: "try to catch the specified pokemon, on success, the pokemon will be added to the user pokedex\n  usage: catch [pokemon name]",
			callback:    callbackCatch,
		}, 
	}
}


// break up the string to lowercase words
func cleanupInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

