package commands

import (
	"fmt"

	"github.com/git-cst/bootdev_pokedex/internal/config"
	"github.com/git-cst/bootdev_pokedex/internal/pokeapi"
	"github.com/git-cst/bootdev_pokedex/internal/pokecache"
)

func commandCatch(c *config.Config, ca *pokecache.Cache, args ...any) error {
	if len(args) == 0 {
		return fmt.Errorf("input pokemon (%s) was of length 0", args)
	}

	pokemon, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("args %v provided was not of the expected type string", pokemon)
	}

	fmt.Printf("Throwing a ball at %s\n...", pokemon)

	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	pokemonData, err := pokeapi.PokemonRequest(endpoint)

	if err != nil {
		return nil
	}

	fmt.Printf("%v", pokemonData)

	return nil
}
