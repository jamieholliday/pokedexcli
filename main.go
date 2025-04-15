package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jamieholliday/pokedexcli/internal"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	nextUrl string
	prevUrl string
	cache   *internal.Cache
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

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int `json:"chance"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var conf = &config{
	nextUrl: "https://pokeapi.co/api/v2/location-area",
	prevUrl: "",
	cache:   internal.NewCache(5 * time.Minute),
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
	getLocationsStrings(c, c.nextUrl)
	return nil
}

func commandMapb(c *config) error {
	getLocationsStrings(c, c.prevUrl)
	return nil
}

func getLocationsStrings(c *config, url string) error {
	if url == "" {
		fmt.Println("You are on the first page.")
		return nil
	}
	var locations []byte
	if c.cache != nil {
		if cachecLocationsData, exists := c.cache.Get(url); exists {
			locations = cachecLocationsData
		}
	}
	if len(locations) == 0 {
		locs, err := getLocationsApi(c, url)
		locations = locs
		c.cache.Add(url, locations)
		if err != nil {
			return err
		}
	}
	locationsJson, err := getLocationsJson(c, locations)
	if err != nil {
		return err
	}
	for _, location := range locationsJson {
		fmt.Println(location)
	}
	return nil
}

func getLocationsApi(c *config, url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, nil

}

func getLocationsJson(c *config, data []byte) ([]string, error) {
	locationAreas := &LocationAreas{}
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
