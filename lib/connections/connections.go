// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package connections

import (
	"bitbucket.org/pztrn/mogrus"
	_ "github.com/go-sql-driver/mysql" // MySQL driver for sqlx
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/proxy"
	"net/http"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/config"
)

// botInitDirect used when no proxy in config file
func botInitDirect(cfg *config.Config, lg *mogrus.LoggerHandler) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.APIToken)
	if err != nil {
		lg.Fatal(err.Error())
	}

	bot.Debug = true

	lg.Info("Bot version: " + config.VERSION)
	lg.Info("Authorized on account @", bot.Self.UserName)

	return bot
}

// botInitWithProxy used when there is proxy in config file
func botInitWithProxy(cfg *config.Config, lg *mogrus.LoggerHandler) *tgbotapi.BotAPI {
	proxyAuth := proxy.Auth{}
	if cfg.Proxy.Username != "" {
		proxyAuth.User = cfg.Proxy.Username
		proxyAuth.Password = cfg.Proxy.Password
	}

	var dialProxy proxy.Dialer
	var err error
	if cfg.Proxy.Username != "" {
		dialProxy, err = proxy.SOCKS5("tcp", cfg.Proxy.Address, &proxyAuth, proxy.Direct)
		if err != nil {
			lg.Fatal(err.Error())
		}
	} else {
		dialProxy, err = proxy.SOCKS5("tcp", cfg.Proxy.Address, &proxyAuth, proxy.Direct)
		if err != nil {
			lg.Fatal(err.Error())
		}
	}

	proxyTransport := &http.Transport{Dial: dialProxy.Dial}
	proxyClient := http.Client{Transport: proxyTransport}

	bot, err := tgbotapi.NewBotAPIWithClient(cfg.Telegram.APIToken, &proxyClient)
	if err != nil {
		lg.Fatal(err.Error())
	}

	bot.Debug = true

	lg.Info("Bot version: " + config.VERSION)
	lg.Info("Authorized on account @", bot.Self.UserName)

	return bot
}

// External functions

// BotInit initializes connection to Telegram
func BotInit(cfg *config.Config, lg *mogrus.LoggerHandler) *tgbotapi.BotAPI {
	if cfg.Proxy.Enabled {
		lg.Info("Using proxy for bot: " + cfg.Proxy.Address)
		return botInitWithProxy(cfg, lg)
	}

	lg.Info("Using direct connection to Telegram")
	return botInitDirect(cfg, lg)
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
