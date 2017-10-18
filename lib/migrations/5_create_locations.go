// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func CreateLocationsUp(tx *sql.Tx) error {
	create_request := "CREATE TABLE `locations` ("
	create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID локации',"
	create_request += "`symbol` varchar(191) COLLATE 'utf8mb4_unicode_520_ci' NOT NULL COMMENT 'Символ локации',"
	create_request += "`name` varchar(191) NOT NULL COMMENT 'Имя локации',"
	create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлена в базу',"
	create_request += "PRIMARY KEY (`id`),"
	create_request += "UNIQUE KEY `id` (`id`),"
	create_request += "KEY `locations_created_at` (`created_at`)"
	create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Локации';"
	_, err := tx.Exec(create_request)
	if err != nil {
		return err
	}

	// Insert locations
	_, err2 := tx.Exec("INSERT INTO `locations` VALUES(NULL, ':evergreen_tree:', 'Лес', NOW());")
	if err2 != nil {
		return err2
	}
	_, err3 := tx.Exec("INSERT INTO `locations` VALUES(NULL, '⛰', 'Горы', NOW());")
	if err3 != nil {
		return err2
	}
	_, err4 := tx.Exec("INSERT INTO `locations` VALUES(NULL, ':rowboat:', 'Озеро', NOW());")
	if err4 != nil {
		return err2
	}
	_, err5 := tx.Exec("INSERT INTO `locations` VALUES(NULL, '🏙:', 'Город', NOW());")
	if err5 != nil {
		return err2
	}
	_, err6 := tx.Exec("INSERT INTO `locations` VALUES(NULL, '🏛', 'Катакомбы', NOW());")
	if err6 != nil {
		return err2
	}
	_, err7 := tx.Exec("INSERT INTO `locations` VALUES(NULL, ':church:', 'Кладбище', NOW());")
	if err7 != nil {
		return err2
	}

	return nil
}

func CreateLocationsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `locations`;")
	if err != nil {
		return err
	}
	return nil
}
