// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
    // stdlib
    "time"
)

type PokememeElement struct {
    Id              int             `db:"id"`
    Pokememe_id     int             `db:"pokememe_id"`
    Element_id      int             `db:"element_id"`
    Created_at      time.Time       `db:"created_at"`
}