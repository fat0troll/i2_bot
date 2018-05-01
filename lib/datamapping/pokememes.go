// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package datamapping

// Pokememe is a struct, which represents pokememes item data
type Pokememe struct {
	ID           int    `yaml:"id"`
	Grade        int    `yaml:"grade"`
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	Attack       int    `yaml:"attack"`
	HP           int    `yaml:"health"`
	MP           int    `yaml:"mana"`
	Defence      int    `db:"defence"`
	Price        int    `yaml:"cost"`
	Purchaseable bool   `yaml:"purchaseable"`
	ImageURL     string `yaml:"image"`
	Elements     []int  `yaml:"elements"`
	Locations    []int  `yaml:"locations"`
}

// PokememeFull is a struct for handling pokememe with all informations about locations and elements
type PokememeFull struct {
	Pokememe  Pokememe
	Locations []Location
	Elements  []Element
}
