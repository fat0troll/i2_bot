// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

// Chat is a struct, which represents `chats` table item in databse.
type Chat struct {
	ID         int        `db:"id"`
	Name       string     `db:"name"`
	ChatType   bool       `db:"chat_type"`
	TelegramID int        `db:"telegram_id"`
	CreatedAt  *time.Time `db:"created_at"`
}
