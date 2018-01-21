// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package forwarder

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"regexp"
)

// ProcessForward process forwards for single-user chats
func (f *Forwarder) ProcessForward(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	text := update.Message.Text
	// Forwards
	var pokememeMsg = regexp.MustCompile("(Уровень)(.+)(Опыт)(.+)\n(Элементы:)(.+)\n(.+)(💙MP)")
	var profileMsg = regexp.MustCompile(`(Онлайн: )(\d+)\n(Турнир через)(.+)\n\n((.*)\n|(.*)\n(.*)\n)(Элементы)(.+)\n(.*)\n\n(.+)(Уровень)(.+)\n`)

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
