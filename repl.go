package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {

	// read stin
	scanner := bufio.NewScanner(os.Stdin)
	
	// REPL running
	for {
		fmt.Printf(" > ")
		scanner.Scan()
		text := scanner.Text()
		
		cleaned := cleanupInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		_, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue		
		}
		availableCommands[commandName].callback()
	}
}


type cliCommand struct {
	name string
	description string
	callback func() error
}
	

// list all available commands (help)
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
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
	}
}


// break up the string to lowercase words
func cleanupInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

