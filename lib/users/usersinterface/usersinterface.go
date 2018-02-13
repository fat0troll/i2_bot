// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package usersinterface

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
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
	ProfileAddEffectsMessage(update *tgbotapi.Update) string
	ProfileMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	UsersList(update *tgbotapi.Update) string
}
