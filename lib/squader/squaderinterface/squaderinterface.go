// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squaderinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// SquaderInterface implements Squader for importing via appcontext.
type SquaderInterface interface {
	Init()

	GetSquadByID(squadID int) (dbmapping.SquadChat, bool)
	GetUserRolesInSquads(playerRaw *dbmapping.Player) ([]dbmapping.SquadPlayerFull, bool)

	AddUserToSquad(update *tgbotapi.Update, adderRaw *dbmapping.Player) string
	CreateSquad(update *tgbotapi.Update) string

	SquadInfo(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	SquadsList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
}
