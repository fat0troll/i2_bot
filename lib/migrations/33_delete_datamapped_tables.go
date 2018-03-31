// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// DeleteDataMappedTablesUp drops `locations`, `elements`, `weapons` and `leagues` tables
// These tables data is rarely changed, so I decided to hold such data in source code
func DeleteDataMappedTablesUp(tx *sql.Tx) error {
	request := "DROP TABLE IF EXISTS `elements`"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	request = "DROP TABLE IF EXISTS `leagues`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "DROP TABLE IF EXISTS `locations`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "DROP TABLE IF EXISTS `weapons`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDataMappedTablesDown returns old tables with all data
func DeleteDataMappedTablesDown(tx *sql.Tx) error {
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
	_, err = tx.Exec("INSERT INTO `locations` VALUES(NULL, '🌲', 'Лес', NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `locations` VALUES(NULL, '⛰', 'Горы', NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `locations` VALUES(NULL, '🚣', 'Озеро', NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `locations` VALUES(NULL, '🏙', 'Город', NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `locations` VALUES(NULL, '🏛', 'Катакомбы', NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `locations` VALUES(NULL, '⛪', 'Кладбище', NOW());")
	if err != nil {
		return err
	}

	request = "CREATE TABLE `elements` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID элемента',"
	request += "`symbol` varchar(191) COLLATE 'utf8mb4_unicode_520_ci' NOT NULL COMMENT 'Символ элемента',"
	request += "`name` varchar(191) NOT NULL COMMENT 'Имя элемента',"
	request += "`league_id` int(11) NOT NULL COMMENT 'ID родной лиги',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `elements_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Элементы';"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	// Insert elements
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '👊', 'Боевой', 1, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '🌀', 'Летающий', 1, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '💀', 'Ядовитый', 1, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '🗿', 'Каменный', 1, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '🔥', 'Огненный', 2, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '⚡', 'Электрический', 2, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '💧', 'Водяной', 2, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '🍀', 'Травяной', 2, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '💩', 'Шоколадный', 3, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '👁', 'Психический', 3, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '👿', 'Темный', 3, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `elements` VALUES(NULL, '⌛', 'Времени', 3, NOW());")
	if err != nil {
		return err
	}

	request = "CREATE TABLE `leagues` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID лиги',"
	request += "`symbol` varchar(191) COLLATE 'utf8mb4_unicode_520_ci' NOT NULL COMMENT 'Символ лиги',"
	request += "`name` varchar(191) NOT NULL COMMENT 'Имя лиги',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлена в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `leagues_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Лиги';"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	// Insert locations
	_, err = tx.Exec("INSERT INTO `leagues` VALUES(NULL, '🈸', 'ИНСТИНКТ', NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `leagues` VALUES(NULL, '🈳', 'ОТВАГА', NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `leagues` VALUES(NULL, '🈵', 'МИСТИКА', NOW());")
	if err != nil {
		return err
	}

	request = "CREATE TABLE `weapons` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID оружия',"
	request += "`name` varchar(191) NOT NULL COMMENT 'Название оружия',"
	request += "`power` int(11) NOT NULL COMMENT 'Атака оружия',"
	request += "`price` int(11) NOT NULL COMMENT 'Цена в магазине',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `weapons_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Оружие';"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO `weapons` VALUES(NULL, 'Бита', 2, 5, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `weapons` VALUES(NULL, 'Стальная бита', 10, 40, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `weapons` VALUES(NULL, 'Чугунная бита ', 200, 500, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `weapons` VALUES(NULL, 'Титановая бита', 2000, 10000, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `weapons` VALUES(NULL, 'Алмазная бита', 10000, 100000, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `weapons` VALUES(NULL, 'Криптонитовая бита', 100000, 500000, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `weapons` VALUES(NULL, 'Буханка из пятёры', 1000000, 5000000, NOW());")
	if err != nil {
		return err
	}

	return nil
}
