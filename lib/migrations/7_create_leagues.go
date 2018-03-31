// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateLeaguesUp creates `leagues` table and fills it with data
func CreateLeaguesUp(tx *sql.Tx) error {
	request := "CREATE TABLE `leagues` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID лиги',"
	request += "`symbol` varchar(191) COLLATE 'utf8mb4_unicode_520_ci' NOT NULL COMMENT 'Символ лиги',"
	request += "`name` varchar(191) NOT NULL COMMENT 'Имя лиги',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлена в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `leagues_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Лиги';"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	// Insert locations
	_, err = tx.Exec("INSERT INTO `leagues` VALUES(NULL, ':u7533:', 'ИНСТИНКТ', NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `leagues` VALUES(NULL, ':u6e80', 'ОТВАГА', NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `leagues` VALUES(NULL, ':u7a7a:', 'МИСТИКА', NOW());")
	if err != nil {
		return err
	}

	return nil
}

// CreateLeaguesDown drops `leagues` table
func CreateLeaguesDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `leagues`;")
	if err != nil {
		return err
	}
	return nil
}
