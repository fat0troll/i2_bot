// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// ChangeProfilePokememesColumnsUp prepares database for latest game update in mid-October
func ChangeProfilePokememesColumnsUp(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `profiles_pokememes` ADD COLUMN `pokememe_attack` INT(11) NOT NULL DEFAULT 0 COMMENT 'Атака покемема' AFTER `pokememe_id`;;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE `profiles_pokememes` DROP COLUMN `pokememe_lvl`;")
	if err != nil {
		return err
	}

	return nil
}

// ChangeProfilePokememesColumnsDown restores mid-October behaviour
func ChangeProfilePokememesColumnsDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `profiles_pokememes` ADD COLUMN `pokememe_lvl` INT(11) NOT NULL DEFAULT 0 COMMENT 'Уровень покемема' AFTER `pokememe_id`;;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE `profiles_pokememes` DROP COLUMN `pokememe_attack`;")
	if err != nil {
		return err
	}

	return nil
}
