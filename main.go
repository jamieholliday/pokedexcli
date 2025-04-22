package main

import (
	"bufio"
	"fmt"
	"github.com/jamieholliday/pokedexcli/internal"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	words := []string{}
	for _, word := range strings.Fields(text) {
		word = strings.TrimSpace(word)
		word = strings.ToLower(word)
		words = append(words, word)
	}
	return words
}

func init() {
	internal.AddCliCommand("help", "Displays a help message", internal.CommandHelp)
	internal.AddCliCommand("exit", "Exit the Pokedex", internal.CommandExit)
	internal.AddCliCommand("map", "Get the next locations", internal.CommandMap)
	internal.AddCliCommand("mapb", "Get the prev locations", internal.CommandMapb)
	internal.AddCliCommand("explore", "Explore a location", internal.CommandExplore)
	internal.AddCliCommand("catch", "Try and catch a pokemon", internal.CommandCatch)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := internal.CreateConfig()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)
		if len(cleanedText) == 0 {
			continue
		}
		commandWord := cleanedText[0]
		args := cleanedText[1:]
		cmd, exists := internal.GetCliCommand(commandWord)
		if !exists {
			fmt.Printf("Unknown command: %s\n", commandWord)
			continue
		}
		if cmd.Callback == nil {
			fmt.Printf("Command not implemented: %s\n", commandWord)
			continue
		}
		err := cmd.Callback(conf, args)
		if err != nil {
			fmt.Printf("Error exexuting command: %s\n", err)
		}
	}
}
