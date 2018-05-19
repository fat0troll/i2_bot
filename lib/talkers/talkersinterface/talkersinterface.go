// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package talkersinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

// TalkersInterface implements Talkers for importing via appcontex
type TalkersInterface interface {
	Init()

	AcademyMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	BastionMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	GamesMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	HelpMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	FAQMessage(update *tgbotapi.Update) string
	RulesMessage(update *tgbotapi.Update) string

	AnyMessageUnauthorized(update *tgbotapi.Update) string
	BanError(update *tgbotapi.Update) string
	BotError(update *tgbotapi.Update) string

	LongMessage(update *tgbotapi.Update) string
	DurakMessage(update *tgbotapi.Update) string
	MatMessage(update *tgbotapi.Update) string

	NewYearMessage2018()
}
