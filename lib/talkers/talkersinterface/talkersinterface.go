// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkersinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// TalkersInterface implements Talkers for importing via appcontex
type TalkersInterface interface {
	Init()
	HelloMessageUnauthorized(update *tgbotapi.Update)
	HelloMessageAuthorized(update *tgbotapi.Update, playerRaw *dbmapping.Player)
	HelpMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player)
	PokememesList(update *tgbotapi.Update)
	PokememeInfo(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	BestPokememesList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	PokememeAddSuccessMessage(update *tgbotapi.Update)
	PokememeAddDuplicateMessage(update *tgbotapi.Update)
	PokememeAddFailureMessage(update *tgbotapi.Update)
	ProfileAddSuccessMessage(update *tgbotapi.Update)
	ProfileAddFailureMessage(update *tgbotapi.Update)
	ProfileMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	AnyMessageUnauthorized(update *tgbotapi.Update)
	GetterError(update *tgbotapi.Update)

	AdminBroadcastMessageCompose(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	AdminBroadcastMessageSend(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	GroupsList(update *tgbotapi.Update) string

	DurakMessage(update *tgbotapi.Update)
	MatMessage(update *tgbotapi.Update)
}
