// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
)

func (dc *DataCache) initLeagues() {
	c.Log.Info("Initializing Leagues storage...")
	dc.leagues = make(map[int]*dbmapping.League)
}

func (dc *DataCache) loadLeagues() {
	c.Log.Info("Load current Leagues data from database to DataCache...")
	leagues := []dbmapping.League{}
	err := c.Db.Select(&leagues, "SELECT * FROM leagues")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}

	dc.leaguesMutex.Lock()
	for i := range leagues {
		dc.leagues[leagues[i].ID] = &leagues[i]
	}
	c.Log.Info("Loaded leagues in DataCache: " + strconv.Itoa(len(dc.leagues)))
	dc.leaguesMutex.Unlock()
}

// External functions

// GetLeagueBySymbol returns league from datacache by emoji
func (dc *DataCache) GetLeagueBySymbol(symbol string) (*dbmapping.League, error) {
	dc.leaguesMutex.Lock()
	for i := range dc.leagues {
		if dc.leagues[i].Symbol == symbol {
			dc.leaguesMutex.Unlock()
			return dc.leagues[i], nil
		}
	}

	dc.leaguesMutex.Unlock()
	return nil, errors.New("There is no league with symbol = " + symbol)
}
