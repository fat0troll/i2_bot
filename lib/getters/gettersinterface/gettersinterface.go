// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package gettersinterface

import (
	// 3rd-party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../../dbmapping"
)

// GettersInterface implements Getters for importing via appcontext.
type GettersInterface interface {
	Init()
	GetOrCreateChat(update *tgbotapi.Update) (dbmapping.Chat, bool)
	GetChatByID(chatID int) (dbmapping.Chat, bool)
	GetOrCreatePlayer(telegramID int) (dbmapping.Player, bool)
	GetPlayerByID(playerID int) (dbmapping.Player, bool)
	GetProfile(playerID int) (dbmapping.Profile, bool)
	GetPokememes() ([]dbmapping.PokememeFull, bool)
	GetBestPokememes(playerID int) ([]dbmapping.PokememeFull, bool)
	GetPokememeByID(pokememeID string) (dbmapping.PokememeFull, bool)
	PossibilityRequiredPokeballs(location int, grade int, lvl int) (float64, int)
}
