// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package main

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/appcontext"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/broadcaster"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/chatter"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datacache"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/forwarder"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/migrations"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/orders"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/pinner"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/pokedexer"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/reminder"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/router"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/squader"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/statistics"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/talkers"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/users"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/welcomer"
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
