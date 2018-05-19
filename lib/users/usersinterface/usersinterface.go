// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package usersinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

// UsersInterface implements Users for importing via appcontex
type UsersInterface interface {
	Init()

	ParseProfile(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	GetPrettyName(user *tgbotapi.User) string
	PlayerBetterThan(playerRaw *dbmapping.Player, powerLevel string) bool

	FindByLevel(update *tgbotapi.Update) string
	FindByName(update *tgbotapi.Update) string
	FindByTopAttack(update *tgbotapi.Update) string
	ForeignProfileMessage(update *tgbotapi.Update) string
	FormatUsername(userName string) string
	ProfileMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	UsersList(update *tgbotapi.Update) string
}
