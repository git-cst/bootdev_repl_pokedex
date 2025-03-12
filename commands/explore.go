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

	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", args)

	// Check if the key exists in the cache
	value, exists := ca.Get(endpoint)
	// If key exists in the cache unmarshal the bytes, print the locations, return
	if exists {
		locations := pokeapi.LocationRequest{}
		err := json.Unmarshal(value, &locations)

		if err != nil {
			return err
		}

		locationHandler(c, &locations)

		return nil
	}

	return nil
}
