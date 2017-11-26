// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/appcontext"
	"lab.pztrn.name/fat0troll/i2_bot/lib/broadcaster"
	"lab.pztrn.name/fat0troll/i2_bot/lib/chatter"
	"lab.pztrn.name/fat0troll/i2_bot/lib/forwarder"
	"lab.pztrn.name/fat0troll/i2_bot/lib/migrations"
	"lab.pztrn.name/fat0troll/i2_bot/lib/orders"
	"lab.pztrn.name/fat0troll/i2_bot/lib/pinner"
	"lab.pztrn.name/fat0troll/i2_bot/lib/pokedexer"
	"lab.pztrn.name/fat0troll/i2_bot/lib/router"
	"lab.pztrn.name/fat0troll/i2_bot/lib/squader"
	"lab.pztrn.name/fat0troll/i2_bot/lib/statistics"
	"lab.pztrn.name/fat0troll/i2_bot/lib/talkers"
	"lab.pztrn.name/fat0troll/i2_bot/lib/users"
	"lab.pztrn.name/fat0troll/i2_bot/lib/welcomer"
	"time"
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

	c.Log.Info("=======================")
	c.Log.Info("= i2_bot initialized. =")
	c.Log.Info("=======================")

	c.Cron.Start()
	c.Log.Info("> Cron started.")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := c.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.From != nil {
				if update.Message.Date > (int(time.Now().Unix()) - 5) {
					c.Router.RouteRequest(&update)
				}
			}
		} else if update.InlineQuery != nil {
			c.Router.RouteInline(&update)
		} else if update.ChosenInlineResult != nil {
			c.Log.Debug(update.ChosenInlineResult.ResultID)
		} else {
			continue
		}
	}
}
