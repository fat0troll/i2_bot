// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateChatsUp creates `chats` table
func CreateChatsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `chats` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID чата',"
	request += "`name` varchar(191) NOT NULL COMMENT 'Имя чата',"
	request += "`chat_type` bool NOT NULL COMMENT 'Тип чата',"
	request += "`telegram_id` int(11) NOT NULL COMMENT 'ID чата в Телеграме',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `chats_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Чаты';"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

// CreateChatsDown drops `chats` table
func CreateChatsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `chats`;")
	if err != nil {
		return err
	}
	return nil
}
