// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"strconv"

	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

func (dc *DataCache) initSquads() {
	c.Log.Info("Initializing Squads storage...")
	dc.squads = make(map[int]*dbmapping.Squad)
	dc.squadsWithChats = make(map[int]*dbmapping.SquadChat)
	dc.squadPlayers = make(map[int]map[int]*dbmapping.SquadPlayerFull)
	dc.squadPlayersRelations = make(map[int]*dbmapping.SquadPlayer)
}

func (dc *DataCache) loadSquads() {
	c.Log.Info("Load current Squads data from database to DataCache...")
	squads := []dbmapping.Squad{}
	err := c.Db.Select(&squads, "SELECT * FROM squads")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}

	squadsPlayersRelations := []dbmapping.SquadPlayer{}
	err = c.Db.Select(&squadsPlayersRelations, "SELECT * FROM squads_players")
	if err != nil {
		c.Log.Fatal(err.Error())
	}

	dc.squadsMutex.Lock()
	for i := range squads {
		squadWithChat := dbmapping.SquadChat{}
		squadWithChat.Squad = squads[i]
		sChat := dc.chats[squads[i].ChatID]
		if sChat != nil {
			squadWithChat.Chat = *sChat

			dc.squads[squads[i].ID] = &squads[i]
			dc.squadsWithChats[squads[i].ID] = &squadWithChat
		}

		dc.squadPlayers[squads[i].ID] = make(map[int]*dbmapping.SquadPlayerFull)
	}

	for i := range squadsPlayersRelations {
		sPlayer := dc.players[squadsPlayersRelations[i].PlayerID]
		sProfile := dc.currentProfiles[squadsPlayersRelations[i].PlayerID]
		sSquad := dc.squadsWithChats[squadsPlayersRelations[i].SquadID]
		if sPlayer != nil && sProfile != nil && sSquad != nil {
			dc.squadPlayersRelations[squadsPlayersRelations[i].ID] = &squadsPlayersRelations[i]

			squadPlayer := dbmapping.SquadPlayerFull{}
			squadPlayer.Player = *sPlayer
			squadPlayer.Profile = *sProfile
			squadPlayer.Squad = *sSquad
			squadPlayer.UserRole = squadsPlayersRelations[i].UserType

			dc.squadPlayers[sSquad.Squad.ID][sPlayer.ID] = &squadPlayer
		} else {
			if sPlayer == nil {
				c.Log.Debug("Alert: player with ID=" + strconv.Itoa(squadsPlayersRelations[i].PlayerID) + "is nil")
			}
			if sProfile == nil {
				c.Log.Debug("Alert: player with ID=" + strconv.Itoa(squadsPlayersRelations[i].PlayerID) + "has no current profile")
			}
			if sSquad == nil {
				c.Log.Debug("Alert: squad with ID=" + strconv.Itoa(squadsPlayersRelations[i].SquadID) + "is nil")
			}
		}
	}
	c.Log.Info("Loaded squads in DataCache: " + strconv.Itoa(len(dc.squads)))
	c.Log.Info("Loaded players relations to squads in DataCache: " + strconv.Itoa(len(dc.squadPlayers)))
	dc.squadsMutex.Unlock()
}

// External functions

// AddPlayerToSquad creates relation between player and squad
func (dc *DataCache) AddPlayerToSquad(relation *dbmapping.SquadPlayer) (int, error) {
	sPlayer, err := c.DataCache.GetPlayerByID(relation.PlayerID)
	if err != nil {
		return 0, err
	}
	sProfile, err := c.DataCache.GetProfileByPlayerID(relation.PlayerID)
	if err != nil {
		return 0, err
	}
	sSquad, err := c.DataCache.GetSquadByID(relation.SquadID)
	if err != nil {
		return 0, err
	}
	dc.squadsMutex.Lock()
	for i := range dc.squadPlayersRelations {
		if dc.squadPlayersRelations[i].SquadID == relation.SquadID && dc.squadPlayersRelations[i].PlayerID == relation.PlayerID {
			dc.squadsMutex.Unlock()
			return 0, errors.New("There is already such a player-squad relation")
		}
	}
	dc.squadsMutex.Unlock()

	_, err = c.Db.NamedExec("INSERT INTO squads_players VALUES(NULL, :squad_id, :player_id, :user_type, :author_id, :created_at)", &relation)
	if err != nil {
		return 0, err
	}

	insertedRelation := dbmapping.SquadPlayer{}
	err = c.Db.Get(&insertedRelation, "SELECT * FROM squads_players WHERE squad_id=? AND player_id=?", relation.SquadID, relation.PlayerID)
	if err != nil {
		return 0, err
	}

	dc.squadsMutex.Lock()
	dc.squadPlayersRelations[insertedRelation.ID] = &insertedRelation
	squadPlayerFull := dbmapping.SquadPlayerFull{}
	squadPlayerFull.Player = *sPlayer
	squadPlayerFull.Profile = *sProfile
	squadPlayerFull.Squad = *sSquad
	squadPlayerFull.UserRole = insertedRelation.UserType
	if dc.squadPlayers[sSquad.Squad.ID] == nil {
		dc.squadPlayers[sSquad.Squad.ID] = make(map[int]*dbmapping.SquadPlayerFull)
	}
	dc.squadPlayers[sSquad.Squad.ID][sPlayer.ID] = &squadPlayerFull
	dc.squadsMutex.Unlock()

	return insertedRelation.ID, nil
}

// GetAllSquadMembers returns all squad members by squad ID
func (dc *DataCache) GetAllSquadMembers(squadID int) []dbmapping.SquadPlayerFull {
	players := []dbmapping.SquadPlayerFull{}
	dc.squadsMutex.Lock()
	for i := range dc.squadPlayers {
		if i == squadID {
			for j := range dc.squadPlayers[i] {
				players = append(players, *dc.squadPlayers[i][j])
			}
		}
	}
	dc.squadsMutex.Unlock()
	return players
}

// GetAllSquadsChats returns all chats belonging to squads
func (dc *DataCache) GetAllSquadsChats() []dbmapping.Chat {
	chats := []dbmapping.Chat{}

	dc.squadsMutex.Lock()
	for i := range dc.squadsWithChats {
		chats = append(chats, dc.squadsWithChats[i].Chat)
	}
	dc.squadsMutex.Unlock()

	return chats
}

// GetAllSquadsWithChats returns all squads with chats
func (dc *DataCache) GetAllSquadsWithChats() []dbmapping.SquadChat {
	squadsWithChats := []dbmapping.SquadChat{}

	dc.squadsMutex.Lock()
	for i := range dc.squadsWithChats {
		squadsWithChats = append(squadsWithChats, *dc.squadsWithChats[i])
	}
	dc.squadsMutex.Unlock()

	return squadsWithChats
}

// GetAvailableSquadsChatsForUser returns all squads chats accessible by user with given ID
func (dc *DataCache) GetAvailableSquadsChatsForUser(userID int) []dbmapping.Chat {
	chats := []dbmapping.Chat{}

	dc.squadsMutex.Lock()
	for i := range dc.squadPlayers {
		for j := range dc.squadPlayers[i] {
			if dc.squadPlayers[i][j].Player.ID == userID {
				chats = append(chats, dc.squadPlayers[i][j].Squad.Chat)
			}
		}
	}
	dc.squadsMutex.Unlock()

	return chats
}

// GetCommandersForSquad returns all players which are commanders of squad with given ID
func (dc *DataCache) GetCommandersForSquad(squadID int) []dbmapping.Player {
	commanders := []dbmapping.Player{}

	dc.squadsMutex.Lock()
	for i := range dc.squadPlayers[squadID] {
		if dc.squadPlayers[squadID][i].Squad.Squad.ID == squadID {
			if dc.squadPlayers[squadID][i].UserRole == "commander" {
				commanders = append(commanders, dc.squadPlayers[squadID][i].Player)
			}
		}
	}
	dc.squadsMutex.Unlock()

	return commanders
}

// GetSquadByID returns squad with given ID
func (dc *DataCache) GetSquadByID(squadID int) (*dbmapping.SquadChat, error) {
	if dc.squadsWithChats[squadID] != nil {
		return dc.squadsWithChats[squadID], nil
	}

	return nil, errors.New("There is no squad with ID=" + strconv.Itoa(squadID))
}

// GetSquadByChatID returns squad with given chat ID
func (dc *DataCache) GetSquadByChatID(chatID int) (*dbmapping.Squad, error) {
	dc.squadsMutex.Lock()
	for i := range dc.squadsWithChats {
		if dc.squadsWithChats[i].Chat.ID == chatID {
			dc.squadsMutex.Unlock()
			return dc.squads[i], nil
		}
	}
	dc.squadsMutex.Unlock()
	return nil, errors.New("There is no squad with chat ID=" + strconv.Itoa(chatID))
}

// GetSquadsChatsBySquadsIDs returns chats for given squad IDs
func (dc *DataCache) GetSquadsChatsBySquadsIDs(squadsIDs []int) []dbmapping.Chat {
	chats := []dbmapping.Chat{}

	dc.squadsMutex.Lock()
	for i := range dc.squadsWithChats {
		for j := range squadsIDs {
			if dc.squadsWithChats[i].Squad.ID == j {
				chats = append(chats, dc.squadsWithChats[i].Chat)
			}
		}
	}
	dc.squadsMutex.Unlock()

	return chats
}

// GetUserRolesInSquads returns all user roles for given user ID
func (dc *DataCache) GetUserRolesInSquads(userID int) []dbmapping.SquadPlayerFull {
	userRoles := []dbmapping.SquadPlayerFull{}

	dc.squadsMutex.Lock()
	for i := range dc.squadPlayers {
		for j := range dc.squadPlayers[i] {
			if dc.squadPlayers[i][j].Player.ID == userID {
				userRoles = append(userRoles, *dc.squadPlayers[i][j])
			}
		}
	}
	dc.squadsMutex.Unlock()

	return userRoles
}

// GetUserRoleInSquad returns user role in specified squad
func (dc *DataCache) GetUserRoleInSquad(squadID int, playerID int) string {
	if dc.squadPlayers[squadID][playerID] != nil {
		return dc.squadPlayers[squadID][playerID].UserRole
	}

	return "none"
}
