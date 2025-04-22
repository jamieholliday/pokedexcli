package internal

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	// "math/rand"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

func CommandCatch(c *Config, args []string) error {
	if args == nil || len(args) == 0 {
		fmt.Println("Please provide a Pokemon name.")
		return nil
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	url := c.PokemonEndpoint + "/" + pokemonName
	getPokemon(c, url)
	return nil
}

func getPokemon(c *Config, url string) error {
	pokemon, err := GetCachedData(c, url)
	if err != nil {
		return err
	}
	pokemonJson, err := getPokemonJson(pokemon)
	if err != nil {
		return err
	}
	tryCatchPokemon(c, pokemonJson)
	return nil
}

func getPokemonJson(data []byte) (*Pokemon, error) {
	pokemon := &Pokemon{}
	json.Unmarshal(data, pokemon)
	return pokemon, nil
}

func tryCatchPokemon(c *Config, pokemon *Pokemon) {
	randNum := rand.Float32()
	rate := 1.0 - (float32(pokemon.BaseExperience) / 255.0)
	result := randNum < rate
	if result {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		c.Pokedex[pokemon.Name] = *pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

}
