// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateElementsUp creates `elements` table and fills it with data
func CreateElementsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `elements` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID элемента',"
	request += "`symbol` varchar(191) COLLATE 'utf8mb4_unicode_520_ci' NOT NULL COMMENT 'Символ элемента',"
	request += "`name` varchar(191) NOT NULL COMMENT 'Имя элемента',"
	request += "`league_id` int(11) NOT NULL COMMENT 'ID родной лиги',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `elements_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Элементы';"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	// Insert elements
	_, err2 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '👊', 'Боевой', 1, NOW());")
	if err2 != nil {
		return err2
	}
	_, err3 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '🌀', 'Летающий', 1, NOW());")
	if err3 != nil {
		return err3
	}
	_, err4 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '💀', 'Ядовитый', 1, NOW());")
	if err4 != nil {
		return err4
	}
	_, err5 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '🗿', 'Каменный', 1, NOW());")
	if err5 != nil {
		return err5
	}
	_, err6 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '🔥', 'Огненный', 2, NOW());")
	if err6 != nil {
		return err6
	}
	_, err7 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '⚡', 'Электрический', 2, NOW());")
	if err7 != nil {
		return err7
	}
	_, err8 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '💧', 'Водяной', 2, NOW());")
	if err8 != nil {
		return err8
	}
	_, err9 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '🍀', 'Травяной', 2, NOW());")
	if err9 != nil {
		return err9
	}
	_, err10 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '💩', 'Шоколадный', 3, NOW());")
	if err10 != nil {
		return err10
	}
	_, err11 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '👁', 'Психический', 3, NOW());")
	if err11 != nil {
		return err11
	}
	_, err12 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '👿', 'Темный', 3, NOW());")
	if err12 != nil {
		return err12
	}
	_, err13 := tx.Exec("INSERT INTO `elements` VALUES(NULL, '⌛', 'Времени', 1, NOW());")
	if err13 != nil {
		return err13
	}

	return nil
}

// CreateElementsDown drops `elements` table
func CreateElementsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `elements`;")
	if err != nil {
		return err
	}
	return nil
}
