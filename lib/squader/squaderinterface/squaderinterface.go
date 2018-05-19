// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squaderinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

// SquaderInterface implements Squader for importing via appcontext.
type SquaderInterface interface {
	Init()

	AddUserToSquad(update *tgbotapi.Update, adderRaw *dbmapping.Player) string

	SquadInfo(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	SquadsList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
}
