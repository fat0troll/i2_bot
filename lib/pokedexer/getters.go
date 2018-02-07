// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"sort"
)

// External functions

func (p *Pokedexer) getBestPokememes(playerID int) ([]*dbmapping.PokememeFull, bool) {
	pokememesArray := make([]*dbmapping.PokememeFull, 0)

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

	pokememesArraySorted := make([]*dbmapping.PokememeFull, 0)

	for i := range allPokememes {
		pokememesArraySorted = append(pokememesArraySorted, allPokememes[i])
	}

	sort.Slice(pokememesArraySorted, func(i, j int) bool {
		return pokememesArraySorted[i].Pokememe.Attack > pokememesArraySorted[j].Pokememe.Attack
	})

	if profileRaw.LevelID < 4 {
		for i := range pokememesArraySorted {
			if (pokememesArraySorted[i].Pokememe.Defence < profileRaw.Power) || (pokememesArraySorted[i].Pokememe.Purchaseable) {
				if allPokememes[i].Pokememe.Grade == profileRaw.LevelID+1 {
					pokememesArray = append(pokememesArray, pokememesArraySorted[i])
				}
			}
		}
	} else if profileRaw.LevelID > 8 {
		// TODO: Remove it on 10th grade pokememes arrival
		for i := range allPokememes {
			if pokememesArraySorted[i].Pokememe.Grade == 9 {
				matchLeague := false
				for j := range pokememesArraySorted[i].Elements {
					if pokememesArraySorted[i].Elements[j].LeagueID == playerRaw.LeagueID {
						matchLeague = true
					}
				}
				if matchLeague {
					if (pokememesArraySorted[i].Pokememe.Defence < profileRaw.Power) || (pokememesArraySorted[i].Pokememe.Purchaseable) {
						pokememesArray = append(pokememesArray, pokememesArraySorted[i])
					}
				}
			}
		}
	} else {
		for i := range allPokememes {
			if pokememesArraySorted[i].Pokememe.Grade == profileRaw.LevelID+1 {
				matchLeague := false
				for j := range pokememesArraySorted[i].Elements {
					if pokememesArraySorted[i].Elements[j].LeagueID == playerRaw.LeagueID {
						matchLeague = true
					}
				}
				if matchLeague {
					if (pokememesArraySorted[i].Pokememe.Defence < profileRaw.Power) || (pokememesArraySorted[i].Pokememe.Purchaseable) {
						pokememesArray = append(pokememesArray, pokememesArraySorted[i])
					}
				}
			}
		}
	}

	return pokememesArray, true
}
