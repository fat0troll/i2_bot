// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package routerinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// RouterInterface implements Router for importing via appcontext.
type RouterInterface interface {
	Init()

	RouteCallback(update tgbotapi.Update) string
	RouteInline(update tgbotapi.Update) string
	RouteRequest(update tgbotapi.Update) string
}
