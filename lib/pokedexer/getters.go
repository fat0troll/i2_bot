// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
)

// External functions

func (p *Pokedexer) getBestPokememes(playerID int) (map[int]*dbmapping.PokememeFull, bool) {
	pokememesArray := make(map[int]*dbmapping.PokememeFull)

	playerRaw, err := c.DataCache.GetPlayerByID(playerID)
	if err != nil {
		c.Log.Error(err.Error())
		return pokememesArray, false
	}
	profileRaw, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return pokememesArray, false
	}

	if playerRaw.LeagueID == 0 {
		return pokememesArray, false
	}

	allPokememes := c.DataCache.GetAllPokememes()
	if profileRaw.LevelID < 4 {
		for i := range allPokememes {
			if allPokememes[i].Pokememe.Grade == profileRaw.LevelID+1 {
				pokememesArray[allPokememes[i].Pokememe.Attack] = allPokememes[i]
			}
		}
	} else if profileRaw.LevelID > 8 {
		// TODO: Remove it on 10th grade pokememes arrival
		for i := range allPokememes {
			if allPokememes[i].Pokememe.Grade == 9 {
				matchLeague := false
				for j := range allPokememes[i].Elements {
					if allPokememes[i].Elements[j].LeagueID == playerRaw.LeagueID {
						matchLeague = true
					}
				}
				if matchLeague {
					pokememesArray[allPokememes[i].Pokememe.Attack] = allPokememes[i]
				}
			}
		}
	} else {
		for i := range allPokememes {
			if allPokememes[i].Pokememe.Grade == profileRaw.LevelID+1 {
				matchLeague := false
				for j := range allPokememes[i].Elements {
					if allPokememes[i].Elements[j].LeagueID == playerRaw.LeagueID {
						matchLeague = true
					}
				}
				if matchLeague {
					pokememesArray[allPokememes[i].Pokememe.Attack] = allPokememes[i]
				}
			}
		}
	}

	return pokememesArray, true
}
