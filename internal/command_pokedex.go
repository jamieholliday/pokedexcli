package internal

import (
	"fmt"
)

func CommandPokedex(c *Config, args []string) error {
	if len(c.Pokedex) > 0 {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range c.Pokedex {
			fmt.Printf("- %s\n", pokemon.Name)
		}
		return nil
	}
	fmt.Println("Pokedex is empty.")
	return nil
}
