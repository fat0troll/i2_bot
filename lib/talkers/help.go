// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/config"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// AcademyMessage gives user link to Bastion
func (t *Talkers) AcademyMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) {
	message := ""

	if playerRaw.LeagueID > 1 {
		message = "Иди нахуй, шпионское отродье"
	} else if playerRaw.LeagueID == 0 {
		message = "Заполни профиль и попробуй ещё раз"
	} else {
		message += "*Академия Инстинкта*\n"
		message += "Чат для обучения новичков предумростям игры расположен по ссылке: https://t.me/joinchat/G2vME04jk02v2etRmumylg"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// BastionMessage gives user link to Bastion
func (t *Talkers) BastionMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) {
	message := ""

	if playerRaw.LeagueID > 1 {
		message = "Иди нахуй, шпионское отродье"
	} else if playerRaw.LeagueID == 0 {
		message = "Заполни профиль и попробуй ещё раз"
	} else {
		message += "*Бастион Инстинкта*\n"
		message += "Общий чат лиги расположен по ссылке: https://t.me/joinchat/G2vME0mIX-QHjjxE\\_JBzoQ"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// HelpMessage gives user all available commands
func (t *Talkers) HelpMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) {
	message := "*Бот Инстинкта Enhanced.*\n\n"
	message += "Текущая версия: *" + config.VERSION + "*\n\n"
	message += "Список команд\n\n"
	message += "\\* /me – посмотреть свой сохраненный профиль в боте\n"
	message += "\\* /best – посмотреть лучших покемонов для поимки\n"
	message += "\\* /advice – посмотреть самых дорогих покемонов для поимки\n"
	message += "\\* /top — топ игроков лиги\n"
	message += "\\* /top\\_my — топ игроков лиги твоего уровня\n"
	message += "\\* /pokedeks – получить список известных боту покемемов\n"
	message += "\\* /reminders — настроить оповещения на Турнир лиг\n"
	message += "\\* /academy — Академия Инстинкта\n"
	message += "\\* /bastion — Бастион Инстинкта\n"
	if c.Users.PlayerBetterThan(playerRaw, "admin") {
		message += "\\* /send\\_all _текст_ — отправить сообщение всем пользователям бота\n"
		message += "\\* /send\\_league _текст_ — отправить сообщение всем пользователям бота, у которых профиль лиги Инстинкт\n"
		message += "\\* /chats — получить список групп, в которых работает бот.\n"
		message += "\\* /squads — получить список отрядов.\n"
		message += "\\* /pin _номера чатов_ _текст_ — отправить сообщение в чаты с номерами. Сообщение будет автоматичекси запинено. Пример: \"/pin 2,3,5 привет мир\". Внимание: между номерами чатов ставятся запятые без пробелов! Всё, что идёт после второго пробела в команде — сообщение\n"
		message += "\\* /pin\\_all _текст_ — отправить сообщение во все группы, где находится бот. Сообщение будет автоматически запинено.\n"
		message += "\\* /orders — просмотреть приказы на атаку\n"
	}
	if c.Users.PlayerBetterThan(playerRaw, "academic") {
		message += "\\* /users —  просмотреть зарегистрированных пользователей бота\n"
		message += "\\* /find\\_level _цифра_ — показать всех игроков соответствующего уровня. Учитываются профили за 72 часа\n"
		message += "\\* /find\\_user _строка_ — найти игрока в боте по его нику или имени. Ник ищется без собачки в начале\n"
	}
	message += "\\* /help – выводит данное сообщение\n"

	message += "\n\n"
	message += "Техническая поддержка бота: https://t.me/joinchat/AAkt5EgFBU9Q9iXJMvDG6A\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// FiveOffer sends all users with 5 pokeballs limit offer for increasing pokeballs limit
func (t *Talkers) FiveOffer(update *tgbotapi.Update) string {
	players := []dbmapping.Player{}

	err := c.Db.Select(&players, "SELECT p.* FROM players p, profiles pp WHERE p.id = pp.player_id AND pp.pokeballs = 5")
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	for i := range players {
		message := "Псст, я тут заметил, что у тебя всего 5 покеболов? Хочешь увеличить их лимит на 2 или даже больше? У всех игроков есть возможность получить бонус!\n\n1. Перейти по ссылке: https://telegram.me/storebot?start=pokemembrobot\n2. Нажать Start\n3. Выбрать ⭐️⭐️⭐️⭐️⭐️\n4. ОБЯЗАТЕЛЬНО написать, что вам нравится в игре (на русском языке). Оставьте большой и красочный отзыв!\n5. Переслать переписку с @storebot в тех поддержку игры @PBhelp<— только ему! и больше никому! (с текстом вашего отзыва)\n6. После проверки получить бонус 🎁 +2 к лимиту ⭕️ А если отзыв понравится админам (и это бывает очень часто), то бонус будет больше!\n7. Проверка - может занять некоторое время. Админы обязательно ответят вам о результатах проверки."

		msg := tgbotapi.NewMessage(int64(players[i].TelegramID), message)
		msg.ParseMode = "Markdown"

		c.Bot.Send(msg)
	}

	message := "Enlarge your pokeballs! Сообщение отправлено."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
