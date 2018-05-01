// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package datamapping

// Level is a struct, which represents level data
type Level struct {
	ID         int `yaml:"id"`
	MaxExp     int `yaml:"max_exp"`
	MaxEgg     int `yaml:"max_egg"`
	LevelStart int `yaml:"level_start"`
}
