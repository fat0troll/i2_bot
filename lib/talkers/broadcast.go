// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// stdlib
	"strconv"
	"strings"
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../dbmapping"
)

// AdminBroadcastMessageCompose saves message for future broadcast
func (t *Talkers) AdminBroadcastMessageCompose(update tgbotapi.Update, playerRaw *dbmapping.Player) string {
	broadcastingMessageBody := strings.Replace(update.Message.Text, "/send_all ", "", 1)

	messageRaw, ok := c.Getters.CreateBroadcastMessage(playerRaw, broadcastingMessageBody, "all")
	if !ok {
		return "fail"
	}

	message := "Сообщение сохранено в базу.\n"
	message += "Выглядеть оно будет так:"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	broadcastingMessage := "*Привет, %username%!*\n\n"
	broadcastingMessage += "*Важное сообщение от администратора " + update.Message.From.FirstName + " " + update.Message.From.LastName + "* (@" + update.Message.From.UserName + ")\n\n"
	broadcastingMessage += messageRaw.Text

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, broadcastingMessage)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	message = "Чтобы отправить сообщение всем, отправь команду /send\\_confirm " + strconv.Itoa(messageRaw.ID)

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

// AdminBroadcastMessageSend sends saved message to all private chats
func (t *Talkers) AdminBroadcastMessageSend(update tgbotapi.Update, playerRaw *dbmapping.Player) string {
	messageNum := strings.Replace(update.Message.Text, "/send_confirm ", "", 1)
	messageNumInt, _ := strconv.Atoi(messageNum)
	messageRaw, ok := c.Getters.GetBroadcastMessageByID(messageNumInt)
	if !ok {
		return "fail"
	}
	if messageRaw.AuthorID != playerRaw.ID {
		return "fail"
	}
	if messageRaw.Status != "new" {
		return "fail"
	}

	broadcastingMessageBody := messageRaw.Text

	privateChats, ok := c.Getters.GetAllPrivateChats()
	if !ok {
		return "fail"
	}

	for i := range privateChats {
		chat := privateChats[i]
		broadcastingMessage := "*Привет, " + chat.Name + "!*\n\n"
		broadcastingMessage += "*Важное сообщение от администратора " + update.Message.From.FirstName + " " + update.Message.From.LastName + "* (@" + update.Message.From.UserName + ")\n\n"
		broadcastingMessage += broadcastingMessageBody

		msg := tgbotapi.NewMessage(int64(chat.TelegramID), broadcastingMessage)
		msg.ParseMode = "Markdown"
		c.Bot.Send(msg)
	}

	messageRaw, ok = c.Getters.UpdateBroadcastMessageStatus(messageRaw.ID, "sent")
	if !ok {
		return "fail"
	}

	message := "Сообщение всем отправлено. Надеюсь, пользователи бота за него тебя не убьют.\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
