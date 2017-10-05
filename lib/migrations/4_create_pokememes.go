// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
    // stdlib
    "database/sql"
)

func CreatePokememesUp(tx *sql.Tx) error {
    create_request := "CREATE TABLE `pokememes` ("
    create_request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID покемема',"
    create_request += "`grade` int(11) NOT NULL COMMENT 'Поколение покемема',"
    create_request += "`name` varchar(191) NOT NULL COMMENT 'Имя покемема',"
    create_request += "`description` TEXT NOT NULL COMMENT 'Описание покемема',"
    create_request += "`attack` int(11) NOT NULL COMMENT 'Атака',"
    create_request += "`hp` int(11) NOT NULL COMMENT 'Здоровье',"
    create_request += "`mp` int(11) NOT NULL COMMENT 'МР',"
    create_request += "`defence` int(11) NOT NULL COMMENT 'Защита',"
    create_request += "`price` int(11) NOT NULL COMMENT 'Стоимость',"
    create_request += "`purchaseable` bool NOT NULL DEFAULT true COMMENT 'Можно купить?',"
    create_request += "`image_url` varchar(191) NOT NULL COMMENT 'Изображение покемема',"
    create_request += "`player_id` int(11) NOT NULL COMMENT 'Кто добавил в базу',"
    create_request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
    create_request += "PRIMARY KEY (`id`),"
    create_request += "UNIQUE KEY `id` (`id`),"
    create_request += "KEY `pokememes_created_at` (`created_at`),"
    create_request += "KEY `pokememes_player_id` (`player_id`)"
    create_request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Покемемы';"
    _, err := tx.Exec(create_request)
    if err != nil {
        return err
    }
    return nil
}

func CreatePokememesDown(tx *sql.Tx) error {
    _, err := tx.Exec("DROP TABLE `pokememes`;")
    if err != nil {
        return err
    }
    return nil
}
