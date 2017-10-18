// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// AnyMessageUnauthorized throws when user can't do something
func (t *Talkers) AnyMessageUnauthorized(update tgbotapi.Update) {
	message := "Извини, действие для тебя недоступно. Возможно, у меня нет твоего профиля или же твои права недостаточны для совершения данного действия\n\n"
	message += "Если тебе кажется, что это ошибка, пиши @fat0troll.\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// GetterError throws when bot can't get something
func (t *Talkers) GetterError(update tgbotapi.Update) {
	message := "Ой, внутренняя ошибка в боте :(\n\n"
	message += "Напиши @fat0troll, приложив форвардом последние сообщения до этого.\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
