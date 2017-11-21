// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexerinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// PokedexerInterface implements Pokedexer for importing via appcontext.
type PokedexerInterface interface {
	ParsePokememe(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	PokememesList(update *tgbotapi.Update)
	PokememeInfo(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	BestPokememesList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	GetPokememes() ([]dbmapping.PokememeFull, bool)
	GetPokememeByID(pokememeID string) (dbmapping.PokememeFull, bool)
}
