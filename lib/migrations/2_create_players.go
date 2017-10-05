// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
    // stdlib
    "database/sql"
)

func CreatePlayersUp(tx *sql.Tx) error {
    create_request := "CREATE TABLE `players` ("
    create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID игрока',"
    create_request += "`telegram_id` int(11) NOT NULL COMMENT 'ID в телеграме',"
    create_request += "`league_id` int(11) COMMENT 'ID лиги' DEFAULT 0,"
    create_request += "`squad_id` int(11) COMMENT 'ID отряда' DEFAULT 0,"
    create_request += "`status` varchar(191) COMMENT 'Статус в лиге' DEFAULT 'common',"
    create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
    create_request += "`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Время последнего обновления',"
    create_request += "PRIMARY KEY (`id`),"
    create_request += "UNIQUE KEY `id` (`id`),"
    create_request += "KEY `players_created_at` (`created_at`),"
    create_request += "KEY `players_updated_at` (`updated_at`)"
    create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Зарегистрированные игроки';"
    _, err := tx.Exec(create_request)
    if err != nil {
        return err
    }
    return nil
}

func CreatePlayersDown(tx *sql.Tx) error {
    _, err := tx.Exec("DROP TABLE `players`;")
    if err != nil {
        return err
    }
    return nil
}
