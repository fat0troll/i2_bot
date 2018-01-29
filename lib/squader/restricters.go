// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squader

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// CleanFlood will clean flood from squads
func (s *Squader) CleanFlood(update *tgbotapi.Update, chatRaw *dbmapping.Chat) string {
	switch s.IsChatASquadEnabled(chatRaw) {
	case "main":
		talker, err := c.DataCache.GetPlayerByTelegramID(update.Message.From.ID)
		if err != nil {
			c.Log.Error(err.Error())
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
