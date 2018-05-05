// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateTournamentReportsUp creates `tournament_reports` table
func CreateTournamentReportsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `tournament_reports` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID рапорта',"
	request += "`player_id` int(11) NOT NULL COMMENT 'Игрок, который сходил на турнир',"
	request += "`tournament_number` int(11) NOT NULL COMMENT 'Номер турнира в игре',"
	request += "`target` varchar(191) NOT NULL COMMENT 'Цель атаки',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `tournament_reports_player_id` (`player_id`),"
	request += "KEY `tournament_repots_tournament_number` (`tournament_number`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Репорты результата Турнира лиг'"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

// CreateTournamentReportsDown drops `tournament_reports` table
func CreateTournamentReportsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `tournament_reports`")
	if err != nil {
		return err
	}

	return nil
}
