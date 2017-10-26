// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/appcontext"
	"lab.pztrn.name/fat0troll/i2_bot/lib/migrations/migrationsinterface"
)

var (
	c *appcontext.Context
)

func New(ac *appcontext.Context) {
	c = ac
	m := &Migrations{}
	c.RegisterMigrationsInterface(migrationsinterface.MigrationsInterface(m))
}
