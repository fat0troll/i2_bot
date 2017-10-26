// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// stdlib
	"log"
	"strconv"
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// BestPokememesList shows list for catching based on player league and grade
func (t *Talkers) BestPokememesList(update tgbotapi.Update, playerRaw dbmapping.Player) string {
	pokememes, ok := c.Getters.GetBestPokememes(playerRaw.ID)
	if !ok {
		log.Printf("Cannot get pokememes from getter!")
		return "fail"
	}

	message := "*Лучшие покемемы для ловли*\n\n"
	for i := range pokememes {
		pk := pokememes[i].Pokememe
		pkL := pokememes[i].Locations
		pkE := pokememes[i].Elements
		message += strconv.Itoa(pk.Grade) + "⃣ "
		message += pk.Name + " (⚔"
		message += c.Parsers.ReturnPoints(pk.Attack)
		message += ", 🛡" + c.Parsers.ReturnPoints(pk.Defence) + ")"
		for i := range pkE {
			message += pkE[i].Symbol
		}
		message += " /pk" + strconv.Itoa(pk.ID) + "\n"
		message += "Локации: "
		for i := range pkL {
			message += pkL[i].Symbol + pkL[i].Name
			if i+1 < len(pkL) {
				message += ", "
			}
		}
		message += "\nКупить: "
		if pk.Purchaseable {
			message += "💲" + c.Parsers.ReturnPoints(pk.Price*3)
		} else {
			message += "Нельзя"
		}
		message += "\n\n"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
