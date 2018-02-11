// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkersinterface

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// TalkersInterface implements Talkers for importing via appcontex
type TalkersInterface interface {
	Init()

	AcademyMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player)
	BastionMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player)
	HelpMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player)
	FiveOffer(update *tgbotapi.Update) string

	AnyMessageUnauthorized(update *tgbotapi.Update) string
	BanError(update *tgbotapi.Update) string
	BotError(update *tgbotapi.Update) string

	LongMessage(update *tgbotapi.Update) string
	DurakMessage(update *tgbotapi.Update) string
	MatMessage(update *tgbotapi.Update) string

	NewYearMessage2018()
}
