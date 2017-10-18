// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

type Element struct {
	Id         int        `db:"id"`
	Symbol     string     `db:"symbol"`
	Name       string     `db:"name"`
	League_id  int        `db:"league_id"`
	Created_at *time.Time `db:"created_at"`
}
