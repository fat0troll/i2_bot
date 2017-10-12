// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package gettersinterface

type GettersInterface interface {
    Init()
    // Possibilities
    PossibilityRequiredPokeballs(location int, grade int, lvl int) int
}
