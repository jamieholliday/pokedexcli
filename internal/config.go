package internal

import (
	"time"
)

type Config struct {
	NextUrl          string
	PrevUrl          string
	Cache            *Cache
	LocationEndpoint string
	PokemonEndpoint  string
	Pokedex          map[string]Pokemon
}

var api = "https://pokeapi.co/api/v2"
var LocationEndpoint = api + "/location-area"

func CreateConfig() *Config {
	return &Config{
		LocationEndpoint: LocationEndpoint,
		PokemonEndpoint:  api + "/pokemon",
		NextUrl:          LocationEndpoint,
		PrevUrl:          "",
		Cache:            NewCache(5 * time.Minute),
		Pokedex:          make(map[string]Pokemon),
	}
}
