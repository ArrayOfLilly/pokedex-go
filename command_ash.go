package main

import (
	"fmt"
)

func callbackAsh(cfg *config, pkdx *map[string]Pokemon, param string) error {
	fmt.Println("\n🐚  You're the very best!")
	fmt.Println("")
	return nil
}
