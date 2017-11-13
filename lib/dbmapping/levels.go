// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// Level is a struct, which represents `levels` table item in databse.
type Level struct {
	ID        int       `db:"id"`
	MaxExp    int       `db:"max_exp"`
	MaxEgg    int       `db:"max_egg"`
	CreatedAt time.Time `db:"created_at"`
}
