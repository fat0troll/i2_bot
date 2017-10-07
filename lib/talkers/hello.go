// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
    // 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../dbmappings"
)

func (t *Talkers) HelloMessageUnauthorized(update tgbotapi.Update) {
    hello_message := "*Бот Инстинкта приветствует тебя!*\n\n"
    hello_message += "Для начала работы с ботом, пожалуйста, перешли от бота игры @PokememBroBot профиль героя.\n"
    hello_message += "Все дальнейшие действия с ботом возможны лишь при наличии профиля игрока."

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, hello_message)
    msg.ParseMode = "Markdown"

    c.Bot.Send(msg)
}

func (t *Talkers) HelloMessageAuthorized(update tgbotapi.Update, player_raw dbmappings.Players) {
    hello_message := "*Бот Инстинкта приветствует тебя. Снова.*\n\n"
    hello_message += "Привет, " + update.Message.From.FirstName + " " + update.Message.From.LastName + "!\n"
	hello_message += "Последнее обновление информации о тебе: " + player_raw.Updated_at.Format("02.01.2006 15:04:05 -0700")
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, hello_message)
    msg.ParseMode = "Markdown"

    c.Bot.Send(msg)
}
