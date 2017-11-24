// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pinnerinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// PinnerInterface implements Pinner for importing via appcontex
type PinnerInterface interface {
	Init()

	PinBattleAlert()
	PinMessageToAllChats(update *tgbotapi.Update) string
}
