// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func FixTimeElementUp(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE `elements` SET league_id=3 WHERE symbol='⌛';")
	if err != nil {
		return err
	}

	return nil
}

func FixTimeElementDown(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE `elements` SET league_id=1 WHERE symbol='⌛';")
	if err != nil {
		return err
	}

	return nil
}
