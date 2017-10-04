// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package connections

import (
    // stdlib
    "log"
    // 3rd-party
	"gopkg.in/telegram-bot-api.v4"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
    // local
    "../config"
)

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

func DBInit(cfg *config.Config) *sqlx.DB {
    database, err := sqlx.Connect("mysql", cfg.Database.User + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + cfg.Database.Port + ")/" + cfg.Database.Database + "?parseTime=true&charset=utf8mb4,utf8")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Database connection established!")
    return database
}
