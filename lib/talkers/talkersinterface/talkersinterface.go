// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkersinterface

import (
    // 3rd party
	"gopkg.in/telegram-bot-api.v4"
    // local
    "../../dbmappings"
)

type TalkersInterface interface {
    Init()
    // Commands
    HelloMessageUnauthorized(update tgbotapi.Update)
    HelloMessageAuthorized(update tgbotapi.Update, player_raw dbmappings.Players)
    HelpMessage(update tgbotapi.Update)
    PokememesList(update tgbotapi.Update, page int)

    // Returns
    PokememeAddSuccessMessage(update tgbotapi.Update)
    PokememeAddDuplicateMessage(update tgbotapi.Update)
    PokememeAddFailureMessage(update tgbotapi.Update)

    // Errors
    AnyMessageUnauthorized(update tgbotapi.Update)

    // Easter eggs
    DurakMessage(update tgbotapi.Update)
    MatMessage(update tgbotapi.Update)
}
