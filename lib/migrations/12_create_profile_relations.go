// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func CreateProfileRelationsUp(tx *sql.Tx) error {
	create_request := "CREATE TABLE `profiles_pokememes` ("
	create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID связи',"
	create_request += "`profile_id` int(11) NOT NULL COMMENT 'ID профиля',"
	create_request += "`pokememe_id` int(11) NOT NULL COMMENT 'ID покемема',"
	create_request += "`pokememe_lvl` int(11) NOT NULL COMMENT 'Уровень покемема',"
	create_request += "`pokememe_rarity` varchar(191) NOT NULL DEFAULT 'common' COMMENT 'Редкость покемема',"
	create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	create_request += "PRIMARY KEY (`id`),"
	create_request += "UNIQUE KEY `id` (`id`),"
	create_request += "KEY `profiles_pokememes_created_at` (`created_at`)"
	create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Связь Профили-Покемемы';"
	_, err := tx.Exec(create_request)
	if err != nil {
		return err
	}
	return nil
}

func CreateProfileRelationsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `profiles_pokememes`;")
	if err != nil {
		return err
	}
	return nil
}
