// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

// RemoveReusableUp removes `reusable` field in `orders` table
func RemoveReusableUp(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `orders` DROP COLUMN `reusable`")
	if err != nil {
		return err
	}

	return nil
}

// RemoveReusableDown reverts `reusable` column
func RemoveReusableDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `orders` ADD COLUMN `reusable` bool NOT NULL DEFAULT true COMMENT 'Можно ли повторить приказ' AFTER `scheduled_at")
	if err != nil {
		return err
	}

	return nil
}
