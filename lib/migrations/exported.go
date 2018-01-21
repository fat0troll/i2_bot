// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/appcontext"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/migrations/migrationsinterface"
)

// Migrations handles all functions of migrations package
type Migrations struct{}

var (
	c *appcontext.Context
)

// New is an initialization function for migrations package
func New(ac *appcontext.Context) {
	c = ac
	m := &Migrations{}
	c.RegisterMigrationsInterface(migrationsinterface.MigrationsInterface(m))
}
