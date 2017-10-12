// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package getters

import (
    // stdlib
    "log"
    "time"
    // local
    "../dbmapping"
)

func (g *Getters) GetPlayerByID(player_id int) (dbmapping.Player, bool) {
    player_raw := dbmapping.Player{}
    err := c.Db.Get(&player_raw, c.Db.Rebind("SELECT * FROM players WHERE id=?"), player_id)
    if err != nil {
        log.Println(err)
        return player_raw, false
    }

    return player_raw, true
}

func (g *Getters) GetOrCreatePlayer(telegram_id int) (dbmapping.Player, bool) {
    player_raw := dbmapping.Player{}
    err := c.Db.Get(&player_raw, c.Db.Rebind("SELECT * FROM players WHERE telegram_id=?"), telegram_id)
    if err != nil {
        log.Printf("Message user not found in database.")
        log.Printf(err.Error())

        // Create "nobody" user
        player_raw.Telegram_id = telegram_id
        player_raw.League_id = 0
        player_raw.Squad_id = 0
        player_raw.Status = "nobody"
        player_raw.Created_at = time.Now().UTC()
        player_raw.Updated_at = time.Now().UTC()
        _, err = c.Db.NamedExec("INSERT INTO players VALUES(NULL, :telegram_id, :league_id, :squad_id, :status, :created_at, :updated_at)", &player_raw)
        if err != nil {
            log.Printf(err.Error())
            return player_raw, false
        }
    } else {
        log.Printf("Message user found in database.")
    }

    return player_raw, true
}
