// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
    // stdlib
    "time"
)

type Pokememe struct {
    Id              int             `db:"id"`
    Grade           int             `db:"grade"`
    Name            string          `db:"name"`
    Description     string          `db:"description"`
    Attack          int             `db:"attack"`
    HP              int             `db:"hp"`
    MP              int             `db:"mp"`
    Defence         int             `db:"defence"`
    Price           int             `db:"price"`
    Purchaseable    bool            `db:"purchaseable"`
    Image_url       string          `db:"image_url"`
    Player_id       int             `db:"player_id"`
    Created_at      time.Time       `db:"created_at"`
}

// Type for handling pokememe with all informations about locations and elements

type PokememeFull struct {
    Pokememe        Pokememe
    Locations       []Location
    Elements        []Element
}
