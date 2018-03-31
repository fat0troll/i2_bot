// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/constants"
)

// AnyMessageUnauthorized throws when user can't do something
func (t *Talkers) AnyMessageUnauthorized(update *tgbotapi.Update) string {
	message := "Извини, действие для тебя недоступно. Возможно, у меня нет твоего профиля или же твои права недостаточны для совершения данного действия\n\n"
	message += "Техническая поддержка бота: https://t.me/joinchat/AAkt5EgFBU9Q9iXJMvDG6A.\n"

	c.Sender.SendMarkdownAnswer(update, message)

	return constants.UserRequestFailed
}

// BanError throws error for persona non grata
func (t *Talkers) BanError(update *tgbotapi.Update) string {
	message := "Вам здесь не рады. Использование бота для вас запрещено."

	c.Sender.SendMarkdownAnswer(update, message)

	return constants.UserRequestForbidden
}

// BotError throws when bot can't do something
func (t *Talkers) BotError(update *tgbotapi.Update) string {
	message := "Ой, внутренняя ошибка в боте :(\n\n"
	message += "Техническая поддержка бота: https://t.me/joinchat/AAkt5EgFBU9Q9iXJMvDG6A. Напиши сюда, приложив скриншоты с перепиской бота.\n"

	c.Sender.SendMarkdownAnswer(update, message)

	return constants.BotError
}
