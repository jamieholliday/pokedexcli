package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	nextUrl string
	prevUrl string
}

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var conf = &config{
	nextUrl: "https://pokeapi.co/api/v2/location-area",
	prevUrl: "",
}

var supportedCommands = map[string]cliCommand{}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range supportedCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *config) error {
	if c.nextUrl == "" {
		fmt.Println("You are on the last page.")
		return nil
	}
	locations, err := getLocations(c, c.nextUrl)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}

func commandMapb(c *config) error {
	if c.prevUrl == "" {
		fmt.Println("You are on the first page.")
		return nil
	}
	locations, err := getLocations(c, c.prevUrl)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}

func getLocations(c *config, url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	defer res.Body.Close()
	locationAreas := &LocationAreas{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}
	json.Unmarshal(data, locationAreas)
	if locationAreas.Next == nil {
		c.nextUrl = ""
	} else {
		c.nextUrl = *locationAreas.Next
	}
	if locationAreas.Previous == nil {
		c.prevUrl = ""
	} else {
		c.prevUrl = *locationAreas.Previous
	}
	locationNames := []string{}
	for _, result := range locationAreas.Results {
		locationNames = append(locationNames, result.Name)
	}

	return locationNames, nil
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

func init() {
	supportedCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	supportedCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	supportedCommands["map"] = cliCommand{
		name:        "map",
		description: "Get the next locations",
		callback:    commandMap,
	}
	supportedCommands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Get the prev locations",
		callback:    commandMapb,
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
		err := cmd.callback(conf)
		if err != nil {
			fmt.Printf("Error exexuting command: %s\n", err)
		}
	}
}
