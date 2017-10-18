// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../config"
)

// HelpMessage gives user all available commands
func (t *Talkers) HelpMessage(update tgbotapi.Update) {
	message := "*Бот Инстинкта Enchanched.*\n\n"
	message += "Текущая версия: *" + config.VERSION + "*\n\n"
	message += "Список команд:\n\n"
	message += "+ /me – посмотреть свой сохраненный профиль в боте\n"
	message += "+ /best – посмотреть лучших покемонов для поимки\n"
	message += "+ /pokedeks – получить список известных боту покемемов\n"
	message += "+ /help – выводит данное сообщение\n"
	message += "\n\n"
	message += "Связаться с автором: @fat0troll\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
