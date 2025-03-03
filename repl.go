package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	words := strings.Fields(lowerString)
	return words
}
