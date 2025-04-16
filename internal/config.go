package internal

import (
	"time"
)

type Config struct {
	NextUrl string
	PrevUrl string
	Cache   *Cache
}

func CreateConfig() *Config {
	return &Config{
		NextUrl: "https://pokeapi.co/api/v2/location-area",
		PrevUrl: "",
		Cache:   NewCache(5 * time.Minute),
	}
}
