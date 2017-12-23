// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squader

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// CleanFlood will clean flood from squads
func (s *Squader) CleanFlood(update *tgbotapi.Update, chatRaw *dbmapping.Chat) string {
	switch s.IsChatASquadEnabled(chatRaw) {
	case "main":
		talker, ok := c.Users.GetOrCreatePlayer(update.Message.From.ID)
		if !ok {
			s.deleteFloodMessage(update)
			return "fail"
		}
		if (update.Message.From.UserName != "i2_bot") && (update.Message.From.UserName != "i2_bot_dev") && !s.isUserAnyCommander(talker.ID) {
			s.deleteFloodMessage(update)
			return "fail"
		}
	}

	return "protection_passed"
}
