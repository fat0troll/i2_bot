// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package broadcaster

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
)

// AdminBroadcastMessageSend sends saved message to all private chats
func (b *Broadcaster) AdminBroadcastMessageSend(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	messageNum := strings.Replace(update.Message.Text, "/send_confirm ", "", 1)
	messageNumInt, _ := strconv.Atoi(messageNum)
	messageRaw, ok := b.getBroadcastMessageByID(messageNumInt)
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

	privateChats, ok := c.Chatter.GetAllPrivateChats()
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

	messageRaw, ok = b.updateBroadcastMessageStatus(messageRaw.ID, "sent")
	if !ok {
		return "fail"
	}

	message := "Сообщение всем отправлено. Надеюсь, пользователи бота за него тебя не убьют.\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
