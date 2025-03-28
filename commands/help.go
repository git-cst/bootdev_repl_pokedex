package commands

import (
	"fmt"

	"github.com/git-cst/bootdev_pokedex/internal/config"
)

func commandHelp(c *config.Config, args ...any) error {
	helpMessage := "Welcome to the Pokedex!\nUsage:\n\n"
	commands := CreateCommands()

	for i := range commands {
		helpMessage += commands[i].name + ": " + commands[i].description + "\n"
	}

	fmt.Println(helpMessage)
	return nil
}
