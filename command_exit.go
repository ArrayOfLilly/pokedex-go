package main

import (
	"fmt"
	"os"
	"time"
)

// exit
func callbackExit() error {

	fmt.Print("\033[2J") // Clear Screen
	fmt.Print("\033[H")  // Cursor Position

	time.Sleep(500)

	os.Exit(0)

	return nil
}