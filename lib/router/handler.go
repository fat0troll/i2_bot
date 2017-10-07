// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
    // 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type RouterHandler struct {}

func (rh RouterHandler) Init() {
    r.Init()
}

func (rh RouterHandler) RouteRequest(update tgbotapi.Update) string {
    return r.RouteRequest(update)
}
