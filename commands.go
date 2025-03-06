package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func createCommands() map[string]cliCommand {
	commands := make(map[string]cliCommand)
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	return commands
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Pokedex did not exit as expected")
}

func commandHelp() error {
	helpMessage := "Welcome to the Pokedex!\nUsage:\n\n"
	commands := createCommands()

	for i := range commands {
		helpMessage += commands[i].name + ": " + commands[i].description + "\n"
	}

	fmt.Println(helpMessage)
	return nil
}
