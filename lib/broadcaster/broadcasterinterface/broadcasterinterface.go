// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package broadcasterinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
)

// BroadcasterInterface implements Broadcaster for importing via appcontex
type BroadcasterInterface interface {
	Init()

	AdminBroadcastMessageCompose(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	AdminBroadcastMessageSend(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
}
