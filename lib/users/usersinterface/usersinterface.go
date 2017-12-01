// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package usersinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// UsersInterface implements Users for importing via appcontex
type UsersInterface interface {
	Init()

	ParseProfile(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	GetProfile(playerID int) (dbmapping.Profile, bool)
	GetOrCreatePlayer(telegramID int) (dbmapping.Player, bool)
	GetPlayerByID(playerID int) (dbmapping.Player, bool)
	PlayerBetterThan(playerRaw *dbmapping.Player, powerLevel string) bool

	ForeignProfileMessage(update *tgbotapi.Update) string
	FormatUsername(userName string) string
	ProfileMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	UsersList(update *tgbotapi.Update) string
}
