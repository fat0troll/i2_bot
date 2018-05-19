// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package datamapping

// League is a struct, which represents league data
type League struct {
	ID          int    `yaml:"id"`
	Symbol      string `yaml:"symbol"`
	Name        string `yaml:"name"`
	NameEnglish string `yaml:"name_english"`
}
