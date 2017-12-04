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

	_, profileExist := c.Users.GetProfile(playerRaw.ID)

	userName := ""
	if newUser.UserName != "" {
		userName += "@" + newUser.UserName
	} else {
		userName += newUser.FirstName
		if newUser.LastName != "" {
			userName += " " + newUser.LastName
		}
	}

	message := "*Бот Инстинкта приветствует тебя, *@"
	message += c.Users.FormatUsername(userName)
	message += "*!*\n\n"

	if profileExist {
		if playerRaw.LeagueID != 1 {
			w.alertSpyUser(update, newUser)
		}
	} else {
		w.alertUserWithoutProfile(update, newUser)
	}

	message += "Приветствую тебя, гость лиги Инстинкт! Для регистрации в Лиге и получения доступа к ее ресурсам и чатам напиши скорее мне, @i2\\_bot, в личку и скинь свой профиль Герой.\n\nГайд для игроков Инстинкта: http://telegra.ph/Dobro-pozhalovat-v-Instinkt-11-22"

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
