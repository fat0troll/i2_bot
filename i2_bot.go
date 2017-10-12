// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package main

import (
	// stdlib
	"time"
    // 3rd-party
	"github.com/go-telegram-bot-api/telegram-bot-api"
    // local
	"./lib/appcontext"
	"./lib/getters"
	"./lib/migrations"
	"./lib/parsers"
	"./lib/router"
	"./lib/talkers"
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
	parsers.New(c)
	talkers.New(c)
	getters.New(c)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := c.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.From == nil {
			continue
		} else if update.Message.Date < (int(time.Now().Unix()) - 1) {
			// Ignore old messages
			continue
		}

        c.Router.RouteRequest(update)
	}
}
