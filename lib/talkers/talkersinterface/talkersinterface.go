// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkersinterface

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../../dbmapping"
)

// TalkersInterface implements Talkers for importing via appcontex
type TalkersInterface interface {
	Init()
	HelloMessageUnauthorized(update tgbotapi.Update)
	HelloMessageAuthorized(update tgbotapi.Update, playerRaw dbmapping.Player)
	HelpMessage(update tgbotapi.Update, playerRaw *dbmapping.Player)
	PokememesList(update tgbotapi.Update, page int)
	PokememeInfo(update tgbotapi.Update, playerRaw dbmapping.Player) string
	BestPokememesList(update tgbotapi.Update, playerRaw dbmapping.Player) string

	PokememeAddSuccessMessage(update tgbotapi.Update)
	PokememeAddDuplicateMessage(update tgbotapi.Update)
	PokememeAddFailureMessage(update tgbotapi.Update)
	ProfileAddSuccessMessage(update tgbotapi.Update)
	ProfileAddFailureMessage(update tgbotapi.Update)
	ProfileMessage(update tgbotapi.Update, playerRaw dbmapping.Player) string

	AnyMessageUnauthorized(update tgbotapi.Update)
	GetterError(update tgbotapi.Update)

	AdminBroadcastMessage(update tgbotapi.Update) string

	DurakMessage(update tgbotapi.Update)
	MatMessage(update tgbotapi.Update)
}
