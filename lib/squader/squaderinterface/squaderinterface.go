// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squaderinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// SquaderInterface implements Squader for importing via appcontext.
type SquaderInterface interface {
	Init()
	CreateSquad(update *tgbotapi.Update) string
	SquadsList(update *tgbotapi.Update) string
	SquadStatictics(squadID int) string
}
