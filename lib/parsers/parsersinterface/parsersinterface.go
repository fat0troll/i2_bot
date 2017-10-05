// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsersinterface

import (
    // local
    "../../dbmappings"
)


type ParsersInterface interface {
    ParsePokememe(text string, player_raw dbmappings.Players) string
}
