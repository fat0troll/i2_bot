// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

type Player struct {
	Id          int       `db:"id"`
	Telegram_id int       `db:"telegram_id"`
	League_id   int       `db:"league_id"`
	Squad_id    int       `db:"squad_id"`
	Status      string    `db:"status"`
	Created_at  time.Time `db:"created_at"`
	Updated_at  time.Time `db:"updated_at"`
}
