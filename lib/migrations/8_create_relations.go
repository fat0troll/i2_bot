// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func CreateRelationsUp(tx *sql.Tx) error {
	create_request := "CREATE TABLE `pokememes_locations` ("
	create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID связи',"
	create_request += "`pokememe_id` int(11) NOT NULL COMMENT 'ID покемема',"
	create_request += "`location_id` int(11) NOT NULL COMMENT 'ID локации',"
	create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	create_request += "PRIMARY KEY (`id`),"
	create_request += "UNIQUE KEY `id` (`id`),"
	create_request += "KEY `pokememes_locations_created_at` (`created_at`)"
	create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Связь Покемемы-Локации';"
	_, err := tx.Exec(create_request)
	if err != nil {
		return err
	}

	create_request = "CREATE TABLE `pokememes_elements` ("
	create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID связи',"
	create_request += "`pokememe_id` int(11) NOT NULL COMMENT 'ID покемема',"
	create_request += "`element_id` int(11) NOT NULL COMMENT 'ID элемента',"
	create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	create_request += "PRIMARY KEY (`id`),"
	create_request += "UNIQUE KEY `id` (`id`),"
	create_request += "KEY `pokememes_elements_created_at` (`created_at`)"
	create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Связь Покемемы-Элементы';"
	_, err2 := tx.Exec(create_request)
	if err2 != nil {
		return err2
	}
	return nil
}

func CreateRelationsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `pokememes_locations`;")
	if err != nil {
		return err
	}
	_, err2 := tx.Exec("DROP TABLE `pokememes_elements`;")
	if err2 != nil {
		return err2
	}
	return nil
}
