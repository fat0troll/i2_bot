// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacacheinterface

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

// DataCacheInterface implements DataCache for importing via appcontext.
type DataCacheInterface interface {
	Init()

	AddPlayer(player *dbmapping.Player) (int, error)
	GetOrCreatePlayerByTelegramID(telegramID int) (*dbmapping.Player, error)
	GetPlayerByID(playerID int) (*dbmapping.Player, error)
	GetPlayerByTelegramID(telegramID int) (*dbmapping.Player, error)
	UpdatePlayerFields(player *dbmapping.Player) (*dbmapping.Player, error)
	UpdatePlayerTimestamp(playerID int) error

	AddProfile(profile *dbmapping.Profile) (int, error)
	GetPlayersWithCurrentProfiles() map[int]*dbmapping.PlayerProfile
	GetProfileByID(profileID int) (*dbmapping.Profile, error)
	GetProfileByPlayerID(playerID int) (*dbmapping.Profile, error)

	AddPokememe(pokememeData map[string]string, pokememeLocations map[string]string, pokememeElements map[string]string) (int, error)
	GetAllPokememes() map[int]*dbmapping.PokememeFull
	GetPokememeByID(pokememeID int) (*dbmapping.PokememeFull, error)
	GetPokememeByName(name string) (*dbmapping.PokememeFull, error)
	DeletePokememeByID(pokememeID int) error

	GetLeagueBySymbol(symbol string) (*dbmapping.League, error)

	GetWeaponTypeByID(weaponID int) (*dbmapping.Weapon, error)
	GetWeaponTypeByName(name string) (*dbmapping.Weapon, error)
}
