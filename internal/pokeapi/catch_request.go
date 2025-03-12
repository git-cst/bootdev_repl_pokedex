package pokeapi

type PokemonInfo struct {
	Name string
}

func PokemonRequest(URL string) (PokemonInfo, error) {
	return PokemonInfo{}, nil
}
