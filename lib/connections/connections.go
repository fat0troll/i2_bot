// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package connections

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/config"
	"lab.pztrn.name/golibs/mogrus"
)

// BotInit initializes connection to Telegram
func BotInit(cfg *config.Config, lg *mogrus.LoggerHandler) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.APIToken)
	if err != nil {
		lg.Fatal(err.Error())
	}

	bot.Debug = true

	lg.Info("Bot version: " + config.VERSION)
	lg.Info("Authorized on account @", bot.Self.UserName)

	return bot
}

// DBInit initializes database connection
func DBInit(cfg *config.Config, lg *mogrus.LoggerHandler) *sqlx.DB {
	database, err := sqlx.Connect("mysql", cfg.Database.User+":"+cfg.Database.Password+"@tcp("+cfg.Database.Host+":"+cfg.Database.Port+")/"+cfg.Database.Database+"?parseTime=true&charset=utf8mb4,utf8")
	if err != nil {
		lg.Fatal(err)
	}
	lg.Info("Database connection established!")
	return database
}
