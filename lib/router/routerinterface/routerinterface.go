// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package routerinterface

import (
    // 3rd party
	"gopkg.in/telegram-bot-api.v4"
)

type RouterInterface interface {
    Init()
    RouteRequest(update tgbotapi.Update) string
}
