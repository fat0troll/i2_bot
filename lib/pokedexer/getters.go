// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"sort"
	"strconv"
	"strings"

	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
)

func (p *Pokedexer) getAdvicePokememes(playerID int, adviceType string) ([]*datamapping.PokememeFull, bool) {
	c.Log.Debug("Getting advice for pokememes...")
	pokememesArray := make([]*datamapping.PokememeFull, 0)

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

	weapon, err := c.DataCache.GetWeaponTypeByID(profileRaw.WeaponID)
	if err != nil {
		c.Log.Debug(err.Error())
	}

	summPower := profileRaw.Power
	if weapon != nil {
		summPower = summPower + weapon.Power
	}

	allPokememes := c.DataCache.GetAllPokememes()
	neededGrade := 0
	if profileRaw.LevelID < 19 {
		neededGrade = profileRaw.LevelID + 1
	} else {
		neededGrade = 19
	}

	c.Log.Debug("This player will search for grade: " + strconv.Itoa(neededGrade))

	for i := range allPokememes {
		if allPokememes[i].Pokememe.Grade == neededGrade {
			matchLeague := false
			if profileRaw.LevelID < 4 {
				matchLeague = true
			} else if adviceType == "best_nofilter" || adviceType == "advice_all" {
				matchLeague = true
			} else {
				for j := range allPokememes[i].Elements {
					if allPokememes[i].Elements[j].LeagueID == playerRaw.LeagueID {
						matchLeague = true
					}
				}
			}
			if matchLeague {
				switch adviceType {
				case "best", "advice":
					if (allPokememes[i].Pokememe.Defence < summPower) || allPokememes[i].Pokememe.Purchaseable {
						pokememesArray = append(pokememesArray, allPokememes[i])
					}
				default:
					pokememesArray = append(pokememesArray, allPokememes[i])
				}
			}
		}
	}

	c.Log.Debug(strconv.Itoa(len(pokememesArray)) + " pokememes passed initial /best filtration.")

	// As we have already filtered this array, we need to sort it and pass to view
	sort.Slice(pokememesArray, func(i, j int) bool {
		if strings.HasPrefix(adviceType, "best") {
			return pokememesArray[i].Pokememe.Attack > pokememesArray[j].Pokememe.Attack
		}
		return pokememesArray[i].Pokememe.Price > pokememesArray[j].Pokememe.Price
	})

	switch adviceType {
	case "best", "advice", "best_nofilter":
		if len(pokememesArray) > 5 {
			pokememesArrayShorted := make([]*datamapping.PokememeFull, 0)
			for i := 0; i < 5; i++ {
				pokememesArrayShorted = append(pokememesArrayShorted, pokememesArray[i])
			}

			pokememesArray = pokememesArrayShorted
		}
	}

	return pokememesArray, true
}
