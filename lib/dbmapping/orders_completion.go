// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"
)

// OrderCompletion is a struct, which represents `orders_completions` table item in databse.
type OrderCompletion struct {
	ID         int       `db:"id"`
	OrderID    int       `db:"order_id"`
	ExecutorID int       `db:"executor_id"`
	CreatedAt  time.Time `db:"created_at"`
}
