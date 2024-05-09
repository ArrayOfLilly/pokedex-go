package main

import (
	"fmt"
)

// exit
func callbackClear(cfg *config, pkdx *map[string]Pokemon, param string) error {

	fmt.Print("\033[2J") // Clear Screen
	fmt.Print("\033[H")  // Cursor Position
	return nil
}
