package pokedex

import "github.com/git-cst/bootdev_pokedex/internal/pokeapi"

type Pokedex struct {
	CaughtPokemon map[string]pokeapi.Pokemon
}
