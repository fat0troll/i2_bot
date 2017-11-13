// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package forwarder

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
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
		if playerRaw.LeagueID == 1 {
			status := c.Parsers.ParsePokememe(text, playerRaw)
			switch status {
			case "ok":
				c.Talkers.PokememeAddSuccessMessage(update)
				return "ok"
			case "dup":
				c.Talkers.PokememeAddDuplicateMessage(update)
				return "ok"
			case "fail":
				c.Talkers.PokememeAddFailureMessage(update)
				return "fail"
			}
		} else {
			c.Talkers.AnyMessageUnauthorized(update)
			return "fail"
		}
	case profileMsg.MatchString(text):
		c.Log.Debug("Profile posted!")
		status := c.Parsers.ParseProfile(update, playerRaw)
		switch status {
		case "ok":
			c.Talkers.ProfileAddSuccessMessage(update)
			return "ok"
		case "fail":
			c.Talkers.ProfileAddFailureMessage(update)
			return "fail"
		}
	default:
		c.Log.Debug(text)
	}

	return "fail"
}
