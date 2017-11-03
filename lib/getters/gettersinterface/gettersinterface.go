// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package gettersinterface

import (
	// 3rd-party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// GettersInterface implements Getters for importing via appcontext.
type GettersInterface interface {
	Init()
	CreateBroadcastMessage(playerRaw *dbmapping.Player, messageBody string, broadcastType string) (dbmapping.Broadcast, bool)
	GetBroadcastMessageByID(messageID int) (dbmapping.Broadcast, bool)
	UpdateBroadcastMessageStatus(messageID int, messageStatus string) (dbmapping.Broadcast, bool)
	GetOrCreateChat(update *tgbotapi.Update) (dbmapping.Chat, bool)
	GetChatByID(chatID int64) (dbmapping.Chat, bool)
	GetAllPrivateChats() ([]dbmapping.Chat, bool)
	UpdateChatTitle(chatRaw dbmapping.Chat, newTitle string) (dbmapping.Chat, bool)
	GetOrCreatePlayer(telegramID int) (dbmapping.Player, bool)
	GetPlayerByID(playerID int) (dbmapping.Player, bool)
	PlayerBetterThan(playerRaw *dbmapping.Player, powerLevel string) bool
	GetProfile(playerID int) (dbmapping.Profile, bool)
	GetPokememes() ([]dbmapping.PokememeFull, bool)
	GetBestPokememes(playerID int) ([]dbmapping.PokememeFull, bool)
	GetPokememeByID(pokememeID string) (dbmapping.PokememeFull, bool)
	PossibilityRequiredPokeballs(location int, grade int, lvl int) (float64, int)
}
