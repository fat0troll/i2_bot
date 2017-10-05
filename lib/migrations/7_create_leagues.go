// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
    // stdlib
    "database/sql"
)

func CreateLeaguesUp(tx *sql.Tx) error {
    create_request := "CREATE TABLE `leagues` ("
    create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID лиги',"
    create_request += "`symbol` varchar(191) COLLATE 'utf8mb4_unicode_520_ci' NOT NULL COMMENT 'Символ лиги',"
    create_request += "`name` varchar(191) NOT NULL COMMENT 'Имя лиги',"
    create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлена в базу',"
    create_request += "PRIMARY KEY (`id`),"
    create_request += "UNIQUE KEY `id` (`id`),"
    create_request += "KEY `leagues_created_at` (`created_at`)"
    create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Лиги';"
    _, err := tx.Exec(create_request)
    if err != nil {
        return err
    }

    // Insert locations
    _, err2 := tx.Exec("INSERT INTO `leagues` VALUES(NULL, ':u7533:', 'ИНСТИНКТ', NOW());")
    if err2 != nil {
        return err2
    }
    _, err3 := tx.Exec("INSERT INTO `leagues` VALUES(NULL, ':u6e80', 'ОТВАГА', NOW());")
    if err3 != nil {
        return err2
    }
    _, err4 := tx.Exec("INSERT INTO `leagues` VALUES(NULL, ':u7a7a:', 'МИСТИКА', NOW());")
    if err4 != nil {
        return err2
    }

    return nil
}

func CreateLeaguesDown(tx *sql.Tx) error {
    _, err := tx.Exec("DROP TABLE `leagues`;")
    if err != nil {
        return err
    }
    return nil
}
