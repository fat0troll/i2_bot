// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

// AddPokememesWealthUp prepares database for latest game update in mid-October
func AddPokememesWealthUp(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `profiles` ADD COLUMN `pokememes_wealth` INT(11) NOT NULL DEFAULT 0 COMMENT 'Стоимость покемонов на руках' AFTER `wealth`;")
	if err != nil {
		return err
	}

	return nil
}

// AddPokememesWealthDown  restores mid-October behaviour
func AddPokememesWealthDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `profiles` DROP COLUMN `pokememes_wealth`;")
	if err != nil {
		return err
	}

	return nil
}
