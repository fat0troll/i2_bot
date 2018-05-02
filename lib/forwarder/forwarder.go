// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package forwarder

import (
	"regexp"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

// ProcessForward process forwards for single-user chats
func (f *Forwarder) ProcessForward(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	text := update.Message.Text

	// Forwards
	var pokememeMsg = regexp.MustCompile(`Dex(.+)\nGrade(.+)\nName(.+)`)
	var profileMsg = regexp.MustCompile(`id(\s)(\d+)\n(Team)(\s)([А-Я]+)\nName(\s)(.*)\nLvl(\s)(\d+)`)

	switch {
	case pokememeMsg.MatchString(text):
		c.Log.Debug("Pokememe posted!")
		if c.Users.PlayerBetterThan(playerRaw, "admin") {
			return c.Pokedexer.ParsePokememe(update, playerRaw)
		}

		return c.Talkers.AnyMessageUnauthorized(update)
	case profileMsg.MatchString(text):
		c.Log.Debug("Profile posted!")
		return c.Users.ParseProfile(update, playerRaw)
	default:
		c.Log.Debug(text)
	}

	return "fail"
}
