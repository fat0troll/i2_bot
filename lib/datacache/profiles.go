// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"strconv"
)

func (dc *DataCache) initProfiles() {
	c.Log.Info("Initializing Profiles storage...")
	dc.profiles = make(map[int]*dbmapping.Profile)
	dc.currentProfiles = make(map[int]*dbmapping.Profile)
}

func (dc *DataCache) loadProfiles() {
	c.Log.Info("Load current Profiles data from database to DataCache...")
	profiles := []dbmapping.Profile{}
	err := c.Db.Select(&profiles, "SELECT * FROM profiles")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}

	dc.profilesMutex.Lock()
	dc.currentProfilesMutex.Lock()
	for i := range profiles {
		dc.profiles[profiles[i].ID] = &profiles[i]

		// Filling current profiles
		if dc.currentProfiles[profiles[i].PlayerID] != nil {
			if dc.currentProfiles[profiles[i].PlayerID].CreatedAt.Before(profiles[i].CreatedAt) {
				dc.currentProfiles[profiles[i].PlayerID] = &profiles[i]
			}
		}

		dc.currentProfiles[profiles[i].PlayerID] = &profiles[i]
	}
	c.Log.Info("Loaded profiles in DataCache: " + strconv.Itoa(len(dc.profiles)))
	c.Log.Info("Loaded current profiles in DataCache: " + strconv.Itoa(len(dc.currentProfiles)))
	dc.profilesMutex.Unlock()
	dc.currentProfilesMutex.Unlock()
}

// External functions

// AddProfile creates new profile in database
func (dc *DataCache) AddProfile(profile *dbmapping.Profile) (int, error) {
	_, err := c.Db.NamedExec("INSERT INTO `profiles` VALUES(NULL, :player_id, :nickname, :telegram_nickname, :level_id, :pokeballs, :wealth, :pokememes_wealth, :exp, :egg_exp, :power, :weapon_id, :crystalls, :created_at)", &profile)
	if err != nil {
		c.Log.Error(err.Error())
		return 0, err
	}
	insertedProfile := dbmapping.Profile{}
	err = c.Db.Get(&insertedProfile, c.Db.Rebind("SELECT * FROM profiles WHERE player_id=? AND created_at=?"), profile.PlayerID, profile.CreatedAt)
	if err != nil {
		c.Log.Error(err.Error())
		return 0, err
	}

	dc.profilesMutex.Lock()
	dc.currentProfilesMutex.Lock()
	dc.profiles[insertedProfile.ID] = &insertedProfile
	dc.currentProfiles[insertedProfile.PlayerID] = &insertedProfile
	dc.currentProfilesMutex.Unlock()
	dc.profilesMutex.Unlock()

	return insertedProfile.ID, nil
}

// GetPlayersWithCurrentProfiles returns users with profiles, if exist
func (dc *DataCache) GetPlayersWithCurrentProfiles() map[int]*dbmapping.PlayerProfile {
	dc.currentProfilesMutex.Lock()
	dc.playersMutex.Lock()

	users := make(map[int]*dbmapping.PlayerProfile)

	for i := range dc.players {
		user := dbmapping.PlayerProfile{}
		user.Player = *dc.players[i]
		if dc.players[i].LeagueID != 0 {
			user.League = *dc.leagues[dc.players[i].LeagueID]
		}

		if dc.currentProfiles[dc.players[i].ID] != nil {
			user.HaveProfile = true
			user.Profile = *dc.currentProfiles[dc.players[i].ID]
		} else {
			user.HaveProfile = false
		}

		users[dc.players[i].ID] = &user
	}

	dc.currentProfilesMutex.Unlock()
	dc.playersMutex.Unlock()
	return users
}

// GetProfileByID returns profile from datacache by ID
func (dc *DataCache) GetProfileByID(profileID int) (*dbmapping.Profile, error) {
	if dc.profiles[profileID] != nil {
		return dc.profiles[profileID], nil
	}

	return nil, errors.New("There is no profile with ID = " + strconv.Itoa(profileID))
}

// GetProfileByPlayerID returns current profile for player
func (dc *DataCache) GetProfileByPlayerID(playerID int) (*dbmapping.Profile, error) {
	if dc.currentProfiles[playerID] != nil {
		c.Log.Debug("DataCache: found current profile for player with ID = " + strconv.Itoa(playerID))
		return dc.currentProfiles[playerID], nil
	}

	return nil, errors.New("There is no profile for user with ID = " + strconv.Itoa(playerID))
}
