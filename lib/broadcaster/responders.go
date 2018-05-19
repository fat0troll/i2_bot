// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package broadcaster

import (
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

// AdminBroadcastMessageCompose saves message for future broadcast
func (b *Broadcaster) AdminBroadcastMessageCompose(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	broadcastingMessageBody := update.Message.CommandArguments()
	messageMode := "none"
	switch update.Message.Command() {
	case "send_all":
		messageMode = "all"
	case "send_league":
		messageMode = "league"
	}

	messageRaw, ok := b.createBroadcastMessage(playerRaw, broadcastingMessageBody, messageMode)
	if !ok {
		return "fail"
	}

	message := "Сообщение сохранено в базу.\n"
	message += "Выглядеть оно будет так:"

	c.Sender.SendMarkdownAnswer(update, message)

	broadcastingMessage := "*Привет, %username%!*\n\n"
	broadcastingMessage += "*Важное сообщение от администратора " + update.Message.From.FirstName + " " + update.Message.From.LastName + "* (@" + update.Message.From.UserName + ")\n\n"
	broadcastingMessage += messageRaw.Text

	c.Sender.SendMarkdownAnswer(update, broadcastingMessage)

	switch update.Message.Command() {
	case "send_all":
		message = "Чтобы отправить сообщение всем, отправь команду /send\\_confirm " + strconv.Itoa(messageRaw.ID)
	case "send_league":
		message = "Чтобы отправить сообщение всем игрокам лиги Инстинкт, отправь команду /send\\_confirm " + strconv.Itoa(messageRaw.ID)
	}

	c.Sender.SendMarkdownAnswer(update, message)

	return "ok"
}
