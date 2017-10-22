// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	// stdlib
	"time"
)

// ProfilePokememe is a struct, which represents `profiles_pokememes` table item in databse.
type ProfilePokememe struct {
	ID             int       `db:"id"`
	ProfileID      int       `db:"profile_id"`
	PokememeID     int       `db:"pokememe_id"`
	PokememeAttack int       `db:"pokememe_attack"`
	PokememeRarity string    `db:"pokememe_rarity"`
	CreatedAt      time.Time `db:"created_at"`
}
