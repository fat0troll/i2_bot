// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateLocationsUp creates `locations` table and fills it with data
func CreateLocationsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `locations` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID локации',"
	request += "`symbol` varchar(191) COLLATE 'utf8mb4_unicode_520_ci' NOT NULL COMMENT 'Символ локации',"
	request += "`name` varchar(191) NOT NULL COMMENT 'Имя локации',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлена в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `locations_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Локации';"
	_, err := tx.Exec(request)
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
	_, err5 := tx.Exec("INSERT INTO `locations` VALUES(NULL, '🏙', 'Город', NOW());")
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

// CreateLocationsDown drops `locations` table
func CreateLocationsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `locations`;")
	if err != nil {
		return err
	}
	return nil
}
