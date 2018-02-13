// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"strconv"
	"time"
)

func (dc *DataCache) initPlayers() {
	c.Log.Info("Initializing Players storage...")
	dc.players = make(map[int]*dbmapping.Player)
}

func (dc *DataCache) loadPlayers() {
	c.Log.Info("Load current Players data from database to DataCache...")
	players := []dbmapping.Player{}
	err := c.Db.Select(&players, "SELECT * FROM players")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}

	dc.playersMutex.Lock()
	for i := range players {
		dc.players[players[i].ID] = &players[i]
	}
	c.Log.Info("Loaded players in DataCache: " + strconv.Itoa(len(dc.players)))
	dc.playersMutex.Unlock()
}

// External functions

// AddPlayer creates new player in database
func (dc *DataCache) AddPlayer(player *dbmapping.Player) (int, error) {
	c.Log.Info("DataCache: Creating new user...")
	_, err := c.Db.NamedExec("INSERT INTO players VALUES(NULL, :telegram_id, :league_id, :status, :created_at, :updated_at)", &player)
	if err != nil {
		c.Log.Error(err.Error())
		return 0, err
	}

	insertedPlayer := dbmapping.Player{}
	err = c.Db.Get(&insertedPlayer, "SELECT * FROM players WHERE telegram_id=? AND created_at=?", player.TelegramID, player.CreatedAt)
	if err != nil {
		c.Log.Error(err.Error())
		return 0, err
	}

	dc.playersMutex.Lock()
	dc.players[insertedPlayer.ID] = &insertedPlayer
	dc.playersMutex.Unlock()

	return insertedPlayer.ID, nil
}

// GetOrCreatePlayerByTelegramID finds user by Telegram ID and creates one if not exist
func (dc *DataCache) GetOrCreatePlayerByTelegramID(telegramID int) (*dbmapping.Player, error) {
	c.Log.Info("DataCache: Getting player with Telegram ID=", telegramID)

	var player *dbmapping.Player
	dc.playersMutex.Lock()
	for i := range dc.players {
		if dc.players[i].TelegramID == telegramID {
			player = dc.players[i]
		}
	}
	dc.playersMutex.Unlock()

	if player == nil {
		c.Log.Info("There is no such user, creating one...")
		newPlayer := dbmapping.Player{}
		newPlayer.TelegramID = telegramID
		newPlayer.LeagueID = 0
		newPlayer.Status = "nobody"
		newPlayer.CreatedAt = time.Now().UTC()
		newPlayer.UpdatedAt = time.Now().UTC()

		newPlayerID, err := dc.AddPlayer(&newPlayer)
		if err != nil {
			return nil, err
		}

		player = dc.players[newPlayerID]
	}

	return player, nil
}

// GetPlayerByID returns player from datacache by ID
func (dc *DataCache) GetPlayerByID(playerID int) (*dbmapping.Player, error) {
	c.Log.Info("DataCache: Getting player with ID = ", playerID)
	if dc.players[playerID] != nil {
		c.Log.Debug("DataCache: found player with ID = " + strconv.Itoa(playerID))
		return dc.players[playerID], nil
	}

	return nil, errors.New("There is no user with ID = " + strconv.Itoa(playerID))
}

// GetPlayerByTelegramID returns player with such Telegram ID
func (dc *DataCache) GetPlayerByTelegramID(telegramID int) (*dbmapping.Player, error) {
	c.Log.Info("DataCache: Getting player with Telegram ID=", telegramID)

	var player *dbmapping.Player
	dc.playersMutex.Lock()
	for i := range dc.players {
		if dc.players[i].TelegramID == telegramID {
			player = dc.players[i]
		}
	}
	dc.playersMutex.Unlock()

	if player != nil {
		return player, nil
	}

	return nil, errors.New("There is no user with Telegram ID = " + strconv.Itoa(telegramID))
}

// UpdatePlayerFields writes new fields to player
func (dc *DataCache) UpdatePlayerFields(player *dbmapping.Player) (*dbmapping.Player, error) {
	if dc.players[player.ID] != nil {
		_, err := c.Db.NamedExec("UPDATE `players` SET league_id=:league_id, status=:status WHERE id=:id", player)
		if err != nil {
			c.Log.Error(err.Error())
			return dc.players[player.ID], err
		}
		dc.playersMutex.Lock()
		dc.players[player.ID].LeagueID = player.LeagueID
		dc.players[player.ID].Status = player.Status
		dc.playersMutex.Unlock()
		return dc.players[player.ID], nil
	}

	return nil, errors.New("There is no user with ID = " + strconv.Itoa(player.ID))
}

// UpdatePlayerTimestamp writes current update time to player
func (dc *DataCache) UpdatePlayerTimestamp(playerID int) error {
	if dc.players[playerID] != nil {
		dc.playersMutex.Lock()
		dc.players[playerID].UpdatedAt = time.Now().UTC()
		_, err := c.Db.NamedExec("UPDATE `players` SET updated_at=:updated_at WHERE id=:id", dc.players[playerID])
		if err != nil {
			c.Log.Error(err.Error())
			dc.playersMutex.Unlock()
			return err
		}
		dc.playersMutex.Unlock()
		return nil
	}

	return errors.New("There is no user with ID = " + strconv.Itoa(playerID))
}
