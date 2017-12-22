// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

// FixLeaguesNamesUp fixes league naming issues
func FixLeaguesNamesUp(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE `leagues` SET name='МИСТИКА' WHERE id=2")
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE `leagues` SET name='ОТВАГА' WHERE id=3")
	if err != nil {
		return err
	}
	return nil
}

// FixLeaguesNamesDown does nothing
func FixLeaguesNamesDown(tx *sql.Tx) error {
	return nil
}
