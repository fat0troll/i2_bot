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
	"../dbmapping"
)

func (t *Talkers) BestPokememesList(update tgbotapi.Update, player_raw dbmapping.Player) string {
	pokememes, ok := c.Getters.GetBestPokememes(player_raw.Id)
	if !ok {
		log.Printf("Cannot get pokememes from getter!")
		return "fail"
	}

	message := "*Лучшие покемемы для ловли*\n\n"
	for i := range pokememes {
		pk := pokememes[i].Pokememe
		pk_l := pokememes[i].Locations
		pk_e := pokememes[i].Elements
		message += strconv.Itoa(pk.Grade) + "⃣ "
		message += pk.Name + " (⚔"
		message += c.Parsers.ReturnPoints(pk.Attack)
		message += ", 🛡" + c.Parsers.ReturnPoints(pk.Defence) + ")"
		for i := range pk_e {
			message += pk_e[i].Symbol
		}
		message += " /pk" + strconv.Itoa(pk.Id) + "\n"
		message += "Локации: "
		for i := range pk_l {
			message += pk_l[i].Symbol + pk_l[i].Name
			if i+1 < len(pk_l) {
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
