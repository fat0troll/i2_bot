// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package broadcaster

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

// AdminBroadcastMessageSend sends saved message to all private chats
func (b *Broadcaster) AdminBroadcastMessageSend(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	messageNum := update.Message.CommandArguments()
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

	profileRaw, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	prettyName := profileRaw.Nickname + " (@" + profileRaw.TelegramNickname + ")"

	privateChats := []dbmapping.Chat{}
	switch messageRaw.BroadcastType {
	case "all":
		privateChats, ok = c.Chatter.GetAllPrivateChats()
		if !ok {
			return "fail"
		}
	case "league":
		privateChats, ok = c.Chatter.GetLeaguePrivateChats()
		if !ok {
			return "fail"
		}
	}

	for i := range privateChats {
		chat := privateChats[i]
		broadcastingMessage := "*Привет, " + chat.Name + "!*\n\n"
		broadcastingMessage += "*Важное сообщение от администратора *" + prettyName + "\n\n"
		broadcastingMessage += broadcastingMessageBody

		msg := tgbotapi.NewMessage(int64(chat.TelegramID), broadcastingMessage)
		msg.ParseMode = "Markdown"
		c.Bot.Send(msg)
	}

	messageRaw, ok = b.updateBroadcastMessageStatus(messageRaw.ID, "sent")
	if !ok {
		return "fail"
	}

	message := "Сообщение отправлено. Надеюсь, пользователи бота за него тебя не убьют.\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
