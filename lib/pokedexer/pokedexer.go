// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"sort"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
	"strconv"
)

func (p *Pokedexer) pokememesListingMessage(update *tgbotapi.Update, page int, pokememesArray map[int]*datamapping.PokememeFull) string {
	message := "📕*Покедекс: " + strconv.Itoa(len(pokememesArray)) + " / 733*\n"
	message += "```\nВсе виды покемемов, которые известны боту. [" + strconv.Itoa(page) + "] (" + strconv.Itoa(((page-1)*35)+1) + "-" + strconv.Itoa(page*35) + ")```"

	var keys []int
	for i := range pokememesArray {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, i := range keys {
		if (i > 35*(page-1)) && (i < (35*page)+1) {
			pk := pokememesArray[i].Pokememe
			pkE := pokememesArray[i].Elements
			message += strconv.Itoa(pk.ID) + ". *[" + strconv.Itoa(pk.Grade)
			message += "]* *" + pk.Name
			message += "* ❤️" + c.Statistics.GetPrintablePoints(pk.HP) + " ⚔️ "
			message += c.Statistics.GetPrintablePoints(pk.Attack) + " 🛡" + c.Statistics.GetPrintablePoints(pk.Defence) + " \\["
			for j := range pkE {
				message += pkE[j].Symbol
			}
			message += "] " + c.Statistics.GetPrintablePoints(pk.Price) + "$ /pk" + strconv.Itoa(pk.ID)
			message += "\n"
		}
	}

	return message
}

func (p *Pokedexer) pokememesListingKeyboard(pokememesArray map[int]*datamapping.PokememeFull) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	rows := make(map[int][]tgbotapi.InlineKeyboardButton)
	rowsCount := int(len(pokememesArray) / (35 * 7))
	for i := 0; i <= rowsCount; i++ {
		rows[i] = []tgbotapi.InlineKeyboardButton{}
	}
	totalPages := int(len(pokememesArray)/35) + 1
	for i := 1; i <= totalPages; i++ {
		btn := tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(i), "pokedeks"+strconv.Itoa(i))
		rows[(i-1)/7] = append(rows[(i-1)/7], btn)
	}
	for i := 0; i <= rowsCount; i++ {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rows[i])
	}

	return &keyboard
}

func (p *Pokedexer) pokememesListing(update *tgbotapi.Update, page int, pokememesArray map[int]*datamapping.PokememeFull) {
	message := p.pokememesListingMessage(update, page, pokememesArray)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = p.pokememesListingKeyboard(pokememesArray)

	c.Bot.Send(msg)
}

func (p *Pokedexer) pokememesListingUpdate(update *tgbotapi.Update, page int, pokememesArray map[int]*datamapping.PokememeFull) {
	message := p.pokememesListingMessage(update, page, pokememesArray)

	messageUpdate := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, message)
	messageUpdate.ParseMode = "Markdown"
	messageUpdate.ReplyMarkup = p.pokememesListingKeyboard(pokememesArray)

	c.Bot.Send(messageUpdate)
}
