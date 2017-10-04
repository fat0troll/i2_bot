// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
    // 3rd party
	"gopkg.in/telegram-bot-api.v4"
)

type RouterHandler struct {}

func (rh RouterHandler) Init() {
    r.Init()
}

func (rh RouterHandler) RouteRequest(update tgbotapi.Update) string {
    return r.RouteRequest(update)
}
