package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/git-cst/bootdev_pokedex/internal/pokeapi"
	"github.com/git-cst/bootdev_pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, *pokecache.Cache, ...any) error
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

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Displays the pokémon available at the provided location",
		callback:    commandExplore,
	}

	return commands
}

func commandExit(c *Config, ca *pokecache.Cache, args ...any) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Pokedex did not exit as expected")
}

func commandHelp(c *Config, ca *pokecache.Cache, args ...any) error {
	helpMessage := "Welcome to the Pokedex!\nUsage:\n\n"
	commands := createCommands()

	for i := range commands {
		helpMessage += commands[i].name + ": " + commands[i].description + "\n"
	}

	fmt.Println(helpMessage)
	return nil
}

func commandExplore(c *Config, ca *pokecache.Cache, args ...any) error {
	if len(args) == 0 {
		return fmt.Errorf("input location (%s) was of length 0", args)
	}

	location, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("args %v provided was not of the expected type string", location)
	}

	return nil
}

func commandMap(c *Config, ca *pokecache.Cache, args ...any) error {
	var endpoint string
	if len(c.nextUrl) == 0 {
		endpoint = "https://pokeapi.co/api/v2/location-area"
	} else {
		endpoint = c.nextUrl
	}

	return fetchLocation(c, ca, endpoint)
}

func commandMapb(c *Config, ca *pokecache.Cache, args ...any) error {
	if len(c.previousUrl) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	return fetchLocation(c, ca, c.previousUrl)
}

func fetchLocation(c *Config, ca *pokecache.Cache, endpoint string) error {
	// Check if the key exists in the cache
	value, exists := ca.Get(endpoint)

	// If key exists in the cache unmarshal the bytes, print the locations, return
	if exists {
		locations := pokeapi.LocationRequest{}
		err := json.Unmarshal(value, &locations)

		if err != nil {
			return err
		}

		locationHandler(c, &locations)

		return nil
	}

	// Key does not exist in the cache therefore get locations, print the locations and add to cache
	locations, err := pokeapi.GetLocation(endpoint)

	if err != nil {
		return err
	}

	locationHandler(c, &locations)

	cacheBytes, err := json.Marshal(locations)

	if err != nil {
		return err
	}

	ca.Add(endpoint, cacheBytes)
	return nil
}

func locationHandler(c *Config, l *pokeapi.LocationRequest) {
	for i := range l.Results {
		fmt.Printf("%v\n", l.Results[i].Name)
	}

	c.nextUrl = l.Next
	c.previousUrl = l.Previous
}
