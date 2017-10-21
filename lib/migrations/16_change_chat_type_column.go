// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func ChangeChatTypeColumnUp(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `chats` MODIFY `chat_type` varchar(191);")
	if err != nil {
		return err
	}

	return nil
}

func ChangeChatTypeColumnDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `chats` MODIFY `chat_type` bool;")
	if err != nil {
		return err
	}

	return nil
}
