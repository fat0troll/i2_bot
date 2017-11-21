// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomerinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// WelcomerInterface implements Welcomer for importing via appcontex
type WelcomerInterface interface {
	Init()

	PrivateWelcomeMessageUnauthorized(update *tgbotapi.Update)
	PrivateWelcomeMessageAuthorized(update *tgbotapi.Update, playerRaw *dbmapping.Player)
	GroupWelcomeMessage(update *tgbotapi.Update) string
}
