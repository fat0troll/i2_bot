// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package routerinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// RouterInterface implements Router for importing via appcontext.
type RouterInterface interface {
	Init()
	RouteRequest(update *tgbotapi.Update) string
}
