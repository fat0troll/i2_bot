// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateProfilesUp creates `profiles` table
func CreateProfilesUp(tx *sql.Tx) error {
	request := "CREATE TABLE `profiles` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID сохраненного профиля',"
	request += "`player_id` int(11) NOT NULL COMMENT 'ID игрока в системе',"
	request += "`nickname` varchar(191) NOT NULL COMMENT 'Ник игрока',"
	request += "`telegram_nickname` varchar(191) NOT NULL COMMENT 'Ник в Телеграме (@)',"
	request += "`level_id` int(11) NOT NULL COMMENT 'Уровень',"
	request += "`exp` int(11) NOT NULL COMMENT 'Опыт',"
	request += "`egg_exp` int(11) NOT NULL COMMENT 'Опыт яйца',"
	request += "`power` int(11) NOT NULL COMMENT 'Сила без оружия',"
	request += "`weapon_id` int(11) NOT NULL COMMENT 'Тип оружия',"
	request += "`crystalls` int(11) NOT NULL COMMENT 'Кристаллы',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `profiles_created_at` (`created_at`),"
	request += "KEY `profiles_nickname` (`nickname`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Профили зарегистрированных игроков';"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}
	return nil
}

// CreateProfilesDown drops `profiles` table
func CreateProfilesDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `profiles`;")
	if err != nil {
		return err
	}
	return nil
}
