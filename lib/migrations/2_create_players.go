// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreatePlayersUp creates `players` table
func CreatePlayersUp(tx *sql.Tx) error {
	request := "CREATE TABLE `players` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID игрока',"
	request += "`telegram_id` int(11) NOT NULL COMMENT 'ID в телеграме',"
	request += "`league_id` int(11) COMMENT 'ID лиги' DEFAULT 0,"
	request += "`squad_id` int(11) COMMENT 'ID отряда' DEFAULT 0,"
	request += "`status` varchar(191) COMMENT 'Статус в лиге' DEFAULT 'common',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Время последнего обновления',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `players_created_at` (`created_at`),"
	request += "KEY `players_updated_at` (`updated_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Зарегистрированные игроки'"

	_, err := tx.Exec(request)

	return err
}

// CreatePlayersDown drops `players` table
func CreatePlayersDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `players`;")

	return err
}
