// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// HelloMessageUnauthorized tell new user what to do.
func (t *Talkers) HelloMessageUnauthorized(update tgbotapi.Update) {
	message := "*Бот Инстинкта приветствует тебя!*\n\n"
	message += "Для начала работы с ботом, пожалуйста, перешли от бота игры @PokememBroBot профиль героя.\n"
	message += "Все дальнейшие действия с ботом возможны лишь при наличии профиля игрока."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// HelloMessageAuthorized greets existing user
func (t *Talkers) HelloMessageAuthorized(update tgbotapi.Update, playerRaw dbmapping.Player) {
	message := "*Бот Инстинкта приветствует тебя. Снова.*\n\n"
	message += "Привет, " + update.Message.From.FirstName + " " + update.Message.From.LastName + "!\n"
	message += "Последнее обновление информации о тебе: " + playerRaw.UpdatedAt.Format("02.01.2006 15:04:05 -0700")
	message += "\nПосмотреть информацию о себе: /me"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
