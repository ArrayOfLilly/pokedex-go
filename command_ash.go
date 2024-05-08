package main

import (
	"fmt"
)


func callbackAsh(cfg *config, param string) error {
	fmt.Println("\n🐚  You're the very best!")
	return nil
}
