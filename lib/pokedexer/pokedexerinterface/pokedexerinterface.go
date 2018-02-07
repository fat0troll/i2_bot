// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexerinterface

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// PokedexerInterface implements Pokedexer for importing via appcontext.
type PokedexerInterface interface {
	ParsePokememe(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	PokememesList(update *tgbotapi.Update)
	PokememeInfo(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	AdvicePokememesList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	DeletePokememe(update *tgbotapi.Update) string
}
