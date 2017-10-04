// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package main

import (
	// stdlib
	"time"
    // 3rd-party
	"gopkg.in/telegram-bot-api.v4"
    // local
	"./lib/appcontext"
	"./lib/router"
)

var (
	c *appcontext.Context
)

func main() {
	c := appcontext.New()
	c.Init()
	router.New(c)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := c.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		} else if update.Message.Date < (int(time.Now().Unix()) - 1) {
			// Ignore old messages
			continue
		}

        c.Router.RouteRequest(update)
	}
}
