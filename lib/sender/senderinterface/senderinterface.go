// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package senderinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// SenderInterface implements Sender for importing via appcontext
type SenderInterface interface {
	Init()

	SendMarkdownAnswer(update *tgbotapi.Update, message string)
	SendMarkdownMessageToChatID(chatID int64, message string)
	SendMarkdownReply(update *tgbotapi.Update, message string, messageID int)
}
