// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmappings

import (
    // stdlib
    "time"
)

type ProfilesPokememes struct {
    Id              int             `db:"id"`
    Profile_id      int             `db:"profile_id"`
    Pokememe_id     int             `db:"pokememe_id"`
    Pokememe_lvl    int             `db:"pokememe_lvl"`
    Pokememe_rarity string          `db:"pokememe_rarity"`
    Created_at      time.Time       `db:"created_at"`
}
