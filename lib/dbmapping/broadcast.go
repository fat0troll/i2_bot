// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

// Broadcast is a struct, which represents `broadcast` table item in databse.
type Broadcast struct {
	ID            int       `db:"id"`
	Text          string    `db:"text"`
	BroadcastType string    `db:"broadcast_type"`
	Status        string    `db:"status"`
	AuthorID      int       `db:"author_id"`
	CreatedAt     time.Time `db:"created_at"`
}
