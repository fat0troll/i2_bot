// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// SquadPlayer is a struct, which represents `squads_players` table item in databse.
type SquadPlayer struct {
	ID        int       `db:"id"`
	SquadID   int       `db:"squad_id"`
	PlayerID  int       `db:"player_id"`
	UserType  string    `db:"user_type"`
	AuthorID  int       `db:"author_id"`
	CreatedAt time.Time `db:"created_at"`
}

// SquadPlayerFull is a struct, which handles all related information
type SquadPlayerFull struct {
	Squad    SquadChat
	Player   Player
	Profile  Profile
	UserRole string
}
