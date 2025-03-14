package commands

import (
	"fmt"

	"github.com/git-cst/bootdev_pokedex/internal/config"
)

func commandInspect(c *config.Config, args ...any) error {
	if len(args) == 0 {
		return fmt.Errorf("input pokemon (%s) was of length 0", args)
	}

	pokemon, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("args %v provided was not of the expected type string", pokemon)
	}

	// Check if the key exists in the cache
	pokemonData, exists := c.Pokedex.CaughtPokemon[pokemon]
	// If key exists in the cache unmarshal the bytes,
	if exists {
		fmt.Printf("Name: %s\n", pokemonData.Name)
		fmt.Printf("Height: %d\n", pokemonData.Height)
		fmt.Printf("Weight: %d\n", pokemonData.Weight)
		fmt.Println("Stats:")

		for Stat, StatValues := range pokemonData.Stats {
			fmt.Printf(" - %s: %d\n", Stat, StatValues.BaseValue)
		}

		fmt.Println("Types:")

		for _, pokemontype := range pokemonData.Types {
			fmt.Printf(" - %s\n", pokemontype)
		}

		return nil
	}

	// Key doesn't exist therefore we haven't caught it yet
	fmt.Printf("you have not caught %s yet...\n", pokemon)
	return nil
}
