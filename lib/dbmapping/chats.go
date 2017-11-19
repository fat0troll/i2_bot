// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// Chat is a struct, which represents `chats` table item in databse.
type Chat struct {
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	ChatType   string    `db:"chat_type"`
	TelegramID int64     `db:"telegram_id"`
	CreatedAt  time.Time `db:"created_at"`
}

// ChatSquad is a stuct, which combines information about chats and squads
type ChatSquad struct {
	Chat     Chat
	Squad    Squad
	ChatRole string
}
