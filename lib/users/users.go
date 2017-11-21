// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Internal functions for Users package

// profileAddSuccessMessage shows profile addition success message
func (u *Users) profileAddSuccessMessage(update *tgbotapi.Update) {
	message := "*Профиль успешно обновлен.*\n\n"
	message += "Функциональность бота держится на актуальности профилей. Обновляйся почаще, и да пребудет с тобой Рандом!\n"
	message += "Сохраненный профиль ты можешь просмотреть командой /me.\n\n"
	message += "/best – посмотреть лучших покемемов для поимки"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// profileAddFailureMessage shows profile addition failure message
func (u *Users) profileAddFailureMessage(update *tgbotapi.Update) {
	message := "*Неудачно получилось :(*\n\n"
	message += "Случилась жуткая ошибка, и мы не смогли записать профиль в базу. Напиши @fat0troll, он разберется."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
