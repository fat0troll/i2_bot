// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

// AddFloodChatIDUp creates `flood_chat_id` column in `squads` table
func AddFloodChatIDUp(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `squads` ADD COLUMN `flood_chat_id` INT(11) NOT NULL DEFAULT 0 COMMENT 'ID группы для общения отряда' AFTER `chat_id`;")
	if err != nil {
		return err
	}

	return nil
}

// AddFloodChatIDDown  destroys `flood_chat_id` column
func AddFloodChatIDDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `squads` DROP COLUMN `flood_chat_id`;")
	if err != nil {
		return err
	}

	return nil
}
