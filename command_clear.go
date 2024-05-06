package main

import (
	"fmt"
)

// exit
func callbackClear(cfg *config) error {

	fmt.Print("\033[2J") // Clear Screen
	fmt.Print("\033[H")  // Cursor Position
	return nil
}