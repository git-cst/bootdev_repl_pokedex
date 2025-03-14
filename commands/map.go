package commands

import (
	"encoding/json"
	"fmt"

	"github.com/git-cst/bootdev_pokedex/internal/config"
	"github.com/git-cst/bootdev_pokedex/internal/pokeapi"
)

func commandMap(c *config.Config, args ...any) error {
	var endpoint string
	if len(c.NextUrl) == 0 {
		endpoint = "https://pokeapi.co/api/v2/location-area"
	} else {
		endpoint = c.NextUrl
	}

	return fetchLocation(c, endpoint)
}

func commandMapb(c *config.Config, args ...any) error {
	if len(c.PreviousUrl) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	return fetchLocation(c, c.PreviousUrl)
}

func fetchLocation(c *config.Config, endpoint string) error {
	// Check if the key exists in the cache
	value, exists := c.Cache.Get(endpoint)

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

	// Key does not exist in the cache therefore get locations, print the locations and add to cache
	locations, err := pokeapi.GetLocation(endpoint)

	if err != nil {
		return err
	}

	locationHandler(c, &locations)

	cacheBytes, err := json.Marshal(locations)

	if err != nil {
		return err
	}

	c.Cache.Add(endpoint, cacheBytes)
	return nil
}

func locationHandler(c *config.Config, l *pokeapi.LocationRequest) {
	for i := range l.Results {
		fmt.Printf("%v\n", l.Results[i].Name)
	}

	c.NextUrl = l.Next
	c.PreviousUrl = l.Previous
}
