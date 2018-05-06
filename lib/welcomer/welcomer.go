// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomer

import (
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (w *Welcomer) groupWelcomeUser(update *tgbotapi.Update, newUser *tgbotapi.User) string {
	playerRaw, err := c.DataCache.GetPlayerByTelegramID(newUser.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	_, profileExist := c.DataCache.GetProfileByPlayerID(playerRaw.ID)

	message := "*Бот Инстинкта приветствует тебя, *"
	message += c.Users.GetPrettyName(newUser)
	message += "*!*\n\n"

	if profileExist == nil {
		if playerRaw.LeagueID != 1 {
			w.alertSpyUser(update, newUser)
		}
	} else {
		c.Log.Info("Following profile error is OK.")
		c.Log.Info(profileExist.Error())
		w.alertUserWithoutProfile(update, newUser)
	}

	message += "Приветствую тебя, гость лиги Инстинкт! Для регистрации в Лиге и получения доступа к ее ресурсам и чатам напиши скорее мне, @i2\\_bot, в личку и скинь свою статистику.\n\n"

	message += "Алгоритм, как зарегистрироваться:\n\n1) Нажимаешь кнопку ниже, и выбираешь @PokememBroBot как цель для сообщения\n2) В ответ получаешь длинное сообщение-статистику.\n3) Полученное сообщение форвардишь в @i2\\_bot (не забудь заранее нажать там /start!)\n4) Готово, вы восхитительны."

	message += "\n\nПравила чатов лиги Инстинкт:/rules. Незнание правил — отягчающее обстоятельство."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	var row []tgbotapi.InlineKeyboardButton
	btn := tgbotapi.NewInlineKeyboardButtonSwitch("Получить статистику у @PokememBroBot", "Статы")
	row = append(row, btn)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	msg.ReplyMarkup = keyboard

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
