// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"sync"

	"source.wtfteam.pro/i2_bot/i2_bot/lib/appcontext"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datacache/datacacheinterface"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
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

	// Chats
	chats      map[int]*dbmapping.Chat
	chatsMutex sync.Mutex
	// Squads
	squads                map[int]*dbmapping.Squad
	squadsWithChats       map[int]*dbmapping.SquadChat
	squadPlayersRelations map[int]*dbmapping.SquadPlayer
	squadPlayers          map[int]map[int]*dbmapping.SquadPlayerFull
	squadsMutex           sync.Mutex

	// Rarely changing data
	// Elements
	elements map[int]*datamapping.Element
	// Leagues
	leagues map[int]*datamapping.League
	// Levels
	levels map[int]*datamapping.Level
	// Locations
	locations map[int]*datamapping.Location
	// Pokememes
	pokememes              map[int]*datamapping.Pokememe
	fullPokememes          map[int]*datamapping.PokememeFull
	pokememesGradeLocation map[int]map[int]int
	// Weapons
	weapons map[int]*datamapping.Weapon
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
	dc.initLevels()
	dc.loadLevels()
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
