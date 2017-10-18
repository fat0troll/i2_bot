// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

type PokememeLocation struct {
	Id          int       `db:"id"`
	Pokememe_id int       `db:"pokememe_id"`
	Location_id int       `db:"location_id"`
	Created_at  time.Time `db:"created_at"`
}
