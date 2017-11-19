// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package chatterinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// ChatterInterface implements Chatter for importing via appcontext.
type ChatterInterface interface {
	Init()

	GetOrCreateChat(update *tgbotapi.Update) (dbmapping.Chat, bool)
	GetChatByID(chatID int64) (dbmapping.Chat, bool)
	GetAllPrivateChats() ([]dbmapping.Chat, bool)
	GetAllGroupChats() ([]dbmapping.Chat, bool)

	UpdateChatTitle(chatRaw *dbmapping.Chat, newTitle string) (*dbmapping.Chat, bool)
	UpdateChatTelegramID(update *tgbotapi.Update) (*dbmapping.Chat, bool)

	GroupsList(update *tgbotapi.Update) string
}
