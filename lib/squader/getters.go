// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squader

import (
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// GetSquadByID returns squad will all support information
func (s *Squader) GetSquadByID(squadID int) (dbmapping.SquadChat, bool) {
	squadFull := dbmapping.SquadChat{}
	squad := dbmapping.Squad{}
	chat := dbmapping.Chat{}
	floodChat := dbmapping.Chat{}

	err := c.Db.Get(&squad, c.Db.Rebind("SELECT * FROM squads WHERE id=?"), squadID)
	if err != nil {
		c.Log.Error(err)
		return squadFull, false
	}

	err = c.Db.Get(&chat, c.Db.Rebind("SELECT * FROM chats WHERE id=?"), squad.ChatID)
	if err != nil {
		c.Log.Error(err)
		return squadFull, false
	}
	err = c.Db.Get(&floodChat, c.Db.Rebind("SELECT * FROM chats WHERE id=?"), squad.FloodChatID)
	if err != nil {
		c.Log.Error(err)
		return squadFull, false
	}

	squadFull.Squad = squad
	squadFull.Chat = chat
	squadFull.FloodChat = floodChat

	return squadFull, true
}

// GetAvailableSquadChatsForUser returns squad chats which user can join
func (s *Squader) GetAvailableSquadChatsForUser(playerRaw *dbmapping.Player) ([]dbmapping.Chat, bool) {
	groupChats := []dbmapping.Chat{}

	err := c.Db.Select(&groupChats, c.Db.Rebind("SELECT ch.* FROM chats ch, squads s, squads_players sp WHERE (s.chat_id=ch.id OR s.flood_chat_id=ch.id) AND sp.player_id = ? AND s.id = sp.squad_id"), playerRaw.ID)
	if err != nil {
		c.Log.Error(err)
		return groupChats, false
	}

	return groupChats, true
}

// GetAllSquadChats returns all main squad chats
func (s *Squader) GetAllSquadChats() ([]dbmapping.Chat, bool) {
	groupChats := []dbmapping.Chat{}

	err := c.Db.Select(&groupChats, "SELECT ch.* FROM chats ch, squads s WHERE s.chat_id=ch.id")
	if err != nil {
		c.Log.Error(err)
		return groupChats, false
	}

	return groupChats, true
}

// GetAllSquadFloodChats returns all flood squad chats
func (s *Squader) GetAllSquadFloodChats() ([]dbmapping.Chat, bool) {
	groupChats := []dbmapping.Chat{}

	err := c.Db.Select(&groupChats, "SELECT ch.* FROM chats ch, squads s WHERE s.flood_chat_id=ch.id")
	if err != nil {
		c.Log.Error(err)
		return groupChats, false
	}

	return groupChats, true
}

// GetUserRolesInSquads lists all user roles
func (s *Squader) GetUserRolesInSquads(playerRaw *dbmapping.Player) ([]dbmapping.SquadPlayerFull, bool) {
	userRoles := []dbmapping.SquadPlayerFull{}
	userRolesRaw := []dbmapping.SquadPlayer{}

	err := c.Db.Select(&userRolesRaw, c.Db.Rebind("SELECT * FROM squads_players WHERE player_id=?"), playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return userRoles, false
	}

	for i := range userRolesRaw {
		userRoleFull := dbmapping.SquadPlayerFull{}
		userRoleFull.Player = *playerRaw
		userProfile, profileOk := c.Users.GetProfile(playerRaw.ID)
		userRoleFull.Profile = userProfile
		userRoleFull.UserRole = userRolesRaw[i].UserType
		squad, squadOk := s.GetSquadByID(userRolesRaw[i].SquadID)
		userRoleFull.Squad = squad

		if profileOk && squadOk {
			userRoles = append(userRoles, userRoleFull)
		}
	}

	return userRoles, true
}

// IsChatASquadEnabled checks group chat for restricting actions for squad
func (s *Squader) IsChatASquadEnabled(chatRaw *dbmapping.Chat) string {
	mainChats, ok := s.GetAllSquadChats()
	if !ok {
		return "no"
	}
	floodChats, ok := s.GetAllSquadFloodChats()
	if !ok {
		return "no"
	}

	for i := range mainChats {
		if *chatRaw == mainChats[i] {
			return "main"
		}
	}

	for i := range floodChats {
		if *chatRaw == floodChats[i] {
			return "flood"
		}
	}

	return "no"
}
