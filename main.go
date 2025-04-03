package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)
		if len(cleanedText) == 0 {
			continue
		}
		firstWord := cleanedText[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}

func cleanInput(text string) []string {
	words := []string{}
	for _, word := range strings.Fields(text) {
		word = strings.TrimSpace(word)
		word = strings.ToLower(word)
		words = append(words, word)
	}
	return words
}
