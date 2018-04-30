// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package datamapping

// Element is a struct, which represents element data
type Element struct {
	ID       int    `yaml:"id"`
	Symbol   string `yaml:"symbol"`
	Name     string `yaml:"name"`
	LeagueID int    `yaml:"league_id"`
}
