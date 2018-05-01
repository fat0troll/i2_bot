// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
)

// AdvicePokememesList shows list for catching
// It may be list of best or most valuable pokememes
func (p *Pokedexer) AdvicePokememesList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	pokememes, ok := p.getAdvicePokememes(playerRaw.ID, update.Message.Command())
	if !ok {
		c.Log.Error("Cannot get pokememes from getter!")
		return "fail"
	}

	profileRaw, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	message := ""
	switch update.Message.Command() {
	case "best":
		message += "*Пять топовых покемемов для поимки*\n\n"
	case "advice":
		message += "*Пять самых дорогих покемемов*\n\n"
	case "best_all":
		message += "*Все топовые покемемы для твоего уровня*\n\n"
	case "advice_all":
		message += "*Все самые дорогие покемемы для твоего уровня*\n\n"
	case "best_nofilter":
		message += "*Пять топовых покемемов для поимки без фильтра по элементам*\n\n"
	}
	for i := range pokememes {
		pk := pokememes[i].Pokememe
		pkL := pokememes[i].Locations
		pkE := pokememes[i].Elements
		message += "*[" + strconv.Itoa(pk.Grade) + "]* "
		message += pk.Name + " (⚔"
		message += c.Statistics.GetPrintablePoints(pk.Attack)
		message += ", 🛡" + c.Statistics.GetPrintablePoints(pk.Defence) + ")"
		for i := range pkE {
			message += pkE[i].Symbol
		}
		message += " /pk" + strconv.Itoa(pk.ID) + "\nЛокации: "
		for i := range pkL {
			message += pkL[i].Symbol + pkL[i].Name
			_, balls := c.Statistics.PossibilityRequiredPokeballs(pkL[i].ID, pk.Grade, profileRaw.LevelID)
			message += " ⭕" + strconv.Itoa(balls)
			if i+1 < len(pkL) {
				message += ", "
			}
		}
		message += "\nКупить: "
		if pk.Purchaseable {
			message += "💲" + c.Statistics.GetPrintablePoints(pk.Price*3)
		} else {
			message += "Нельзя"
		}
		if update.Message.Command() == "advice" || update.Message.Command() == "advice_all" {
			message += "\nСтоимость продажи: 💲" + c.Statistics.GetPrintablePoints(pk.Price)
		}
		if len(message) > 4000 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			msg.ParseMode = "Markdown"

			c.Bot.Send(msg)

			message = ""
		} else {
			message += "\n\n"
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

// PokememesList lists all known pokememes
func (p *Pokedexer) PokememesList(update *tgbotapi.Update) {
	pageNumber := strings.Replace(update.Message.Text, "/pokedex", "", 1)
	pageNumber = strings.Replace(pageNumber, "/pokedeks", "", 1)
	page, _ := strconv.Atoi(pageNumber)
	if page == 0 {
		page = 1
	}
	pokememesArray := c.DataCache.GetAllPokememes()
	p.pokememesListing(update, page, pokememesArray)
}

// PokememesListUpdater updates page in pokedeks message
func (p *Pokedexer) PokememesListUpdater(update *tgbotapi.Update) string {
	pageNumber := strings.Replace(update.CallbackQuery.Data, "pokedeks", "", 1)
	page, _ := strconv.Atoi(pageNumber)
	if page == 0 {
		page = 1
	}
	pokememesArray := c.DataCache.GetAllPokememes()
	p.pokememesListingUpdate(update, page, pokememesArray)

	return "ok"
}

// PokememeInfo shows information about single pokememe based on internal ID
func (p *Pokedexer) PokememeInfo(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	pokememeNumber := strings.Replace(update.Message.Text, "/pk", "", 1)
	var calculatePossibilites = true
	profileRaw, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		calculatePossibilites = false
	}

	pokememeID, _ := strconv.Atoi(pokememeNumber)
	pokememe, err := c.DataCache.GetPokememeByID(pokememeID)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	pk := pokememe.Pokememe

	message := "*[" + strconv.Itoa(pk.Grade) + "]* *" + pk.Name + "*\n"
	message += pk.Description + "\n\n"
	message += "Элементы:"
	for i := range pokememe.Elements {
		message += " " + pokememe.Elements[i].Symbol
	}
	message += "\n⚔ Атака: *" + c.Statistics.GetPrintablePoints(pk.Attack)
	message += "*\n❤️ HP: *" + c.Statistics.GetPrintablePoints(pk.HP)
	message += "*\n💙 MP: *" + c.Statistics.GetPrintablePoints(pk.MP)
	if pk.Defence != pk.Attack {
		message += "*\n🛡Защита: *" + c.Statistics.GetPrintablePoints(pk.Defence) + "* _(сопротивляемость покемема к поимке)_"
	} else {
		message += "*"
	}
	message += "\nСтоимость: *" + c.Statistics.GetPrintablePoints(pk.Price)
	message += "*\nКупить: *"
	if pk.Purchaseable {
		message += "Можно"
	} else {
		message += "Нельзя"
	}
	message += "*\nОбитает:"
	for i := range pokememe.Locations {
		message += " *" + pokememe.Locations[i].Name + "*"
		if (i + 1) < len(pokememe.Locations) {
			message += ","
		}
	}

	if c.Users.PlayerBetterThan(playerRaw, "academic") {
		message += "\nВероятность поимки:"
		idx := 1
		for idx < 23 {
			levelHeaderPosted := false
			for i := range pokememe.Locations {
				percentile, pokeballs := c.Statistics.PossibilityRequiredPokeballs(pokememe.Locations[i].ID, pk.Grade, idx)
				if pokeballs > 0 {
					if !levelHeaderPosted {
						message += "\nУровень: " + strconv.Itoa(idx)
						levelHeaderPosted = true
					}
					message += "\n" + pokememe.Locations[i].Name + " – "
					message += strconv.FormatFloat(percentile, 'f', 2, 64) + "% или "
					message += strconv.Itoa(pokeballs) + "⭕"
				}
			}
			idx++
		}
	} else if calculatePossibilites {
		if (pk.Grade < profileRaw.LevelID+2) && (pk.Grade > profileRaw.LevelID-3) {
			message += "\nВероятность поимки:"
			for i := range pokememe.Locations {
				percentile, pokeballs := c.Statistics.PossibilityRequiredPokeballs(pokememe.Locations[i].ID, pk.Grade, profileRaw.LevelID)
				message += "\n" + pokememe.Locations[i].Name + " – "
				message += strconv.FormatFloat(percentile, 'f', 2, 64) + "% или "
				message += strconv.Itoa(pokeballs) + "⭕"
			}
		}
	}

	message += "\n" + pk.ImageURL

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for i := range pokememe.Locations {
		c.Log.Info("wow, location")
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonSwitch(pokememe.Locations[i].Symbol+pokememe.Locations[i].Name, pokememe.Locations[i].Symbol+pokememe.Locations[i].Name)
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	msg.ReplyMarkup = keyboard
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
