// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

// DeletePokememe deletes pokememe by its ID
func (p *Pokedexer) DeletePokememe(update *tgbotapi.Update) string {
	pokememeNum, _ := strconv.Atoi(update.Message.CommandArguments())
	if pokememeNum == 0 {
		return "fail"
	}

	err := c.DataCache.DeletePokememeByID(pokememeNum)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	message := "Покемем удалён."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
