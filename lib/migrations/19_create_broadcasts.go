// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func CreateBroadcastsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `broadcasts` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID сообщения',"
	request += "`text` text NOT NULL COMMENT 'Тело сообщения',"
	request += "`broadcast_type` varchar(191) NOT NULL COMMENT 'Тип широковещательного сообщения',"
	request += "`status` varchar(191) NOT NULL DEFAULT 'new' COMMENT 'Статус сообщения',"
	request += "`author_id` int(11) NOT NULL COMMENT 'ID автора',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `broadcasts_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Широковещательные сообщения';"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

func CreateBroadcastsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `broadcasts`;")
	if err != nil {
		return err
	}
	return nil
}
