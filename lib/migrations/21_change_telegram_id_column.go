// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// ChangeTelegramIDColumnUp changes `telegram_id` column of `chats` table
func ChangeTelegramIDColumnUp(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `chats` MODIFY `telegram_id` bigint;")
	if err != nil {
		return err
	}

	return nil
}

// ChangeTelegramIDColumnDown changes `telegram_id` column of `chats` table
func ChangeTelegramIDColumnDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `chats` MODIFY `telegram_id` bigint;")
	if err != nil {
		return err
	}

	return nil
}
