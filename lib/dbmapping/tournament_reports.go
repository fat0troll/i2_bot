// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// TournamentReport is a struct, which represents `tournament_reports` table item in database.
type TournamentReport struct {
	ID               int       `db:"id"`
	PlayerID         int       `db:"player_id"`
	TournamentNumber int       `db:"tournament_number"`
	Target           string    `db:"target"`
	CreatedAt        time.Time `db:"created_at"`
}
