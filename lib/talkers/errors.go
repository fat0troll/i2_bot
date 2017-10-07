// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
    // 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (t *Talkers) AnyMessageUnauthorized(update tgbotapi.Update) {
    error_message := "Извини, действие для тебя недоступно. Возможно, у меня нет твоего профиля или же твои права недостаточны для совершения данного действия\n\n"
    error_message += "Если тебе кажется, что это ошибка, пиши @fat0troll.\n"

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, error_message)
    msg.ParseMode = "Markdown"

    c.Bot.Send(msg)
}
