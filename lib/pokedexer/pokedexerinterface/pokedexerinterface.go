// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexerinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

// PokedexerInterface implements Pokedexer for importing via appcontext.
type PokedexerInterface interface {
	ParsePokememe(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	PokememesList(update *tgbotapi.Update)
	PokememesListUpdater(update *tgbotapi.Update) string
	PokememeInfo(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	AdvicePokememesList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
}
