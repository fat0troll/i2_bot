// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkersinterface

import (
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../../dbmapping"
)

type TalkersInterface interface {
	Init()
	// Commands
	HelloMessageUnauthorized(update tgbotapi.Update)
	HelloMessageAuthorized(update tgbotapi.Update, player_raw dbmapping.Player)
	HelpMessage(update tgbotapi.Update)
	PokememesList(update tgbotapi.Update, page int)
	PokememeInfo(update tgbotapi.Update, player_raw dbmapping.Player) string
	BestPokememesList(update tgbotapi.Update, player_raw dbmapping.Player) string

	// Returns
	PokememeAddSuccessMessage(update tgbotapi.Update)
	PokememeAddDuplicateMessage(update tgbotapi.Update)
	PokememeAddFailureMessage(update tgbotapi.Update)
	ProfileAddSuccessMessage(update tgbotapi.Update)
	ProfileAddFailureMessage(update tgbotapi.Update)
	ProfileMessage(update tgbotapi.Update, player_raw dbmapping.Player) string

	// Errors
	AnyMessageUnauthorized(update tgbotapi.Update)
	GetterError(update tgbotapi.Update)

	// Easter eggs
	DurakMessage(update tgbotapi.Update)
	MatMessage(update tgbotapi.Update)
}
