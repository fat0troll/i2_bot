// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package broadcaster

import (
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/constants"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

// AdminBroadcastMessageSend sends saved message to all private chats
func (b *Broadcaster) AdminBroadcastMessageSend(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	messageNum := update.Message.CommandArguments()
	messageNumInt, _ := strconv.Atoi(messageNum)
	messageRaw, ok := b.getBroadcastMessageByID(messageNumInt)
	if !ok {
		return constants.BotError
	}
	if messageRaw.AuthorID != playerRaw.ID {
		return constants.UserRequestForbidden
	}
	if messageRaw.Status != "new" {
		return constants.UserRequestFailed
	}

	broadcastingMessageBody := messageRaw.Text

	profileRaw, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return constants.UserRequestFailed
	}

	prettyName := profileRaw.Nickname + " (@" + profileRaw.TelegramNickname + ")"

	privateChats := []dbmapping.Chat{}
	switch messageRaw.BroadcastType {
	case "all":
		privateChats = c.DataCache.GetAllPrivateChats()
	case "league":
		privateChats = c.DataCache.GetLeaguePrivateChats()
	}

	for i := range privateChats {
		chat := privateChats[i]
		broadcastingMessage := "*Привет, " + chat.Name + "!*\n\n"
		broadcastingMessage += "*Важное сообщение от администратора *" + prettyName + "\n\n"
		broadcastingMessage += broadcastingMessageBody

		c.Sender.SendMarkdownMessageToChatID(chat.TelegramID, broadcastingMessage)
	}

	messageRaw, ok = b.updateBroadcastMessageStatus(messageRaw.ID, "sent")
	if !ok {
		return constants.BotError
	}

	message := "Сообщение отправлено. Надеюсь, пользователи бота за него тебя не убьют.\n"

	c.Sender.SendMarkdownAnswer(update, message)

	return constants.UserRequestSuccess
}
