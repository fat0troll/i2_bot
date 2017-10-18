// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// RouterHandler is a handler for router package
type RouterHandler struct{}

// Init is an initialization function of router
func (rh RouterHandler) Init() {
	r.Init()
}

// RouteRequest decides, what to do with user input
func (rh RouterHandler) RouteRequest(update tgbotapi.Update) string {
	return r.RouteRequest(update)
}
