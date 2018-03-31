// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
	"strconv"
)

func (dc *DataCache) initLeagues() {
	c.Log.Info("Initializing Leagues storage...")
	dc.leagues = make(map[int]*datamapping.League)
}

func (dc *DataCache) loadLeagues() {
	c.Log.Info("Load current Leagues data to DataCache...")
	leagues := dc.getLeagues()

	for i := range leagues {
		dc.leagues[leagues[i].ID] = &leagues[i]
	}
	c.Log.Info("Loaded leagues in DataCache: " + strconv.Itoa(len(dc.leagues)))
}

func (dc *DataCache) getLeagues() []datamapping.League {
	leagues := []datamapping.League{}

	leagues = append(leagues, datamapping.League{1, "🈸", "ИНСТИНКТ"})
	leagues = append(leagues, datamapping.League{2, "🈳 ", "МИСТИКА"})
	leagues = append(leagues, datamapping.League{3, "🈵", "ОТВАГА"})

	return leagues
}

// External functions

// GetLeagueByID returns league from datacache by ID
func (dc *DataCache) GetLeagueByID(leagueID int) (*datamapping.League, error) {
	if dc.leagues[leagueID] != nil {
		return dc.leagues[leagueID], nil
	}

	return nil, errors.New("There is no league with ID = " + strconv.Itoa(leagueID))
}

// GetLeagueBySymbol returns league from datacache by emoji
func (dc *DataCache) GetLeagueBySymbol(symbol string) (*datamapping.League, error) {
	for i := range dc.leagues {
		if dc.leagues[i].Symbol == symbol {
			return dc.leagues[i], nil
		}
	}

	return nil, errors.New("There is no league with symbol = " + symbol)
}
