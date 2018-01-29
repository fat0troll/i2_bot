// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package statistics

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
)

// SquadStatictics generates statistics message snippet. Public due to usage in chats list
func (s *Statistics) SquadStatictics(squadID int) string {
	squadMembersWithInformation := []dbmapping.SquadPlayerFull{}
	squadMembers := []dbmapping.SquadPlayer{}

	squad, ok := c.Squader.GetSquadByID(squadID)
	if !ok {
		return "Невозможно получить информацию о данном отряде. Возможно, он пуст или произошла ошибка."
	}

	err := c.Db.Select(&squadMembers, c.Db.Rebind("SELECT * FROM squads_players WHERE squad_id=?"), squadID)
	if err != nil {
		c.Log.Error(err.Error())
		return "Невозможно получить информацию о данном отряде. Возможно, он пуст или произошла ошибка."
	}

	for i := range squadMembers {
		fullInfo := dbmapping.SquadPlayerFull{}

		playerRaw, err := c.DataCache.GetPlayerByID(squadMembers[i].PlayerID)
		if err != nil {
			c.Log.Error(err.Error())
		}
		profileRaw, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
		if err != nil {
			c.Log.Error(err.Error())
		}

		fullInfo.Squad = squad
		fullInfo.Player = *playerRaw
		fullInfo.Profile = *profileRaw

		squadMembersWithInformation = append(squadMembersWithInformation, fullInfo)
	}

	message := "Количество человек в отряде: " + strconv.Itoa(len(squadMembersWithInformation)) + "\n"

	summAttack := 0
	for i := range squadMembersWithInformation {
		summAttack += squadMembersWithInformation[i].Profile.Power
	}
	message += "Суммарная атака: " + strconv.Itoa(summAttack) + " очков.\n"

	return message
}
