package commands

import (
	"errors"
	"fmt"

	"github.com/git-cst/bootdev_pokedex/internal/config"
)

func commandPokedex(c *config.Config, args ...any) error {
	fmt.Println("Your pokedex:")

	if len(c.Pokedex.CaughtPokemon) == 0 {
		return errors.New("no pokemon caught yet")
	}

	for _, pokemon := range c.Pokedex.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
