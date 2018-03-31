// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package sender

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendMarkdownAnswer sends markdown-powered message as reply
func (s *Sender) SendMarkdownAnswer(update *tgbotapi.Update, message string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	_, err := c.Bot.Send(msg)
	if err != nil {
		c.Log.Error(err.Error())
	}
}

// SendMarkdownMessageToChatID sends markdown-powered message to specified chat
func (s *Sender) SendMarkdownMessageToChatID(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "Markdown"

	_, err := c.Bot.Send(msg)
	if err != nil {
		c.Log.Error(err.Error())
	}
}
