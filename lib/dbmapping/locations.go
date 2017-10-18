// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

type Location struct {
	Id         int       `db:"id"`
	Symbol     string    `db:"symbol"`
	Name       string    `db:"name"`
	Created_at time.Time `db:"created_at"`
}
