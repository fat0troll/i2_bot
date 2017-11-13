// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package getters

import (
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// GetProfile returns last saved profile of player
func (g *Getters) GetProfile(playerID int) (dbmapping.Profile, bool) {
	profileRaw := dbmapping.Profile{}
	err := c.Db.Get(&profileRaw, c.Db.Rebind("SELECT * FROM profiles WHERE player_id=? ORDER BY created_at DESC LIMIT 1"), playerID)
	if err != nil {
		c.Log.Error(err)
		return profileRaw, false
	}

	return profileRaw, true
}
