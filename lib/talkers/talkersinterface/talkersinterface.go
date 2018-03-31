// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package talkersinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

// TalkersInterface implements Talkers for importing via appcontex
type TalkersInterface interface {
	Init()

	AcademyMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player)
	BastionMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player)
	HelpMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player)
	FAQMessage(update *tgbotapi.Update) string

	AnyMessageUnauthorized(update *tgbotapi.Update) string
	BanError(update *tgbotapi.Update) string
	BotError(update *tgbotapi.Update) string

	LongMessage(update *tgbotapi.Update) string
	DurakMessage(update *tgbotapi.Update) string
	MatMessage(update *tgbotapi.Update) string

	NewYearMessage2018()
}
