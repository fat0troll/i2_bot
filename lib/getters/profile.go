// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package getters

import (
    // stdlib
    "log"
    // local
    "../dbmapping"
)

func (g *Getters) GetProfile(player_id int) (dbmapping.Profile, bool) {
    profile_raw := dbmapping.Profile{}
    err := c.Db.Get(&profile_raw, c.Db.Rebind("SELECT * FROM profiles WHERE player_id=? ORDER BY created_at DESC LIMIT 1"), player_id)
    if err != nil {
        log.Println(err)
        return profile_raw, false
    }

    return profile_raw, true
}
