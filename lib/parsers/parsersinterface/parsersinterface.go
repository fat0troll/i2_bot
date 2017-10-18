// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsersinterface

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../../dbmapping"
)

type ParsersInterface interface {
	ParsePokememe(text string, player_raw dbmapping.Player) string
	ParseProfile(update tgbotapi.Update, player_raw dbmapping.Player) string
	ReturnPoints(points int) string
}
