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
    "../router/routerinterface"
)

type Context struct {
    Cfg         *config.Config
    Bot         *tgbotapi.BotAPI
	Router      routerinterface.RouterInterface
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
