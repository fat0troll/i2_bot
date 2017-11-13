// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/config"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// HelpMessage gives user all available commands
func (t *Talkers) HelpMessage(update tgbotapi.Update, playerRaw *dbmapping.Player) {
	message := "*Бот Инстинкта Enchanched.*\n\n"
	message += "Текущая версия: *" + config.VERSION + "*\n\n"
	message += "Список команд\n\n"
	message += "+ /me – посмотреть свой сохраненный профиль в боте\n"
	message += "+ /best – посмотреть лучших покемонов для поимки\n"
	message += "+ /pokedeks – получить список известных боту покемемов\n"
	if c.Getters.PlayerBetterThan(playerRaw, "admin") {
		message += "+ /send\\_all _текст_ — отправить сообщение всем пользователям бота\n"
		message += "+ /group\\_chats — получить список групп, в которых работает бот.\n"
	}
	message += "+ /help – выводит данное сообщение\n"

	message += "\n\n"
	message += "Связаться с автором: @fat0troll\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
