// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func CreateProfilesUp(tx *sql.Tx) error {
	create_request := "CREATE TABLE `profiles` ("
	create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID сохраненного профиля',"
	create_request += "`player_id` int(11) NOT NULL COMMENT 'ID игрока в системе',"
	create_request += "`nickname` varchar(191) NOT NULL COMMENT 'Ник игрока',"
	create_request += "`telegram_nickname` varchar(191) NOT NULL COMMENT 'Ник в Телеграме (@)',"
	create_request += "`level_id` int(11) NOT NULL COMMENT 'Уровень',"
	create_request += "`exp` int(11) NOT NULL COMMENT 'Опыт',"
	create_request += "`egg_exp` int(11) NOT NULL COMMENT 'Опыт яйца',"
	create_request += "`power` int(11) NOT NULL COMMENT 'Сила без оружия',"
	create_request += "`weapon_id` int(11) NOT NULL COMMENT 'Тип оружия',"
	create_request += "`crystalls` int(11) NOT NULL COMMENT 'Кристаллы',"
	create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	create_request += "PRIMARY KEY (`id`),"
	create_request += "UNIQUE KEY `id` (`id`),"
	create_request += "KEY `profiles_created_at` (`created_at`),"
	create_request += "KEY `profiles_nickname` (`nickname`)"
	create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Профили зарегистрированных игроков';"
	_, err := tx.Exec(create_request)
	if err != nil {
		return err
	}
	return nil
}

func CreateProfilesDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `profiles`;")
	if err != nil {
		return err
	}
	return nil
}
