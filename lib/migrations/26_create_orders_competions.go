// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// CreateOrdersCompletionsUp creates `orders_completions` table
func CreateOrdersCompletionsUp(tx *sql.Tx) error {
	request := "CREATE TABLE `orders_completions` ("
	request += "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID приказа',"
	request += "`order_id` int(11) NOT NULL COMMENT 'Выполненный приказ',"
	request += "`executor_id` int(11) NOT NULL COMMENT 'ID выполнившего приказ',"
	request += "`created_at` datetime NOT NULL COMMENT 'Добавлен в базу',"
	request += "PRIMARY KEY (`id`),"
	request += "UNIQUE KEY `id` (`id`),"
	request += "KEY `orders_completions_created_at` (`created_at`)"
	request += ") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Исполнение приказов'"
	_, err := tx.Exec(request)
	if err != nil {
		return err
	}

	return nil
}

// CreateOrdersCompletionsDown drops `orders_completions` table
func CreateOrdersCompletionsDown(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `orders_completions`")
	if err != nil {
		return err
	}
	return nil
}
