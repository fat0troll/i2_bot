// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
    // 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
    // local
    "../config"
)

func (t *Talkers) HelpMessage(update tgbotapi.Update) {
    help_message := "*Бот Инстинкта Enchanched.*\n\n"
    help_message += "Текущая версия: *" + config.VERSION + "*\n\n"
    help_message += "Список команд:\n\n"
	help_message += "+ /me – посмотреть свой сохраненный профиль в боте\n"
	help_message += "+ /pokedeks – получить список известных боту покемемов\n"
    help_message += "+ /help – выводит данное сообщение\n"
    help_message += "\n\n"
    help_message += "Связаться с автором: @fat0troll\n"

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, help_message)
    msg.ParseMode = "Markdown"

    c.Bot.Send(msg)
}
