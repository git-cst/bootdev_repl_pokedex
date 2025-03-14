package config

import (
	"github.com/git-cst/bootdev_pokedex/internal/pokecache"
	"github.com/git-cst/bootdev_pokedex/internal/pokedex"
)

type Config struct {
	NextUrl     string
	PreviousUrl string
	Pokedex     pokedex.Pokedex
	Cache       *pokecache.Cache
}
