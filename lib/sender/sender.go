// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package sender

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendMarkdownAnswer sends markdown-powered message as answer
func (s *Sender) SendMarkdownAnswer(update *tgbotapi.Update, message string) {
	c.Log.Debug("Sending Markdown answer...")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = parseMode

	_, err := c.Bot.Send(msg)
	if err != nil {
		c.Log.Error(err.Error())
	}
}

// SendMarkdownMessageToChatID sends markdown-powered message to specified chat
func (s *Sender) SendMarkdownMessageToChatID(chatID int64, message string) {
	c.Log.Debug("Sending Markdown message to specified chat...")
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = parseMode

	_, err := c.Bot.Send(msg)
	if err != nil {
		c.Log.Error(err.Error())
	}
}

// SendMarkdownReply sends markdown-powered message as reply
func (s *Sender) SendMarkdownReply(update *tgbotapi.Update, message string, messageID int) {
	c.Log.Debug("Sending Markdown reply...")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = parseMode
	msg.ReplyToMessageID = messageID

	_, err := c.Bot.Send(msg)
	if err != nil {
		c.Log.Error(err.Error())
	}
}
