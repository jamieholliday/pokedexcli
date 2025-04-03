package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func init() {
	// Add help command on init as it references the supportedCommands map
	supportedCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
}

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
		cmd, exists := supportedCommands[firstWord]
		if !exists {
			fmt.Printf("Unknown command: %s\n", firstWord)
			continue
		}
		if cmd.callback == nil {
			fmt.Printf("Command not implemented: %s\n", firstWord)
			continue
		}
		err := cmd.callback()
		if err != nil {
			fmt.Printf("Error exexuting command: %s\n", err)
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range supportedCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
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
