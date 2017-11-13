// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// Weapon is a struct, which represents `weapons` table item in databse.
type Weapon struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Power     int       `db:"power"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}
