package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("No Pokemon have been caught yet")
	}

	fmt.Println("Your Pokedex:")
	for _, pok := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pok.Name)
	}

	return nil
}
