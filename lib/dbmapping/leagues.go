// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// League is a struct, which represents `leagues` table item in databse.
type League struct {
	ID        int        `db:"id"`
	Symbol    string     `db:"symbol"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
}
