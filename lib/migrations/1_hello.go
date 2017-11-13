// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// HelloUp is the first migration, added for testing purposes
func HelloUp(tx *sql.Tx) error {
	c.Log.Printf("Migration framework loaded. All systems are OK.")

	return nil
}
