// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
	"errors"
)

// DeletePokememesTablesUp drops `pokememes`, `pokememes_elements` and `pokememes_locations` tables
// These tables data is rarely changed, so I decided to hold such data in source code
func DeletePokememesTablesUp(tx *sql.Tx) error {
	request := "DROP TABLE IF EXISTS `pokememes`"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	request = "DROP TABLE IF EXISTS `pokememes_elements`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "DROP TABLE IF EXISTS `pokememes_locations`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

// DeletePokememesTablesDown does nothing, because after nerf old information isn't needed at all
func DeletePokememesTablesDown(tx *sql.Tx) error {
	return errors.New("This migration is irreversible, as nerf itself")
}
