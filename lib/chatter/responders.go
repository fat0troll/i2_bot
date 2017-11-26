// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package chatter

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

// GroupsList lists all chats where bot exist
func (ct *Chatter) GroupsList(update *tgbotapi.Update) string {
	groupChats, ok := ct.getAllGroupChatsWithSquads()
	if !ok {
		return "fail"
	}

	bastionChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
	defaultChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.DefaultID, 10, 64)
	hqChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)

	message := "*Бот состоит в следующих групповых чатах:*\n"

	for i := range groupChats {
		message += "---\n"
		message += "\\[#" + strconv.Itoa(groupChats[i].Chat.ID) + "] _" + c.Users.FormatUsername(groupChats[i].Chat.Name) + "_\n"
		message += "Telegram ID: " + strconv.FormatInt(groupChats[i].Chat.TelegramID, 10) + "\n"
		if groupChats[i].ChatRole == "squad" {
			message += "Статистика отряда:\n"
			message += c.Statistics.SquadStatictics(groupChats[i].Squad.ID)
		} else if groupChats[i].ChatRole == "flood" {
			message += "Является флудочатом отряда №" + strconv.Itoa(groupChats[i].Squad.ID) + "\n"
		} else {
			if groupChats[i].Chat.TelegramID == bastionChatID {
				message += "Является бастионом лиги\n"
			}

			if groupChats[i].Chat.TelegramID == defaultChatID {
				message += "Является чатом по умолчанию лиги\n"
			}

			if groupChats[i].Chat.TelegramID == hqChatID {
				message += "Является чатом совета лиги\n"
			}
		}
	}

	message += "\nЧтобы создать отряд, введите команду /make\\_squad _X Y_, где _X_ — номер чата с пинами (в нём позволено писать лишь боту и командирам), а _Y_ — чат-флудилка для общения отряда."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
