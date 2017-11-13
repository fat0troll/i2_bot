// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// PokememeLocation is a struct, which represents `pokememes_locations` table item in databse.
type PokememeLocation struct {
	ID         int       `db:"id"`
	PokememeID int       `db:"pokememe_id"`
	LocationID int       `db:"location_id"`
	CreatedAt  time.Time `db:"created_at"`
}
