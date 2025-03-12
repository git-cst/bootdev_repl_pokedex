package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/git-cst/bootdev_pokedex/commands"
	"github.com/git-cst/bootdev_pokedex/internal/config"
	"github.com/git-cst/bootdev_pokedex/internal/pokecache"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := commands.CreateCommands()
	config := config.Config{}
	cache := pokecache.NewCache()

	for {
		fmt.Print("Pokedex > ")
		success := scanner.Scan()
		if !success {
			return
		}

		userInput := scanner.Text()
		if len(userInput) == 0 {
			continue
		}

		cleanUserInput := cleanInput(userInput)

		var arg any
		if len(cleanUserInput) > 1 {
			arg = cleanUserInput[1]
		} else {
			arg = nil
		}

		command, ok := commands[cleanUserInput[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			callbackError := command.Callback(&config, cache, arg)

			if callbackError != nil {
				fmt.Println(callbackError)
			}
		}
	}
}

func cleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	words := strings.Fields(lowerString)
	return words
}
