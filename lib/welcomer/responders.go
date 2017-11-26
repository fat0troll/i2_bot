// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomer

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
)

// PrivateWelcomeMessageUnauthorized tell new user what to do.
func (w *Welcomer) PrivateWelcomeMessageUnauthorized(update *tgbotapi.Update) {
	message := "*Бот Инстинкта приветствует тебя!*\n\n"
	message += "Для начала работы с ботом, пожалуйста, перешли от бота игры @PokememBroBot профиль героя.\n"
	message += "Все дальнейшие действия с ботом возможны лишь при наличии профиля игрока."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// PrivateWelcomeMessageAuthorized greets existing user
func (w *Welcomer) PrivateWelcomeMessageAuthorized(update *tgbotapi.Update, playerRaw *dbmapping.Player) {
	message := "*Бот Инстинкта приветствует тебя. Снова.*\n\n"
	message += "Привет, " + update.Message.From.FirstName + " " + update.Message.From.LastName + "!\n"
	message += "Последнее обновление информации о тебе: " + playerRaw.UpdatedAt.Format("02.01.2006 15:04:05 -0700")
	message += "\nПосмотреть информацию о себе: /me"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// PrivateWelcomeMessageSpecial greets existing user with `special` access
func (w *Welcomer) PrivateWelcomeMessageSpecial(update *tgbotapi.Update, playerRaw *dbmapping.Player) {
	message := "*Бот Инстинкта приветствует тебя. Снова.*\n\n"
	message += "Привет, " + update.Message.From.FirstName + " " + update.Message.From.LastName + "!\n"
	message += "\nБудь аккуратен, суперюзер!"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// GroupWelcomeMessage welcomes new user on group or bot itself
func (w *Welcomer) GroupWelcomeMessage(update *tgbotapi.Update) string {
	newUsers := *update.Message.NewChatMembers
	for i := range newUsers {
		newUser := newUsers[i]
		if (newUser.UserName == "i2_bot") || (newUser.UserName == "i2_dev_bot") {
			w.groupStartMessage(update)
		} else {
			defaultGroupID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
			if update.Message.Chat.ID == defaultGroupID {
				w.groupWelcomeUser(update, &newUser)
			}
		}
	}

	return "ok"
}
