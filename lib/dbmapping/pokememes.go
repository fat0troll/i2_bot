// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

// Pokememe is a struct, which represents `pokememes` table item in databse.
type Pokememe struct {
	ID           int       `db:"id"`
	Grade        int       `db:"grade"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	Attack       int       `db:"attack"`
	HP           int       `db:"hp"`
	MP           int       `db:"mp"`
	Defence      int       `db:"defence"`
	Price        int       `db:"price"`
	Purchaseable bool      `db:"purchaseable"`
	ImageURL     string    `db:"image_url"`
	PlayerID     int       `db:"player_id"`
	CreatedAt    time.Time `db:"created_at"`
}

// PokememeFull is a struct for handling pokememe with all informations about locations and elements
type PokememeFull struct {
	Pokememe  Pokememe
	Locations []Location
	Elements  []Element
}
