// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func CreateWeaponsAndAddWealthUp(tx *sql.Tx) error {
	create_request := "CREATE TABLE `weapons` ("
	create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID оружия',"
	create_request += "`name` varchar(191) NOT NULL COMMENT 'Название оружия',"
	create_request += "`power` int(11) NOT NULL COMMENT 'Атака оружия',"
	create_request += "`price` int(11) NOT NULL COMMENT 'Цена в магазине',"
	create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	create_request += "PRIMARY KEY (`id`),"
	create_request += "UNIQUE KEY `id` (`id`),"
	create_request += "KEY `weapons_created_at` (`created_at`)"
	create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Оружие';"
	_, err := tx.Exec(create_request)
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

	_, err = tx.Exec("ALTER TABLE `profiles` ADD COLUMN `wealth` INT(11) NOT NULL COMMENT 'Денег на руках' AFTER `pokeballs`;")
	if err != nil {
		return err
	}

	return nil
}

func CreateWeaponsAndAddWealthDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `profiles` DROP COLUMN `wealth`;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE `weapons`;")
	if err != nil {
		return err
	}

	return nil
}
