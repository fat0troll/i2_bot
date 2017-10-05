// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package appcontext

import (
    // 3rd-party
    "github.com/jmoiron/sqlx"
	"gopkg.in/telegram-bot-api.v4"
    // local
    "../config"
    "../connections"
	// interfaces
    "../migrations/migrationsinterface"
    "../parsers/parsersinterface"
    "../router/routerinterface"
)

type Context struct {
    Cfg         *config.Config
    Bot         *tgbotapi.BotAPI
    Migrations  migrationsinterface.MigrationsInterface
	Router      routerinterface.RouterInterface
    Parsers     parsersinterface.ParsersInterface
	Db 			*sqlx.DB
}

func (c *Context) Init() {
    c.Cfg = config.New()
    c.Cfg.Init()
    c.Bot = connections.BotInit(c.Cfg)
	c.Db = connections.DBInit(c.Cfg)
}

func (c *Context) RegisterRouterInterface(ri routerinterface.RouterInterface) {
	c.Router = ri
	c.Router.Init()
}

func (c *Context) RegisterMigrationsInterface(mi migrationsinterface.MigrationsInterface) {
    c.Migrations = mi
    c.Migrations.Init()
}

func (c *Context) RegisterParsersInterface(pi parsersinterface.ParsersInterface) {
    c.Parsers = pi
}

func (c *Context) RunDatabaseMigrations() {
    c.Migrations.SetDialect("mysql")
    c.Migrations.Migrate()
}
