// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmappings

import (
    // stdlib
    "time"
)

type Profiles struct {
    Id                  int             `db:"id"`
    Player_id           int             `db:"player_id"`
    Nickname            string          `db:"nickname"`
    TelegramNickname    string          `db:"telegram_nickname"`
    Level_id            int             `db:"level_id"`
    Exp                 int             `db:"exp"`
    Egg_exp             int             `db:"egg_exp"`
    Power               int             `db:"power"`
    Weapon_id           int             `db:"weapon_id"`
    Crystalls           int             `db:"crystalls"`
    Created_at          *time.Time      `db:"created_at"`
}
