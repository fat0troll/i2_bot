// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/appcontext"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/datacache/datacacheinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"sync"
)

var (
	c *appcontext.Context
)

// DataCache is a function-handling struct for package datacache.
// Also, it's a data cache: it handles all data, powered by DataCache functions.
type DataCache struct {
	// Players — users of bot
	players      map[int]*dbmapping.Player
	playersMutex sync.Mutex
	// Profiles - game profiles, no matter, actual or not
	profiles      map[int]*dbmapping.Profile
	profilesMutex sync.Mutex
	// Current profiles - actual profiles for players, mostly used by bot
	// Note: int in this array for player ID, not for profile ID
	currentProfiles      map[int]*dbmapping.Profile
	currentProfilesMutex sync.Mutex
	// Pokememes
	pokememes      map[int]*dbmapping.Pokememe
	pokememesMutex sync.Mutex
	// Pokememes with all supported data
	fullPokememes      map[int]*dbmapping.PokememeFull
	fullPokememesMutex sync.Mutex

	// Elements
	elements      map[int]*dbmapping.Element
	elementsMutex sync.Mutex
	// Leagues
	leagues      map[int]*dbmapping.League
	leaguesMutex sync.Mutex
	// Locations
	locations      map[int]*dbmapping.Location
	locationsMutex sync.Mutex
	// Weapons
	weapons      map[int]*dbmapping.Weapon
	weaponsMutex sync.Mutex
}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	dc := &DataCache{}
	c.RegisterDataCacheInterface(datacacheinterface.DataCacheInterface(dc))
}

// Init is a initialization function for package
func (dc *DataCache) Init() {
	c.Log.Info("Initializing DataCache...")

	dc.initElements()
	dc.loadElements()
	dc.initLeagues()
	dc.loadLeagues()
	dc.initLocations()
	dc.loadLocations()
	dc.initWeapons()
	dc.loadWeapons()
	dc.initPokememes()
	dc.loadPokememes()
	dc.initPlayers()
	dc.loadPlayers()
	dc.initProfiles()
	dc.loadProfiles()
}
