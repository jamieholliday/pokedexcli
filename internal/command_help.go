package internal

import (
	"fmt"
)

func CommandHelp(c *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range SupportedCommands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
