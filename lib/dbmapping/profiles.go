// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// Profile is a struct, which represents `profiles` table item in database.
type Profile struct {
	ID               int       `db:"id"`
	PlayerID         int       `db:"player_id"`
	Nickname         string    `db:"nickname"`
	TelegramNickname string    `db:"telegram_nickname"`
	LevelID          int       `db:"level_id"`
	Pokeballs        int       `db:"pokeballs"`
	Wealth           int       `db:"wealth"`
	PokememesWealth  int       `db:"pokememes_wealth"`
	Exp              int       `db:"exp"`
	EggExp           int       `db:"egg_exp"`
	Power            int       `db:"power"`
	WeaponID         int       `db:"weapon_id"`
	Crystalls        int       `db:"crystalls"`
	CreatedAt        time.Time `db:"created_at"`
}
