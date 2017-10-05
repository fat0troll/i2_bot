// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
    // local
    "../appcontext"
    "../migrations/migrationsinterface"
)

var (
    c *appcontext.Context
)

func New(ac *appcontext.Context) {
    c = ac
    m := &Migrations{}
    c.RegisterMigrationsInterface(migrationsinterface.MigrationsInterface(m))
}
