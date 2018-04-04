// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package chatter

import (
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// GroupsList lists all chats where bot exist
func (ct *Chatter) GroupsList(update *tgbotapi.Update) string {
	groupChats := c.DataCache.GetAllGroupChats()

	academyChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.AcademyID, 10, 64)
	bastionChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
	defaultChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.DefaultID, 10, 64)
	hqChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
	gamesChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.GamesID, 10, 64)

	message := "*Бот состоит в следующих групповых чатах:*\n"

	for i := range groupChats {
		message += "---\n"
		message += "\\[#" + strconv.Itoa(groupChats[i].ID) + "] _" + c.Users.FormatUsername(groupChats[i].Name) + "_\n"
		message += "Telegram ID: " + strconv.FormatInt(groupChats[i].TelegramID, 10) + "\n"
		squad, squadExistErr := c.DataCache.GetSquadByChatID(groupChats[i].ID)
		if squadExistErr == nil {
			message += "Статистика отряда:\n"
			message += c.Statistics.SquadStatictics(squad.ID)
		} else {
			if groupChats[i].TelegramID == academyChatID {
				message += "Является академией лиги\n"
			}
			if groupChats[i].TelegramID == bastionChatID {
				message += "Является бастионом лиги\n"
			}

			if groupChats[i].TelegramID == defaultChatID {
				message += "Является чатом по умолчанию лиги\n"
			}

			if groupChats[i].TelegramID == hqChatID {
				message += "Является чатом совета лиги\n"
			}

			if groupChats[i].TelegramID == gamesChatID {
				message += "Является игровым чатом\n"
			}
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
