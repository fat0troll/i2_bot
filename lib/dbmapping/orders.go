// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package dbmapping

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

// Order is a struct, which represents `orders` table item in databse.
type Order struct {
	ID           int            `db:"id"`
	Target       string         `db:"target"`
	TargetSquads string         `db:"target_squads"`
	Scheduled    bool           `db:"scheduled"`
	ScheduledAt  mysql.NullTime `db:"scheduled_at"`
	Status       string         `db:"status"`
	AuthorID     int            `db:"author_id"`
	CreatedAt    time.Time      `db:"created_at"`
}
