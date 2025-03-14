package commands

import (
	"github.com/git-cst/bootdev_pokedex/internal/config"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(*config.Config, ...any) error
}

func CreateCommands() map[string]cliCommand {
	commands := make(map[string]cliCommand)

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the pokedex",
		Callback:    commandExit,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		Callback:    commandHelp,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 map locations",
		Callback:    commandMap,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 map locations",
		Callback:    commandMapb,
	}

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Displays the pokémon available at the provided location",
		Callback:    commandExplore,
	}

	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Attempts to catch the pokemon provided as an argument",
		Callback:    commandCatch,
	}

	commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Shows the information about a pokemon you've already caught",
		Callback:    commandInspect,
	}

	commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Displays all pokemon that you have caught so far",
		Callback:    commandPokedex,
	}

	return commands
}
