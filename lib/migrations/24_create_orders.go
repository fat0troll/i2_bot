// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateOrdersUp creates `orders` table
func CreateOrdersUp(tx *sql.Tx) error {
	request := "CREATE TABLE `orders` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID приказа',"
	request += "`target` varchar(191) NOT NULL COMMENT 'Цель приказа',"
	request += "`target_squads` varchar(191) NOT NULL COMMENT 'Отряды, для которых этот приказ действителен',"
	request += "`scheduled` bool NOT NULL DEFAULT false COMMENT 'Является ли запланированным',"
	request += "`scheduled_at` datetime COMMENT 'Время запланированного пина',"
	request += "`reusable` bool NOT NULL DEFAULT true COMMENT 'Можно ли повторить приказ',"
	request += "`status` varchar(191) NOT NULL DEFAULT 'new' COMMENT 'Статус приказа',"
	request += "`author_id` int(11) NOT NULL COMMENT 'ID автора приказа',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `orders_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=4201 DEFAULT CHARSET=utf8mb4 COMMENT='Приказы'"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	// Fill some default templates to send
	_, err = tx.Exec("INSERT INTO `orders` VALUES(NULL, 'M', 'all', false, NULL, true, 'new', 1, NOW())")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `orders` VALUES(NULL, 'O', 'all', false, NULL, true, 'new', 1, NOW())")
	if err != nil {
		return err
	}

	return nil
}

// CreateOrdersDown drops `chats` table
func CreateOrdersDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `orders`")
	if err != nil {
		return err
	}
	return nil
}
