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
	HelpMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player)

	AnyMessageUnauthorized(update *tgbotapi.Update) string
	BotError(update *tgbotapi.Update) string

	LongMessage(update *tgbotapi.Update)
	DurakMessage(update *tgbotapi.Update)
	MatMessage(update *tgbotapi.Update)
}
