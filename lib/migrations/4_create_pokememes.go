// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreatePokememesUp creates `pokememes` table
func CreatePokememesUp(tx *sql.Tx) error {
	request := "CREATE TABLE `pokememes` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID покемема',"
	request += "`grade` int(11) NOT NULL COMMENT 'Поколение покемема',"
	request += "`name` varchar(191) NOT NULL COMMENT 'Имя покемема',"
	request += "`description` TEXT NOT NULL COMMENT 'Описание покемема',"
	request += "`attack` int(11) NOT NULL COMMENT 'Атака',"
	request += "`hp` int(11) NOT NULL COMMENT 'Здоровье',"
	request += "`mp` int(11) NOT NULL COMMENT 'МР',"
	request += "`defence` int(11) NOT NULL COMMENT 'Защита',"
	request += "`price` int(11) NOT NULL COMMENT 'Стоимость',"
	request += "`purchaseable` bool NOT NULL DEFAULT true COMMENT 'Можно купить?',"
	request += "`image_url` varchar(191) NOT NULL COMMENT 'Изображение покемема',"
	request += "`player_id` int(11) NOT NULL COMMENT 'Кто добавил в базу',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `pokememes_created_at` (`created_at`),"
	request += "KEY `pokememes_player_id` (`player_id`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Покемемы'"

	_, err := tx.Exec(request)

	return err
}

// CreatePokememesDown drops `pokememes` table
func CreatePokememesDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `pokememes`")

	return err
}
