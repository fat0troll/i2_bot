// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"sort"
	"strconv"
)

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

	for i := range allPokememes {
		// Adding only affordable pokememes...
		if (allPokememes[i].Pokememe.Defence < profileRaw.Power) || allPokememes[i].Pokememe.Purchaseable {
			// ...and only of needed grade (+1 until 9)
			neededGrade := 0
			if profileRaw.LevelID < 9 {
				neededGrade = profileRaw.LevelID + 1
			} else {
				neededGrade = 9
			}
			if allPokememes[i].Pokememe.Grade == neededGrade {
				// ...and only of our elements if our level past 4
				matchLeague := false
				if profileRaw.LevelID < 4 {
					matchLeague = true
				} else {
					for j := range allPokememes[i].Elements {
						if allPokememes[i].Elements[j].LeagueID == playerRaw.LeagueID {
							matchLeague = true
						}
					}
				}
				if matchLeague {
					pokememesArray = append(pokememesArray, allPokememes[i])
				}
			}
		}
	}

	c.Log.Debug(strconv.Itoa(len(pokememesArray)) + " pokememes passed initial /best filtration.")

	// As we have already filtered this array, we need to sort it and pass to view
	sort.Slice(pokememesArray, func(i, j int) bool {
		return pokememesArray[i].Pokememe.Attack > pokememesArray[j].Pokememe.Attack
	})

	if len(pokememesArray) > 5 {
		idx := 0

		pokememesArrayShorted := make([]*dbmapping.PokememeFull, 0)

		for i := range pokememesArray {
			if idx < 5 {
				pokememesArrayShorted = append(pokememesArrayShorted, pokememesArray[i])
			}
			idx++
		}

		pokememesArray = pokememesArrayShorted
	}

	return pokememesArray, true
}

func (p *Pokedexer) getHighPricedPokememes(playerID int) ([]*dbmapping.PokememeFull, bool) {
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

	for i := range allPokememes {
		// Adding only affordable pokememes...
		// Only by force: who need to buy pokememe for selling?
		if allPokememes[i].Pokememe.Defence < profileRaw.Power {
			// ...and only of needed grade (+1 until 9)
			neededGrade := 0
			if profileRaw.LevelID < 9 {
				neededGrade = profileRaw.LevelID + 1
			} else {
				neededGrade = 9
			}
			if allPokememes[i].Pokememe.Grade == neededGrade {
				pokememesArray = append(pokememesArray, allPokememes[i])
			}
		}
	}

	c.Log.Debug(strconv.Itoa(len(pokememesArray)) + " pokememes passed initial /advice filtration.")

	// As we have already filtered this array, we need to sort it and pass to view
	sort.Slice(pokememesArray, func(i, j int) bool {
		return pokememesArray[i].Pokememe.Price > pokememesArray[j].Pokememe.Price
	})

	if len(pokememesArray) > 10 {
		idx := 0

		pokememesArrayShorted := make([]*dbmapping.PokememeFull, 0)

		for i := range pokememesArray {
			if idx < 10 {
				pokememesArrayShorted = append(pokememesArrayShorted, pokememesArray[i])
			}
			idx++
		}

		pokememesArray = pokememesArrayShorted
	}

	return pokememesArray, true
}
