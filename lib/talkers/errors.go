// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// AnyMessageUnauthorized throws when user can't do something
func (t *Talkers) AnyMessageUnauthorized(update *tgbotapi.Update) string {
	message := "Извини, действие для тебя недоступно. Возможно, у меня нет твоего профиля или же твои права недостаточны для совершения данного действия\n\n"
	message += "Техническая поддержка бота: https://t.me/joinchat/AAkt5EgFBU9Q9iXJMvDG6A.\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "fail"
}

// BotError throws when bot can't do something
func (t *Talkers) BotError(update *tgbotapi.Update) string {
	message := "Ой, внутренняя ошибка в боте :(\n\n"
	message += "Техническая поддержка бота: https://t.me/joinchat/AAkt5EgFBU9Q9iXJMvDG6A. Напиши сюда, приложив скриншоты с перепиской бота.\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "fail"
}
