// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

// GroupsList lists all chats where bot exist
func (t *Talkers) GroupsList(update *tgbotapi.Update) string {
	groupChats, ok := c.Getters.GetAllGroupChatsWithSquads()
	if !ok {
		return "fail"
	}

	message := "*Бот состоит в следующих групповых чатах:*\n"

	for i := range groupChats {
		message += "---\n"
		message += "[#" + strconv.Itoa(groupChats[i].Chat.ID) + "] _" + groupChats[i].Chat.Name + "_\n"
		message += "Telegram ID: " + strconv.FormatInt(groupChats[i].Chat.TelegramID, 10) + "\n"
		if groupChats[i].IsSquad {
			message += "Является отрядом <статистика>\n"
		} else {
			message += "Не является отрядом. Сделать отрядом: /make\\_squad" + strconv.Itoa(groupChats[i].Chat.ID) + "\n"
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
