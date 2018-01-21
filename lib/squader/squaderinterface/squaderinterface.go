// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squaderinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
)

// SquaderInterface implements Squader for importing via appcontext.
type SquaderInterface interface {
	Init()

	GetAllSquadChats() ([]dbmapping.Chat, bool)
	GetAllSquadFloodChats() ([]dbmapping.Chat, bool)
	GetAvailableSquadChatsForUser(playerRaw *dbmapping.Player) ([]dbmapping.Chat, bool)
	GetCommandersForSquadViaChat(chatRaw *dbmapping.Chat) ([]dbmapping.Player, bool)
	GetSquadByID(squadID int) (dbmapping.SquadChat, bool)
	GetSquadChatsBySquadsIDs(squadsID string) ([]dbmapping.Chat, bool)
	GetUserRolesInSquads(playerRaw *dbmapping.Player) ([]dbmapping.SquadPlayerFull, bool)
	IsChatASquadEnabled(chatRaw *dbmapping.Chat) string

	AddUserToSquad(update *tgbotapi.Update, adderRaw *dbmapping.Player) string
	CreateSquad(update *tgbotapi.Update) string

	SquadInfo(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	SquadsList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	CleanFlood(update *tgbotapi.Update, chatRaw *dbmapping.Chat) string
}
