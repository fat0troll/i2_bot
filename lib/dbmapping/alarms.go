// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// Alarm is a struct, which represents `alarms` table item in database.
type Alarm struct {
	ID           int       `db:"id"`
	PlayerID     int       `db:"player_id"`
	TurnirNumber int       `db:"turnir_number"` // From 1 to 12
	CreatedAt    time.Time `db:"created_at"`
}
