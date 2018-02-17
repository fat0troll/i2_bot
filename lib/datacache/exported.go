// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/appcontext"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datacache/datacacheinterface"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
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

	// Chats
	chats      map[int]*dbmapping.Chat
	chatsMutex sync.Mutex
	// Squads
	squads                map[int]*dbmapping.Squad
	squadsWithChats       map[int]*dbmapping.SquadChat
	squadPlayersRelations map[int]*dbmapping.SquadPlayer
	squadPlayers          map[int]map[int]*dbmapping.SquadPlayerFull
	squadsMutex           sync.Mutex

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
	dc.initChats()
	dc.loadChats()
	dc.initSquads()
	dc.loadSquads()
}
