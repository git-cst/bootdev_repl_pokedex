package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/git-cst/bootdev_pokedex/internal/config"
	"github.com/git-cst/bootdev_pokedex/internal/pokeapi"
)

func commandCatch(c *config.Config, args ...any) error {
	if len(args) == 0 {
		return fmt.Errorf("input pokemon (%s) was of length 0", args)
	}

	pokemon, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("args %v provided was not of the expected type string", pokemon)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	// Check if the key exists in the cache
	value, exists := c.Cache.Get(endpoint)
	// If key exists in the cache unmarshal the bytes, attempt catch, return
	if exists {
		pokemonData := pokeapi.Pokemon{}
		err := json.Unmarshal(value, &pokemonData)
		if err != nil {
			return err
		}

		catchHandler(c, &pokemonData)
		return nil
	}

	// Key does not exist in the cache therefore explore the location, print the pokemon and add to cache
	pokemonData, err := pokeapi.PokemonRequest(endpoint)
	if err != nil {
		return err
	}

	catchHandler(c, &pokemonData)

	cacheBytes, err := json.Marshal(pokemonData)
	if err != nil {
		return err
	}

	c.Cache.Add(endpoint, cacheBytes)
	return nil
}

func catchHandler(c *config.Config, p *pokeapi.Pokemon) {
	// setup catch chances
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	modifier := 0.4
	baseExperience := float64(p.BaseExperience)
	baseChance := 1000.00
	minChance := 200.00
	chanceToCatch := max(minChance, baseChance/(1+modifier*baseExperience/100.00))
	randomRoll := rng.Float64() * float64(baseChance)

	if randomRoll < chanceToCatch {
		fmt.Printf("%s was caught!\n", p.Name)
		c.Pokedex.CaughtPokemon[p.Name] = *p
	} else {
		fmt.Printf("%s escaped!\n", p.Name)
	}
}
