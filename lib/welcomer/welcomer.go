// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomer

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func (w *Welcomer) groupWelcomeUser(update *tgbotapi.Update, newUser *tgbotapi.User) string {
	playerRaw, ok := c.Users.GetOrCreatePlayer(newUser.ID)
	if !ok {
		return "fail"
	}

	profileRaw, profileExist := c.Users.GetProfile(playerRaw.ID)

	message := "*Бот Инстинкта приветствует тебя, *@"
	message += c.Users.FormatUsername(newUser.UserName)
	message += "*!*\n\n"

	if profileExist {
		if playerRaw.LeagueID == 1 {
			message += "Рад тебя видеть! Не забывай обновлять профиль почаще, и да пребудет с тобой Рандом!\n"
			message += "Последнее обновление твоего профиля: " + profileRaw.CreatedAt.Format("02.01.2006 15:04:05") + "."
		} else {
			message += "Обнови профиль, отправив его мне в личку."

			w.alertSpyUser(update, newUser)
		}
	} else {
		// newbie
		message += "Перешли мне свой профиль для дальнейших инструкций.\n"

		w.alertUserWithoutProfile(update, newUser)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

func (w *Welcomer) groupStartMessage(update *tgbotapi.Update) string {
	message := "*Бот Инстинкта приветствует этот чатик!*\n\n"
	message += "На службе здравого смысла с " + time.Now().Format("02.01.2006 15:04:05") + "."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
