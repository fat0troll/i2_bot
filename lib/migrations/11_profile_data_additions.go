// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// ProfileDataAdditionsUp creates some helping databases for profiles
func ProfileDataAdditionsUp(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `profiles` ADD `pokeballs` INT(11) DEFAULT 5 NOT NULL COMMENT 'Покеболы' AFTER `level_id`;")
	if err != nil {
		return err
	}

	request := "CREATE TABLE `levels` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID уровня и его номер',"
	request += "`max_exp` int(11) NOT NULL COMMENT 'Опыт для прохождения уровня',"
	request += "`max_egg` int(11) NOT NULL COMMENT 'Опыт для открытия яйца',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `levels_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Уровни';"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	// Insert levels
	_, err = tx.Exec("INSERT INTO `levels` VALUES(NULL, 200, 6, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `levels` VALUES(NULL, 400, 12, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `levels` VALUES(NULL, 800, 24, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `levels` VALUES(NULL, 1600, 48, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `levels` VALUES(NULL, 3200, 96, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `levels` VALUES(NULL, 6400, 192, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `levels` VALUES(NULL, 12800, 384, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `levels` VALUES(NULL, 25600, 768, NOW());")
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `levels` VALUES(NULL, 51200, 1536, NOW());")
	if err != nil {
		return err
	}

	return nil
}

// ProfileDataAdditionsDown drops `levels` table and `pokeballs` column of `profiles` table
func ProfileDataAdditionsDown(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE `profiles` DROP COLUMN `pokeballs`;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE `levels`;")
	if err != nil {
		return err
	}

	return nil
}
