// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// Squad is a struct, which represents `squads` table item in databse.
type Squad struct {
	ID         int       `db:"id"`
	ChatID     int       `db:"chat_id"`
	MinLevel   int       `db:"min_level"`
	MaxLevel   int       `db:"max_level"`
	InviteLink string    `db:"invite_link"`
	CreatedAt  time.Time `db:"created_at"`
}

// SquadChat is a stuct, which combines information about chats and squads
type SquadChat struct {
	Squad Squad
	Chat  Chat
}
