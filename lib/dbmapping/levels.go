// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
    // stdlib
    "time"
)

type Level struct {
    Id                  int             `db:"id"`
    Max_exp             int             `db:"max_exp"`
    Max_egg             int             `db:"max_egg"`
    Created_at          time.Time       `db:"created_at"`
}
