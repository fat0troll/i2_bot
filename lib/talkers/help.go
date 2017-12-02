// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/config"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// AcademyMessage gives user link to Bastion
func (t *Talkers) AcademyMessage(update *tgbotapi.Update) {
	message := "*Академия Инстинкта*\n"
	message += "Чат для обучения новичков предумростям игры расположен по ссылке: https://t.me/joinchat/G2vME04jk02v2etRmumylg"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// BastionMessage gives user link to Bastion
func (t *Talkers) BastionMessage(update *tgbotapi.Update) {
	message := "*Бастион Инстинкта*\n"
	message += "Общий чат лиги расположен по ссылке: https://t.me/joinchat/G2vME0mIX-QHjjxE\\_JBzoQ"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// HelpMessage gives user all available commands
func (t *Talkers) HelpMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) {
	message := "*Бот Инстинкта Enchanched.*\n\n"
	message += "Текущая версия: *" + config.VERSION + "*\n\n"
	message += "Список команд\n\n"
	message += "+ /me – посмотреть свой сохраненный профиль в боте\n"
	message += "+ /best – посмотреть лучших покемонов для поимки\n"
	message += "+ /pokedeks – получить список известных боту покемемов\n"
	message += "+ /academy — Академия Инстинкта\n"
	message += "+ /bastion — Бастион Инстинкта\n"
	if c.Users.PlayerBetterThan(playerRaw, "admin") {
		message += "+ /send\\_all _текст_ — отправить сообщение всем пользователям бота\n"
		message += "+ /send\\_league _текст_ — отправить сообщение всем пользователям бота, у которых профиль лиги Инстинкт\n"
		message += "+ /chats — получить список групп, в которых работает бот.\n"
		message += "+ /squads — получить список отрядов.\n"
		message += "+ /pin _номера чатов_ _текст_ — отправить сообщение в чаты с номерами. Сообщение будет автоматичекси запинено. Пример: \"/pin 2,3,5 привет мир\". Внимание: между номерами чатов ставятся запятые без пробелов! Всё, что идёт после второго пробела в команде — сообщение\n"
		message += "+ /pin\\_all _текст_ — отправить сообщение во все группы, где находится бот. Сообщение будет автоматически запинено.\n"
		message += "+ /orders — просмотреть приказы на атаку\n"
	}
	if c.Users.PlayerBetterThan(playerRaw, "academic") {
		message += "+ /users —  просмотреть зарегистрированных пользователей бота\n"
		message += "+ /find\\_user _строка_ — найти игрока в боте по его нику или имени. Ник ищется без собачки в начале\n"
	}
	message += "+ /help – выводит данное сообщение\n"

	message += "\n\n"
	message += "Связаться с автором: @fat0troll\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
