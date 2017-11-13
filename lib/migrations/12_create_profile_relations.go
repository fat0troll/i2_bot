// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateProfileRelationsUp creates profile-pokememes relation table
func CreateProfileRelationsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `profiles_pokememes` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID связи',"
	request += "`profile_id` int(11) NOT NULL COMMENT 'ID профиля',"
	request += "`pokememe_id` int(11) NOT NULL COMMENT 'ID покемема',"
	request += "`pokememe_lvl` int(11) NOT NULL COMMENT 'Уровень покемема',"
	request += "`pokememe_rarity` varchar(191) NOT NULL DEFAULT 'common' COMMENT 'Редкость покемема',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `profiles_pokememes_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Связь Профили-Покемемы';"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}
	return nil
}

// CreateProfileRelationsDown drops profile-pokememes relation table
func CreateProfileRelationsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `profiles_pokememes`;")
	if err != nil {
		return err
	}
	return nil
}
