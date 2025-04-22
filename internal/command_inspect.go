package internal

import (
	"fmt"
)

func CommandInspect(c *Config, args []string) error {
	if args == nil || len(args) == 0 {
		fmt.Println("Please provide a Pokemon name.")
		return nil
	}

	if pokemon, ok := c.Pokedex[args[0]]; !ok {
		fmt.Println("Pokemon not found in Pokedex.")

	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("- %s\n", pokemonType.Type.Name)
		}
	}
	return nil
}
