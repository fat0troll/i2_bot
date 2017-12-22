// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

// FixLocationsUp fixes location issues
func FixLocationsUp(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE `locations` SET symbol='🏙' WHERE id=4")
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE `locations` SET symbol='⛪️' WHERE id=6")
	if err != nil {
		return err
	}
	return nil
}

// FixLocationsDown does nothing
func FixLocationsDown(tx *sql.Tx) error {
	return nil
}
