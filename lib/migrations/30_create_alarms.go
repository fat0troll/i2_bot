// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateAlarmsUp creates `alarms` table
func CreateAlarmsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `alarms` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID настройки оповещения',"
	request += "`player_id` int(11) NOT NULL COMMENT 'Игрок, которому отправляется оповещение',"
	request += "`turnir_number` int(11) NOT NULL COMMENT 'Номер турнира (от 1 до 12)',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `alarms_player_id` (`player_id`),"
	request += "KEY `alarms_turnir_number` (`turnir_number`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Оповещения на Турнир лиг'"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

// CreateAlarmsDown drops `alarms` table
func CreateAlarmsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `alarms`")
	if err != nil {
		return err
	}

	return nil
}
