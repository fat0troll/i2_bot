// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

// AddKarmaToPlayersUp creates `karma` column in `players` table
func AddKarmaToPlayersUp(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `players` ADD COLUMN `karma` INT(11) NOT NULL DEFAULT 250 COMMENT 'Карма игрока' AFTER `status`")
	if err != nil {
		return err
	}

	return nil
}

// AddKarmaToPlayersDown destroys `karma` column
func AddKarmaToPlayersDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `players` DROP COLUMN `karma`")
	if err != nil {
		return err
	}

	return nil
}
