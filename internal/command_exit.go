package internal

import (
	"fmt"
	"os"
)

func CommandExit(c *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
