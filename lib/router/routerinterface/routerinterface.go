// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package routerinterface

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type RouterInterface interface {
	Init()
	RouteRequest(update tgbotapi.Update) string
}
