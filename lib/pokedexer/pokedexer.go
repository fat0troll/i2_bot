// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"sort"
	"strconv"
)

func (p *Pokedexer) pokememesListing(update *tgbotapi.Update, page int, pokememesArray map[int]*dbmapping.PokememeFull) {
	message := "*Известные боту покемемы*\n"
	message += "Список отсортирован по грейду и алфавиту.\n"
	message += "Покедекс: " + strconv.Itoa(len(pokememesArray)) + " / 296\n"
	message += "Отображаем покемемов с " + strconv.Itoa(((page-1)*50)+1) + " по " + strconv.Itoa(page*50) + "\n"
	if len(pokememesArray) > page*50 {
		message += "Переход на следующую страницу: /pokedeks" + strconv.Itoa(page+1)
	}
	if page > 1 {
		message += "\nПереход на предыдущую страницу: /pokedeks" + strconv.Itoa(page-1)
	}
	message += "\n\n"

	var keys []int
	for i := range pokememesArray {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, i := range keys {
		if (i+1 > 50*(page-1)) && (i+1 < (50*page)+1) {
			pk := pokememesArray[i].Pokememe
			pkE := pokememesArray[i].Elements
			message += strconv.Itoa(i+1) + ". " + strconv.Itoa(pk.Grade)
			message += "⃣ *" + pk.Name
			message += "* (" + c.Statistics.GetPrintablePoints(pk.HP) + "-" + c.Statistics.GetPrintablePoints(pk.MP) + ") ⚔️ *"
			message += c.Statistics.GetPrintablePoints(pk.Attack) + "* \\["
			for j := range pkE {
				message += pkE[j].Symbol
			}
			message += "] " + c.Statistics.GetPrintablePoints(pk.Price) + "$ /pk" + strconv.Itoa(pk.ID)
			message += "\n"
		}
	}

	if len(pokememesArray) > page*50 {
		message += "\n"
		message += "Переход на следующую страницу: /pokedeks" + strconv.Itoa(page+1)
	}
	if page > 1 {
		message += "\nПереход на предыдущую страницу: /pokedeks" + strconv.Itoa(page-1)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

func (p *Pokedexer) pokememeAddSuccessMessage(update *tgbotapi.Update, newPokememeID int) {
	message := "*Покемем успешно добавлен.*\n\n"
	message += "Посмотреть всех известных боту покемемов можно командой /pokedeks\n"
	message += "Посмотреть свежедобавленного покемема можно командой /pk" + strconv.Itoa(newPokememeID)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

func (p *Pokedexer) pokememeAddDuplicateMessage(update *tgbotapi.Update) {
	message := "*Мы уже знаем об этом покемеме*\n\n"
	message += "Посмотреть всех известных боту покемемов можно командой /pokedeks\n\n"
	message += "Если у покемема изменились описание или характеристики, напиши @fat0troll для обновления базы."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

func (p *Pokedexer) pokememeAddFailureMessage(update *tgbotapi.Update) {
	message := "*Неудачно получилось :(*\n\n"
	message += "Случилась жуткая ошибка, и мы не смогли записать покемема в базу. Напиши @fat0troll, он разберется.\n\n"
	message += "Посмотреть всех известных боту покемемов можно командой /pokedeks"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
