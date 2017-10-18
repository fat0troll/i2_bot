// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package gettersinterface

import (
	// local
	"../../dbmapping"
)

// GettersInterface implements Getters for importing via appcontext.
type GettersInterface interface {
	Init()
	GetOrCreatePlayer(telegram_id int) (dbmapping.Player, bool)
	GetPlayerByID(player_id int) (dbmapping.Player, bool)
	GetProfile(player_id int) (dbmapping.Profile, bool)
	GetPokememes() ([]dbmapping.PokememeFull, bool)
	GetBestPokememes(player_id int) ([]dbmapping.PokememeFull, bool)
	GetPokememeByID(pokememe_id string) (dbmapping.PokememeFull, bool)
	PossibilityRequiredPokeballs(location int, grade int, lvl int) (float64, int)
}
