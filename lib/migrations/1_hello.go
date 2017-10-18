// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
	"log"
)

// First migration, added for testing purposes

func HelloUp(tx *sql.Tx) error {
	log.Printf("Migration framework loaded. All systems are OK.")

	return nil
}
