// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package datamapping

// Weapon is a struct, which represents weapon data
type Weapon struct {
	ID    int    `yaml:"id"`
	Name  string `yaml:"name"`
	Power int    `yaml:"power"`
	Price int    `yaml:"price"`
}
