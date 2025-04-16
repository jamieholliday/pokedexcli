package internal

import (
	"encoding/json"
	"fmt"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap(c *Config, args []string) error {
	if c.NextUrl == "" {
		fmt.Println("You are on the first page.")
		return nil
	}
	getLocationsStrings(c, c.NextUrl)
	return nil
}

func CommandMapb(c *Config, args []string) error {
	if c.PrevUrl == "" {
		fmt.Println("You are on the last page.")
		return nil
	}
	getLocationsStrings(c, c.PrevUrl)
	return nil
}

func getLocationsStrings(c *Config, url string) error {
	locations, err := GetCachedData(c, url)
	if err != nil {
		return err
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

func getLocationsJson(c *Config, data []byte) ([]string, error) {
	locationAreas := &LocationAreas{}
	json.Unmarshal(data, locationAreas)
	if locationAreas.Next == nil {
		c.NextUrl = ""
	} else {
		c.NextUrl = *locationAreas.Next
	}
	if locationAreas.Previous == nil {
		c.PrevUrl = ""
	} else {
		c.PrevUrl = *locationAreas.Previous
	}
	locationNames := []string{}
	for _, result := range locationAreas.Results {
		locationNames = append(locationNames, result.Name)
	}

	return locationNames, nil
}
