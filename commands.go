package main

import (
	"fmt"
	"os"

	"github.com/git-cst/bootdev_pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 map locations",
		callback:    commandMap,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 map locations",
		callback:    commandMapb,
	}

	return commands
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Pokedex did not exit as expected")
}

func commandHelp(c *Config) error {
	helpMessage := "Welcome to the Pokedex!\nUsage:\n\n"
	commands := createCommands()

	for i := range commands {
		helpMessage += commands[i].name + ": " + commands[i].description + "\n"
	}

	fmt.Println(helpMessage)
	return nil
}

func commandMap(c *Config) error {
	var endpoint string
	if len(c.nextUrl) == 0 {
		endpoint = "https://pokeapi.co/api/v2/location-area"
	} else {
		endpoint = c.nextUrl
	}

	locations, err := pokeapi.GetLocation(endpoint)

	if err != nil {
		return err
	}

	c.nextUrl = locations.Next
	c.previousUrl = locations.Previous

	for i := range locations.Results {
		fmt.Printf("%v\n", locations.Results[i].Name)
	}

	return nil
}

func commandMapb(c *Config) error {
	var endpoint string
	if len(c.previousUrl) == 0 {
		fmt.Println("you're on the first page")
		return nil
	} else {
		endpoint = c.previousUrl
	}

	locations, err := pokeapi.GetLocation(endpoint)

	if err != nil {
		return err
	}

	c.nextUrl = locations.Next
	c.previousUrl = locations.Previous

	for i, _ := range locations.Results {
		fmt.Printf("%v\n", locations.Results[i].Name)
	}

	return nil
}
