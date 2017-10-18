// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package connections

import (
	// stdlib
	"log"
	// 3rd-party
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	// local
	"../config"
)

// BotInit initializes connection to Telegram
func BotInit(cfg *config.Config) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.APIToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Bot version: " + config.VERSION)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

// DBInit initializes database connection
func DBInit(cfg *config.Config) *sqlx.DB {
	database, err := sqlx.Connect("mysql", cfg.Database.User+":"+cfg.Database.Password+"@tcp("+cfg.Database.Host+":"+cfg.Database.Port+")/"+cfg.Database.Database+"?parseTime=true&charset=utf8mb4,utf8")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database connection established!")
	return database
}
