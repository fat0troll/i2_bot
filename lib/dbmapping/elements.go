// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// Element is a struct, which represents `elements` table item in databse.
type Element struct {
	ID        int        `db:"id"`
	Symbol    string     `db:"symbol"`
	Name      string     `db:"name"`
	LeagueID  int        `db:"league_id"`
	CreatedAt *time.Time `db:"created_at"`
}
