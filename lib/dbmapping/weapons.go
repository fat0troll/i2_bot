// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
    // stdlib
    "time"
)

type Weapon struct {
    Id              int             `db:"id"`
    Name            string          `db:"name"`
    Power           int             `db:"power"`
    Price           int             `db:"price"`
    Created_at      time.Time       `db:"created_at"`
}