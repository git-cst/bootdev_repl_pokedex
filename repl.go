package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	nextUrl     string
	previousUrl string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := createCommands()
	config := Config{}

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

		command, ok := commands[cleanUserInput[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			callbackError := command.callback(&config)

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
