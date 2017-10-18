// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// ProfileAddSuccessMessage shows profile addition success message
func (t *Talkers) ProfileAddSuccessMessage(update tgbotapi.Update) {
	message := "*Профиль успешно обновлен.*\n\n"
	message += "Функциональность бота держится на актуальности профилей. Обновляйся почаще, и да пребудет с тобой Рандом!\n"
	message += "Сохраненный профиль ты можешь просмотреть командой /me.\n\n"
	message += "/best – посмотреть лучших покемемов для поимки"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// ProfileAddFailureMessage shows profile addition failure message
func (t *Talkers) ProfileAddFailureMessage(update tgbotapi.Update) {
	message := "*Неудачно получилось :(*\n\n"
	message += "Случилась жуткая ошибка, и мы не смогли записать профиль в базу. Напиши @fat0troll, он разберется."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
