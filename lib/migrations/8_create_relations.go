// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateRelationsUp creates tables for pokememes-locations and pokememes-elements links
func CreateRelationsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `pokememes_locations` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID связи',"
	request += "`pokememe_id` int(11) NOT NULL COMMENT 'ID покемема',"
	request += "`location_id` int(11) NOT NULL COMMENT 'ID локации',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `pokememes_locations_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Связь Покемемы-Локации';"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	request = "CREATE TABLE `pokememes_elements` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID связи',"
	request += "`pokememe_id` int(11) NOT NULL COMMENT 'ID покемема',"
	request += "`element_id` int(11) NOT NULL COMMENT 'ID элемента',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `pokememes_elements_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Связь Покемемы-Элементы';"
	_, err2 := tx.Exec(request)
	if err2 != nil {
		return err2
	}
	return nil
}

// CreateRelationsDown drops pokememe-* relations tables
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
