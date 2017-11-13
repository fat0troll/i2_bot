// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package forwarderinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// ForwarderInterface implements Getters for importing via appcontext.
type ForwarderInterface interface {
	Init()
	ProcessForward(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
}
