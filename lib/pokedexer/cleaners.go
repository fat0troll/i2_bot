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

	pokememe, ok := p.GetPokememeByID(strconv.Itoa(pokememeNum))
	if !ok {
		return "fail"
	}

	_, err := c.Db.NamedExec("DELETE FROM pokememes WHERE id=:id", &pokememe.Pokememe)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	_, err = c.Db.NamedExec("DELETE FROM pokememes_elements WHERE pokememe_id=:id", &pokememe.Pokememe)
	if err != nil {
		c.Log.Debug(err.Error())
	}

	_, err = c.Db.NamedExec("DELETE FROM pokememes_locations WHERE pokememe_id=:id", &pokememe.Pokememe)
	if err != nil {
		c.Log.Debug(err.Error())
	}

	_, err = c.Db.NamedExec("DELETE FROM profiles_pokememes WHERE pokememe_id=:id", &pokememe.Pokememe)
	if err != nil {
		c.Log.Debug(err.Error())
	}

	message := "Покемем удалён."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
