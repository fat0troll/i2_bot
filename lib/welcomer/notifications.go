// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomer

import (
	// stdlib
	"strconv"
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (w *Welcomer) alertUserWithoutProfile(update tgbotapi.Update) string {
	alertGroupID, _ := strconv.ParseInt(c.Cfg.Notifications.GroupID, 10, 64)
	chat, ok := c.Getters.GetOrCreateChat(&update)
	if !ok {
		return "fail"
	}

	message := "*Новый вход пользователя без профиля в чат с ботом!*\n"
	message += "В чат _" + chat.Name + "_ вошёл некто @" + update.Message.NewChatMember.UserName
	message += ". Он получил уведомление о том, что ему нужно создать профиль в боте."

	msg := tgbotapi.NewMessage(alertGroupID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

func (w *Welcomer) alertSpyUser(update tgbotapi.Update) string {
	alertGroupID, _ := strconv.ParseInt(c.Cfg.Notifications.GroupID, 10, 64)
	chat, ok := c.Getters.GetOrCreateChat(&update)
	if !ok {
		return "fail"
	}

	message := "*Шпион в деле!*\n"
	message += "В чат _" + chat.Name + "_ вошёл некто @" + update.Message.NewChatMember.UserName
	message += ". У него профиль другой лиги. Ждём обновлений."

	msg := tgbotapi.NewMessage(alertGroupID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
