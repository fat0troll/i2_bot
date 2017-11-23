// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

// AddUserTypeUp creates `user_type` column in `squads_players` table
func AddUserTypeUp(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `squads_players` ADD COLUMN `user_type` varchar(191) NOT NULL DEFAULT 'common' COMMENT 'Уровень игрока' AFTER `player_id`;")
	if err != nil {
		return err
	}

	return nil
}

// AddUserTypeDown  destroys `user_type` column
func AddUserTypeDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `squads_players` DROP COLUMN `user_type`;")
	if err != nil {
		return err
	}

	return nil
}
