// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// ChangeSquadsTableUp changes `sqauds` to new format
func ChangeSquadsTableUp(tx *sql.Tx) error {
	request := "ALTER TABLE `squads` DROP COLUMN `flood_chat_id`"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	request = "ALTER TABLE `squads` DROP COLUMN `author_id`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "ALTER TABLE `squads` ADD COLUMN `min_level` int(11) NOT NULL DEFAULT 0 AFTER `chat_id`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "ALTER TABLE `squads` ADD COLUMN `max_level` int(11) NOT NULL DEFAULT 0 AFTER `min_level`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "ALTER TABLE `squads` ADD COLUMN `invite_link` varchar(191) NOT NULL DEFAULT 'https://example.com' AFTER `max_level`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

// ChangeSquadsTableDown reverts `squads` to old format
func ChangeSquadsTableDown(tx *sql.Tx) error {
	request := "ALTER TABLE `squads` ADD COLUMN `flood_chat_id` int(11) NOT NULL AFTER `chat_id`"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	request = "ALTER TABLE `squads` ADD COLUMN `author_id` int(11) NOT NULL AFTER `flood_chat_id"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "ALTER TABLE `squads` DROP COLUMN `min_level`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "ALTER TABLE `squads` DROP COLUMN `max_level`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "ALTER TABLE `squads` DROP COLUMN `invite_link`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}
