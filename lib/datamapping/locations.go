// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package datamapping

// Location is a struct, which represents location data
type Location struct {
	ID     int    `yaml:"id"`
	Symbol string `yaml:"symbol"`
	Name   string `yaml:"name"`
}
