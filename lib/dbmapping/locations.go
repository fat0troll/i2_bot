// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// Location is a struct, which represents `locations` table item in databse.
type Location struct {
	ID        int       `db:"id"`
	Symbol    string    `db:"symbol"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
