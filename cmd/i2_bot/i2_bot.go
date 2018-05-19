// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package main

import (
	"github.com/fat0troll/i2_bot/lib/appcontext"
	"github.com/fat0troll/i2_bot/lib/broadcaster"
	"github.com/fat0troll/i2_bot/lib/chatter"
	"github.com/fat0troll/i2_bot/lib/datacache"
	"github.com/fat0troll/i2_bot/lib/forwarder"
	"github.com/fat0troll/i2_bot/lib/migrations"
	"github.com/fat0troll/i2_bot/lib/orders"
	"github.com/fat0troll/i2_bot/lib/pinner"
	"github.com/fat0troll/i2_bot/lib/pokedexer"
	"github.com/fat0troll/i2_bot/lib/reminder"
	"github.com/fat0troll/i2_bot/lib/router"
	"github.com/fat0troll/i2_bot/lib/sender"
	"github.com/fat0troll/i2_bot/lib/squader"
	"github.com/fat0troll/i2_bot/lib/statistics"
	"github.com/fat0troll/i2_bot/lib/talkers"
	"github.com/fat0troll/i2_bot/lib/users"
	"github.com/fat0troll/i2_bot/lib/welcomer"
)

var (
	c *appcontext.Context
)

func main() {
	c := appcontext.New()
	c.Init()

	router.New(c)
	migrations.New(c)
	c.RunDatabaseMigrations()
	datacache.New(c)
	sender.New(c)
	forwarder.New(c)
	pokedexer.New(c)
	pinner.New(c)
	talkers.New(c)
	broadcaster.New(c)
	welcomer.New(c)
	chatter.New(c)
	squader.New(c)
	users.New(c)
	statistics.New(c)
	orders.New(c)
	reminder.New(c)

	c.Log.Info("=======================")
	c.Log.Info("= i2_bot initialized. =")
	c.Log.Info("=======================")

	c.Cron.Start()
	c.Log.Info("> Cron started.")

	c.StartBot()
}
