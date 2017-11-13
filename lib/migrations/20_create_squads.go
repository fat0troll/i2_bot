// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateSquadsUp creates `squads` table
func CreateSquadsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `squads` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID отряда',"
	request += "`chat_id` int(11) NOT NULL COMMENT 'ID чата в базе',"
	request += "`author_id` int(11) NOT NULL COMMENT 'ID автора отряда',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `squads_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Отряды';"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	request = "CREATE TABLE `squads_players` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID связи',"
	request += "`squad_id` int(11) NOT NULL COMMENT 'ID отряда в базе',"
	request += "`player_id` int(11) NOT NULL COMMENT 'ID игрока',"
	request += "`author_id` int(11) NOT NULL COMMENT 'ID добавившего в отряд',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлено в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `squads_players_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Связь Отряды-Игроки';"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	request = "ALTER TABLE `players` DROP COLUMN `squad_id`"
	_, err = tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

// CreateSquadsDown drops `squads` table
func CreateSquadsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `squads`;")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE `squads_players`;")
	if err != nil {
		return err
	}
	_, err = tx.Exec("MODIFY TABLE `players` ADD COLUMN `league_id` int(11) COMMENT 'ID лиги' DEFAULT 0")
	if err != nil {
		return err
	}

	return nil
}
