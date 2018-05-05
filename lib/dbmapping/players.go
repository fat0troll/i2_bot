// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"

	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
)

// Player is a struct, which represents `players` table item in databse.
type Player struct {
	ID         int       `db:"id"`
	TelegramID int       `db:"telegram_id"`
	LeagueID   int       `db:"league_id"`
	Status     string    `db:"status"`
	Karma      int       `db:"karma"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// PlayerProfile is a struch which handles all user information
type PlayerProfile struct {
	Player      Player
	Profile     Profile
	League      datamapping.League
	HaveProfile bool
}
