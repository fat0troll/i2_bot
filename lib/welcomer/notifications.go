// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomer

import (
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (w *Welcomer) alertUserWithoutProfile(update *tgbotapi.Update, newUser *tgbotapi.User) string {
	alertGroupID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
	chat, err := c.DataCache.GetOrCreateChat(update)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	userName := ""
	if newUser.UserName != "" {
		userName += "@" + newUser.UserName
	} else {
		userName += newUser.FirstName
		if newUser.LastName != "" {
			userName += " " + newUser.LastName
		}
	}

	message := "*Новый вход пользователя без профиля в чат с ботом!*\n"
	message += "В чат _" + chat.Name + "_ вошёл некто " + c.Users.FormatUsername(userName)
	message += ". Он получил уведомление о том, что ему нужно создать профиль в боте."

	msg := tgbotapi.NewMessage(alertGroupID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

func (w *Welcomer) alertSpyUser(update *tgbotapi.Update, newUser *tgbotapi.User) string {
	alertGroupID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
	chat, err := c.DataCache.GetOrCreateChat(update)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	userName := ""
	if newUser.UserName != "" {
		userName += "@" + newUser.UserName
	} else {
		userName += newUser.FirstName
		if newUser.LastName != "" {
			userName += " " + newUser.LastName
		}
	}

	message := "*Шпион в деле!*\n"
	message += "В чат _" + chat.Name + "_ вошёл некто " + c.Users.FormatUsername(userName)
	message += ". У него профиль другой лиги. Ждём обновлений."

	msg := tgbotapi.NewMessage(alertGroupID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
