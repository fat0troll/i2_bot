// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmappings

import (
    // stdlib
    "time"
)

type Leagues struct {
    Id              int             `db:"id"`
    Symbol          string          `db:"symbol"`
    Name            string          `db:"league_id"`
    Created_at      *time.Time      `db:"created_at"`
}
