// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// PokememeAddSuccessMessage shows pokememe adding success message
func (t *Talkers) PokememeAddSuccessMessage(update *tgbotapi.Update) {
	message := "*Покемем успешно добавлен.*\n\n"
	message += "Посмотреть всех известных боту покемемов можно командой /pokedeks"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// PokememeAddDuplicateMessage shows pokememe add duplication message
func (t *Talkers) PokememeAddDuplicateMessage(update *tgbotapi.Update) {
	message := "*Мы уже знаем об этом покемеме*\n\n"
	message += "Посмотреть всех известных боту покемемов можно командой /pokedeks\n\n"
	message += "Если у покемема изменились описание или характеристики, напиши @fat0troll для обновления базы."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// PokememeAddFailureMessage shows pokememe add error message
func (t *Talkers) PokememeAddFailureMessage(update *tgbotapi.Update) {
	message := "*Неудачно получилось :(*\n\n"
	message += "Случилась жуткая ошибка, и мы не смогли записать покемема в базу. Напиши @fat0troll, он разберется.\n\n"
	message += "Посмотреть всех известных боту покемемов можно командой /pokedeks"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
