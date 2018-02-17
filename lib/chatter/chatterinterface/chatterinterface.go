// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package chatterinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

// ChatterInterface implements Chatter for importing via appcontext.
type ChatterInterface interface {
	Init()

	BanUserFromChat(user *tgbotapi.User, chatRaw *dbmapping.Chat)
	ProtectChat(update *tgbotapi.Update, playerRaw *dbmapping.Player, chatRaw *dbmapping.Chat) string

	UpdateChatTitle(chatRaw *dbmapping.Chat, newTitle string) (*dbmapping.Chat, bool)
	UpdateChatTelegramID(update *tgbotapi.Update) (*dbmapping.Chat, bool)

	GroupsList(update *tgbotapi.Update) string
}
