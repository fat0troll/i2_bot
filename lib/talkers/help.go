// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/config"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"time"
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
	message += "\\* /best – посмотреть 5 лучших покемонов для поимки\n"
	message += "\\* /advice – посмотреть 5 самых дорогих покемонов для поимки\n"
	message += "\\* /best\\_all – посмотреть всех лучших покемонов для поимки\n"
	message += "\\* /advice\\_all – посмотреть всех самых дорогих покемонов для поимки\n"
	if playerRaw.CreatedAt.Before(time.Now().Add(30 * 24 * 60 * 60 * time.Second)) {
		message += "\\* /best\\_nofilter – посмотреть пять лучших покемонов для поимки _без фильтра по элементам_ (полезно для сборки максимально топовой руки на высоком уровне)\n"
	}
	message += "\\* /top — топ игроков лиги\n"
	message += "\\* /top\\_my — топ игроков лиги твоего уровня\n"
	message += "\\* /pokedeks – получить список известных боту покемемов\n"
	message += "\\* /reminders — настроить оповещения на Турнир лиг\n"
	message += "\\* /academy — Академия Инстинкта\n"
	message += "\\* /bastion — Бастион Инстинкта\n"
	message += "\\* /faq — получить ответы на часто задаваемые вопросы\n"
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
	message += "\n*Благодарности*:\n"
	message += "Для поддержания сервера и его стабильной оплаты нужны средства. К сожалению, далеко не всегда они находятся в нужный момент, но всегда есть люди, готовые помочь. Я благодарю их за поддержку:\n\n"
	message += "\\* @Lenorag\n"
	message += "\\* @vanushinvi\n"
	message += "\\* @TechniqueOne\n"
	message += "\\* @Antropophag\n"
	message += "\\* @nokio404\n"
	message += "Выразить благодарность и попасть в список: 4377 7300 0246 7362\n"
	message += "_Топ ранжируется по размеру благодарности. Здесь может быть ваша реклама!_"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

// FAQMessage prints frequently asked questions
func (t *Talkers) FAQMessage(update *tgbotapi.Update) string {
	message := "*Часто задаваемые вопросы:*\n\n"

	message += "💎Кристаллы\n"
	message += "_Зачем тут донатить?_\n"
	message += "О: Самое полезное на начальных уровнях это:\n1.  Увеличение лимита ⭕️покеболов до 24 за 100💎\n2.  Турболовля за 20💎, покеболы приходят не раз в час, а каждые 40 секунд, время действия турболовли 30 минут, за это время ты получишь 45 покеболов. Охота также ускоряется: вместо 2 минут, охота длится 40 секунд.\n"
	message += "_Где и как задонатить?_\n"
	message += "В игровом боте напиши /crystalls и перейди по выданной ссылке. Стоимость кристаллов зависит от количества: чем больше ты их купишь за один раз, тем дешевле они обойдутся (от 3 до 10 рублей за 1💎). Если это твой первый донат, то купленное тобой количество кристаллов удвоится. Имеет смысл в первый раз закупать большое количество!\n\n"

	message += "⭕️Покеболы\n"
	message += "_Как увеличить лимит покеболов?_\n"
	message += "Вариантов несколько:\n1. Купить увеличение лимита покеболов до 24 за 100💎\n2. За каждый 30 покемем в покедексе ты получишь +1 к лимиту покеболов\n3. Отправь игровому боту /spam и получи ссылку. Каждый новый игрок получивший 2 уровень и пришедший по твоей ссылке принесет тебе 1 покебол и увеличение лимита покеболов на 1.\n4.  Перейти по ссылке: https://telegram.me/storebot?start=pokemembrobot, оставить отзыв и переслать переписку с @storebot в тех поддержку игры @PBhelp\nТак как за кристаллы ты увеличиваешь покедекс не НА 24, а ДО 24. То имеет смысл купить увеличение как можно раньше.\n"
	message += "_Какой лимит покеболов самый оптимальный?_\n"
	message += "Как показывает практика, лучший лимит (но и самый труднодостижимый) — 90 шаров. Достигнув его и заполнив целиком, вы сможете, запустив ТурбоЛовлю, сходить на 45 охот подряд.\n\n"

	message += "🐱Покемемы\n"
	message += "_Я поймал второго покемона, но не могу его взять в руку, почему?_\n"
	message += "Покемоны 4 и более уровней не могут повторяться в руке (т.е. должны быть разными).\n"
	message += "_А щит на покемемах что дает?_\n"
	message += "Чтобы поймать покемема с щитом, тебе нужно иметь атаку больше их щита.\n"
	message += "_Что такое MP?_\n"
	message += "Пока не знаем (в игре не используется).\n"
	message += "_Есть смысл дальше вкачивать покемемы?_\n"
	message += "/best\\_all поможет сориентироваться, хватает ли тебе атаки на топового покемема твоего уровня, или нет.\n"
	message += "_Какие покемоны для меня топ? Где ловить топов?_\n"
	message += "/best, /best\\_all\n"
	message += "_Что значат перед названием покемема ромбы?_\n"
	message += "🔸- Илитный: +15% ⚔️ атака. Максимальный уровень - 8.\n🔶- Супер Илитный: +25% ⚔️ атака. Максимальный уровень - 10.\n🔹- Либераха: +10% ⚔️ атаки покемемам с теми же элементами, что у Либерахи. Максимальный уровень - 8.\n🔷- Супер Либераха: +20% атаки покемемам с теми же элементами, что у Супер Либерахи. Максимальный уровень - 10.\n\n"

	message += "🥚 Яйца\n"
	message += "_Что такое яйцо и зачем оно?_\n"
	message += "Яйцо — это бонус, который ты получаешь в битвах лиг, из яйца дают покемема твоей стихии и на 1 грейд больше твоего лвла (максимум 9го), так же яйцо можно получить за ежедневную активность.\n"
	message += "_Сколько яиц я получу в битве?_\n"
	message += "1 яйцо за участие в бою, +2 яйца если мы проиграем и красным и синим. А также по +2 яйца если выиграем за явным преимуществом, если оппоненты выставят против нас в 20 раз меньше атаки. Максимум можно получить по 5 яиц с битвы.\n"
	message += "_Если у меня яйцо 96/96 и я, не открывая его апну лвл - оно станет 96/192 или 0/192?_\n"
	message += "Яйца остаются, увеличивается необходимое количество для открытия, в твоем случае станет 96/192\n"
	message += "_Если не открывать 🥚яйцо, когда оно наберется до нужного количества, оно не будет дальше расти?_\n"
	message += "Не будет."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
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
