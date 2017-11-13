// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsersinterface

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// ParsersInterface implements Parsers for importing via appcontext.
type ParsersInterface interface {
	ParsePokememe(text string, playerRaw *dbmapping.Player) string
	ParseProfile(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	ReturnPoints(points int) string
}
