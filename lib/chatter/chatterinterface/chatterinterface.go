// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package chatterinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

// ChatterInterface implements Chatter for importing via appcontext.
type ChatterInterface interface {
	Init()

	BanUserFromChat(user *tgbotapi.User, chatRaw *dbmapping.Chat)
	ProtectChat(update *tgbotapi.Update, playerRaw *dbmapping.Player, chatRaw *dbmapping.Chat) string

	GroupsList(update *tgbotapi.Update) string
}
