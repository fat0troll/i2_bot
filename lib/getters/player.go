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

// GetPlayerByID returns dbmapping.Player instance with given ID.
func (g *Getters) GetPlayerByID(playerID int) (dbmapping.Player, bool) {
	playerRaw := dbmapping.Player{}
	err := c.Db.Get(&playerRaw, c.Db.Rebind("SELECT * FROM players WHERE id=?"), playerID)
	if err != nil {
		log.Println(err)
		return playerRaw, false
	}

	return playerRaw, true
}

// GetOrCreatePlayer seeks for player in database via Telegram ID.
// In case, when there is no player with such ID, new player will be created.
func (g *Getters) GetOrCreatePlayer(telegramID int) (dbmapping.Player, bool) {
	playerRaw := dbmapping.Player{}
	err := c.Db.Get(&playerRaw, c.Db.Rebind("SELECT * FROM players WHERE telegram_id=?"), telegramID)
	if err != nil {
		log.Printf("Message user not found in database.")
		log.Printf(err.Error())

		// Create "nobody" user
		playerRaw.TelegramID = telegramID
		playerRaw.LeagueID = 0
		playerRaw.SquadID = 0
		playerRaw.Status = "nobody"
		playerRaw.CreatedAt = time.Now().UTC()
		playerRaw.UpdatedAt = time.Now().UTC()
		_, err = c.Db.NamedExec("INSERT INTO players VALUES(NULL, :telegram_id, :league_id, :squad_id, :status, :created_at, :updated_at)", &playerRaw)
		if err != nil {
			log.Printf(err.Error())
			return playerRaw, false
		}
	} else {
		log.Printf("Message user found in database.")
	}

	return playerRaw, true
}

// PlayerBetterThan return true, if profile is more or equal powerful than
// provided power level
func (g *Getters) PlayerBetterThan(playerRaw *dbmapping.Player, powerLevel string) bool {
	var isBetter = false
	switch playerRaw.Status {
	case "owner":
		isBetter = true
	case "admin":
		if powerLevel != "owner" {
			isBetter = true
		}
	default:
		isBetter = false
	}

	return isBetter
}
