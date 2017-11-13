// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomerinterface

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// WelcomerInterface implements Welcomer for importing via appcontex
type WelcomerInterface interface {
	Init()
	WelcomeMessage(update tgbotapi.Update) string
}
