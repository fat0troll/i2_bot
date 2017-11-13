// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package appcontext

import (
	// 3rd-party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/config"
	"lab.pztrn.name/fat0troll/i2_bot/lib/connections"
	// interfaces
	"lab.pztrn.name/fat0troll/i2_bot/lib/getters/gettersinterface"
	"lab.pztrn.name/fat0troll/i2_bot/lib/migrations/migrationsinterface"
	"lab.pztrn.name/fat0troll/i2_bot/lib/parsers/parsersinterface"
	"lab.pztrn.name/fat0troll/i2_bot/lib/router/routerinterface"
	"lab.pztrn.name/fat0troll/i2_bot/lib/talkers/talkersinterface"
	"lab.pztrn.name/fat0troll/i2_bot/lib/welcomer/welcomerinterface"
)

// Context is an application context struct
type Context struct {
	Cfg        *config.Config
	Bot        *tgbotapi.BotAPI
	Migrations migrationsinterface.MigrationsInterface
	Router     routerinterface.RouterInterface
	Parsers    parsersinterface.ParsersInterface
	Db         *sqlx.DB
	Talkers    talkersinterface.TalkersInterface
	Getters    gettersinterface.GettersInterface
	Welcomer   welcomerinterface.WelcomerInterface
}

// Init is a initialization function for context
func (c *Context) Init() {
	c.Cfg = config.New()
	c.Cfg.Init()
	c.Bot = connections.BotInit(c.Cfg)
	c.Db = connections.DBInit(c.Cfg)
}

// RegisterRouterInterface registering router interface in application
func (c *Context) RegisterRouterInterface(ri routerinterface.RouterInterface) {
	c.Router = ri
	c.Router.Init()
}

// RegisterMigrationsInterface registering migrations interface in application
func (c *Context) RegisterMigrationsInterface(mi migrationsinterface.MigrationsInterface) {
	c.Migrations = mi
	c.Migrations.Init()
}

// RegisterParsersInterface registering parsers interface in application
func (c *Context) RegisterParsersInterface(pi parsersinterface.ParsersInterface) {
	c.Parsers = pi
}

// RegisterTalkersInterface registering talkers interface in application
func (c *Context) RegisterTalkersInterface(ti talkersinterface.TalkersInterface) {
	c.Talkers = ti
}

// RegisterGettersInterface registering getters interface in application
func (c *Context) RegisterGettersInterface(gi gettersinterface.GettersInterface) {
	c.Getters = gi
}

// RegisterWelcomerInterface registering welcomer interface in application
func (c *Context) RegisterWelcomerInterface(wi welcomerinterface.WelcomerInterface) {
	c.Welcomer = wi
}

// RunDatabaseMigrations applies migrations on bot's startup
func (c *Context) RunDatabaseMigrations() {
	c.Migrations.SetDialect("mysql")
	c.Migrations.Migrate()
}
