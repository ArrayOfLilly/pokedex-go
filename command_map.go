package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, param string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("-----------------------------------")
	fmt.Println("Location Areas:")
	fmt.Println("-----------------------------------")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous

	fmt.Println("")
	return nil
}

func callbackMapBack(cfg *config, param string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("-----------------------------------")
	fmt.Println("Location Areas:")
	fmt.Println("-----------------------------------")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous

	fmt.Println("")
	return nil
}
