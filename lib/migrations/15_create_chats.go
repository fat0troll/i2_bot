// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func CreateChatsUp(tx *sql.Tx) error {
	create_request := "CREATE TABLE `chats` ("
	create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID чата',"
	create_request += "`name` varchar(191) NOT NULL COMMENT 'Имя чата',"
	create_request += "`chat_type` bool NOT NULL COMMENT 'Тип чата',"
	create_request += "`telegram_id` int(11) NOT NULL COMMENT 'ID чата в Телеграме',"
	create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	create_request += "PRIMARY KEY (`id`),"
	create_request += "UNIQUE KEY `id` (`id`),"
	create_request += "KEY `chats_created_at` (`created_at`)"
	create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Чаты';"
	_, err := tx.Exec(create_request)
	if err != nil {
		return err
	}

	return nil
}

func CreateChatsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `chats`;")
	if err != nil {
		return err
	}
	return nil
}
