package commands

import (
	"fmt"
	"os"

	"github.com/git-cst/bootdev_pokedex/internal/config"
	"github.com/git-cst/bootdev_pokedex/internal/pokecache"
)

func commandExit(c *config.Config, ca *pokecache.Cache, args ...any) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Pokedex did not exit as expected")
}
