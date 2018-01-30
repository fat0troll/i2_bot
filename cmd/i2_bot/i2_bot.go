// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package main

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/appcontext"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/broadcaster"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/chatter"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/datacache"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/forwarder"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/migrations"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/orders"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/pinner"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/pokedexer"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/reminder"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/router"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/squader"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/statistics"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/talkers"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/users"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/welcomer"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
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

	_, err := c.Bot.SetWebhook(tgbotapi.NewWebhook(c.Cfg.Telegram.WebHookDomain + c.Bot.Token))
	if err != nil {
		c.Log.Fatal(err.Error())
	}

	updates := c.Bot.ListenForWebhook("/" + c.Bot.Token)
	go http.ListenAndServe(c.Cfg.Telegram.ListenAddress, nil)

	c.Log.Info("Listening on " + c.Cfg.Telegram.ListenAddress)
	c.Log.Info("Webhook URL: " + c.Cfg.Telegram.WebHookDomain + c.Bot.Token)

	for update := range updates {
		if update.Message != nil {
			if update.Message.From != nil {
				if update.Message.Date > (int(time.Now().Unix()) - 5) {
					go c.Router.RouteRequest(&update)
				}
			}
		} else if update.InlineQuery != nil {
			c.Router.RouteInline(&update)
		} else if update.CallbackQuery != nil {
			c.Router.RouteCallback(&update)
		} else if update.ChosenInlineResult != nil {
			c.Log.Debug(update.ChosenInlineResult.ResultID)
		} else {
			continue
		}
	}
}
