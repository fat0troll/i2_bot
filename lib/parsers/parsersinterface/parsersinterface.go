// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsersinterface

import (
    // 3rd party
    "github.com/go-telegram-bot-api/telegram-bot-api"
    // local
    "../../dbmappings"
)


type ParsersInterface interface {
    ParsePokememe(text string, player_raw dbmappings.Players) string
    ParseProfile(update tgbotapi.Update, player_raw dbmappings.Players) string
    ReturnPoints(points int) string
}
