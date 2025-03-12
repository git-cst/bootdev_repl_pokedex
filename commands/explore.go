package commands

import (
	"encoding/json"
	"fmt"

	"github.com/git-cst/bootdev_pokedex/internal/config"
	"github.com/git-cst/bootdev_pokedex/internal/pokeapi"
	"github.com/git-cst/bootdev_pokedex/internal/pokecache"
)

func commandExplore(c *config.Config, ca *pokecache.Cache, args ...any) error {
	if len(args) == 0 {
		return fmt.Errorf("input location (%s) was of length 0", args)
	}

	location, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("args %v provided was not of the expected type string", location)
	}

	fmt.Printf("Exploring %s...\n", location)
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)

	// Check if the key exists in the cache
	value, exists := ca.Get(endpoint)
	// If key exists in the cache unmarshal the bytes, print the pokemon, return
	if exists {
		exploration := pokeapi.ExploreRequest{}
		err := json.Unmarshal(value, &exploration)
		if err != nil {
			return err
		}

		exploreHandler(&exploration)

		return nil
	}

	// Key does not exist in the cache therefore explore the location, print the pokemon and add to cache
	exploration, err := pokeapi.ExploreLocation(endpoint)
	if err != nil {
		return err
	}

	exploreHandler(&exploration)

	cacheBytes, err := json.Marshal(exploration)
	if err != nil {
		return err
	}

	ca.Add(endpoint, cacheBytes)
	return nil
}

func exploreHandler(er *pokeapi.ExploreRequest) {
	fmt.Println("Found Pokémon:")
	for pokemon := range er.Pokemon {
		fmt.Printf(" - %s\n", pokemon)
	}
}
