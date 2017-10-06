// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
    // 3rd party
	"gopkg.in/telegram-bot-api.v4"
    // local
    "../config"
)

func (t *Talkers) HelpMessage(update tgbotapi.Update) {
    help_message := "*Бот Инстинкта Enchanched.*\n\n"
    help_message += "Текущая версия: *" + config.VERSION + "*\n\n"
    help_message += "Список команд:\n\n"
    help_message += "+ /help – выводит данное сообщение\n"
    help_message += "\n\n"
    help_message += "Связаться с автором: @fat0troll\n"

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, help_message)
    msg.ParseMode = "Markdown"

    c.Bot.Send(msg)
}
