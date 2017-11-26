// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomer

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (w *Welcomer) alertUserWithoutProfile(update *tgbotapi.Update, newUser *tgbotapi.User) string {
	alertGroupID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
	chat, ok := c.Chatter.GetOrCreateChat(update)
	if !ok {
		return "fail"
	}

	message := "*Новый вход пользователя без профиля в чат с ботом!*\n"
	message += "В чат _" + chat.Name + "_ вошёл некто @" + newUser.UserName
	message += ". Он получил уведомление о том, что ему нужно создать профиль в боте."

	msg := tgbotapi.NewMessage(alertGroupID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

func (w *Welcomer) alertSpyUser(update *tgbotapi.Update, newUser *tgbotapi.User) string {
	alertGroupID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
	chat, ok := c.Chatter.GetOrCreateChat(update)
	if !ok {
		return "fail"
	}

	message := "*Шпион в деле!*\n"
	message += "В чат _" + chat.Name + "_ вошёл некто @" + newUser.UserName
	message += ". У него профиль другой лиги. Ждём обновлений."

	msg := tgbotapi.NewMessage(alertGroupID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}