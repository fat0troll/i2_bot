// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// AddIsActiveToPokememesUp adds `is_active` to `pokememes`
func AddIsActiveToPokememesUp(tx *sql.Tx) error {
	request := "ALTER TABLE `pokememes` ADD COLUMN `is_active` tinyint(1) DEFAULT 1 NOT NULL COMMENT 'Является ли покемем играющим в данный момент?'"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

// AddIsActiveToPokememesDown removes `is_active` from `pokememes` table
func AddIsActiveToPokememesDown(tx *sql.Tx) error {
	request := "ALTER TABLE `pokememes` DROP COLUMN `is_active`"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}
