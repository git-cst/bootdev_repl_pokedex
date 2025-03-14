package commands

import (
	"fmt"
	"os"

	"github.com/git-cst/bootdev_pokedex/internal/config"
)

func commandExit(c *config.Config, args ...any) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Pokedex did not exit as expected")
}
