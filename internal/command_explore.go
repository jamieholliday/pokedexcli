package internal

import (
	"encoding/json"
	"fmt"
)

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

func CommandExplore(c *Config, args []string) error {
	if args == nil || len(args) == 0 {
		fmt.Println("Please provide a location name.")
		return nil
	}
	locationName := args[0]
	url := c.LocationEndpoint + "/" + locationName
	fmt.Printf("Exploring %s...\n", locationName)
	getLocation(c, url)
	return nil
}

func getLocation(c *Config, url string) error {
	location, err := GetCachedData(c, url)
	if err != nil {
		return err
	}
	locationJson, err := getLocationJson(location)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, location := range locationJson {
		fmt.Println(location)
	}
	return nil
}

func getLocationJson(data []byte) ([]string, error) {
	location := &Location{}
	json.Unmarshal(data, location)
	pokemonNames := []string{}
	for _, result := range location.PokemonEncounters {
		pokemonNames = append(pokemonNames, result.Pokemon.Name)
	}

	return pokemonNames, nil
}
