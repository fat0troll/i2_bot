// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// PokememeElement is a struct, which represents `pokememes_elements` table item in databse.
type PokememeElement struct {
	ID         int       `db:"id"`
	PokememeID int       `db:"pokememe_id"`
	ElementID  int       `db:"element_id"`
	CreatedAt  time.Time `db:"created_at"`
}
