// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

type Profile struct {
	Id               int       `db:"id"`
	Player_id        int       `db:"player_id"`
	Nickname         string    `db:"nickname"`
	TelegramNickname string    `db:"telegram_nickname"`
	Level_id         int       `db:"level_id"`
	Pokeballs        int       `db:"pokeballs"`
	Wealth           int       `db:"wealth"`
	Exp              int       `db:"exp"`
	Egg_exp          int       `db:"egg_exp"`
	Power            int       `db:"power"`
	Weapon_id        int       `db:"weapon_id"`
	Crystalls        int       `db:"crystalls"`
	Created_at       time.Time `db:"created_at"`
}
