// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package broadcaster

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
)

// AdminBroadcastMessageCompose saves message for future broadcast
func (b *Broadcaster) AdminBroadcastMessageCompose(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	broadcastingMessageBody := update.Message.CommandArguments()
	messageRaw, ok := dbmapping.Broadcast{}, false

	switch update.Message.Command() {
	case "send_all":
		messageRaw, ok = b.createBroadcastMessage(playerRaw, broadcastingMessageBody, "all")
		if !ok {
			return "fail"
		}
	case "send_league":
		messageRaw, ok = b.createBroadcastMessage(playerRaw, broadcastingMessageBody, "league")
		if !ok {
			return "fail"
		}
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

	switch update.Message.Command() {
	case "send_all":
		message = "Чтобы отправить сообщение всем, отправь команду /send\\_confirm " + strconv.Itoa(messageRaw.ID)
	case "send_league":
		message = "Чтобы отправить сообщение всем игрокам лиги Инстинкт, отправь команду /send\\_confirm " + strconv.Itoa(messageRaw.ID)
	}

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
