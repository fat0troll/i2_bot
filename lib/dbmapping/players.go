// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

// Player is a struct, which represents `players` table item in databse.
type Player struct {
	ID         int       `db:"id"`
	TelegramID int       `db:"telegram_id"`
	LeagueID   int       `db:"league_id"`
	Status     string    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
