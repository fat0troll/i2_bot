// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
)

// Internal functions

func (p *Pokedexer) formFullPokememes(pokememes []dbmapping.Pokememe) ([]dbmapping.PokememeFull, bool) {
	pokememesArray := []dbmapping.PokememeFull{}
	elements := []dbmapping.Element{}
	err := c.Db.Select(&elements, "SELECT * FROM elements")
	if err != nil {
		c.Log.Error(err)
		return pokememesArray, false
	}
	locations := []dbmapping.Location{}
	err = c.Db.Select(&locations, "SELECT * FROM locations")
	if err != nil {
		c.Log.Error(err)
		return pokememesArray, false
	}
	pokememesElements := []dbmapping.PokememeElement{}
	err = c.Db.Select(&pokememesElements, "SELECT * FROM pokememes_elements")
	if err != nil {
		c.Log.Error(err)
		return pokememesArray, false
	}
	pokememesLocations := []dbmapping.PokememeLocation{}
	err = c.Db.Select(&pokememesLocations, "SELECT * FROM pokememes_locations")
	if err != nil {
		c.Log.Error(err)
		return pokememesArray, false
	}

	for i := range pokememes {
		fullPokememe := dbmapping.PokememeFull{}
		elementsListed := []dbmapping.Element{}
		locationsListed := []dbmapping.Location{}

		for j := range pokememesLocations {
			if pokememesLocations[j].PokememeID == pokememes[i].ID {
				for l := range locations {
					if pokememesLocations[j].LocationID == locations[l].ID {
						locationsListed = append(locationsListed, locations[l])
					}
				}
			}
		}

		for k := range pokememesElements {
			if pokememesElements[k].PokememeID == pokememes[i].ID {
				for e := range elements {
					if pokememesElements[k].ElementID == elements[e].ID {
						elementsListed = append(elementsListed, elements[e])
					}
				}
			}
		}

		fullPokememe.Pokememe = pokememes[i]
		fullPokememe.Elements = elementsListed
		fullPokememe.Locations = locationsListed

		pokememesArray = append(pokememesArray, fullPokememe)
	}

	return pokememesArray, true
}

// External functions

// GetPokememes returns all existing pokememes, known by bot
func (p *Pokedexer) GetPokememes() ([]dbmapping.PokememeFull, bool) {
	pokememesArray := []dbmapping.PokememeFull{}
	pokememes := []dbmapping.Pokememe{}
	err := c.Db.Select(&pokememes, "SELECT * FROM pokememes ORDER BY grade asc, name asc")
	if err != nil {
		c.Log.Error(err)
		return pokememesArray, false
	}

	pokememesArray, ok := p.formFullPokememes(pokememes)
	return pokememesArray, ok
}

func (p *Pokedexer) getBestPokememes(playerID int) ([]dbmapping.PokememeFull, bool) {
	pokememesArray := []dbmapping.PokememeFull{}
	playerRaw, ok := c.Users.GetPlayerByID(playerID)
	if !ok {
		return pokememesArray, ok
	}
	profileRaw, ok := c.Users.GetProfile(playerID)
	if !ok {
		return pokememesArray, ok
	}

	if playerRaw.LeagueID == 0 {
		return pokememesArray, false
	}

	// TODO: make it more complicated
	pokememes := []dbmapping.Pokememe{}
	err := c.Db.Select(&pokememes, c.Db.Rebind("SELECT p.* FROM pokememes p, pokememes_elements pe, elements e WHERE e.league_id = ? AND p.grade = ? AND pe.element_id = e.id AND pe.pokememe_id = p.id ORDER BY p.attack DESC"), playerRaw.LeagueID, profileRaw.LevelID+1)
	if err != nil {
		c.Log.Error(err)
		return pokememesArray, false
	}

	pokememesArray, ok = p.formFullPokememes(pokememes)
	return pokememesArray, ok
}

// GetPokememeByID returns single pokememe based on internal ID in database
func (p *Pokedexer) GetPokememeByID(pokememeID string) (dbmapping.PokememeFull, bool) {
	fullPokememe := dbmapping.PokememeFull{}
	pokememe := dbmapping.Pokememe{}
	err := c.Db.Get(&pokememe, c.Db.Rebind("SELECT * FROM pokememes WHERE id=?"), pokememeID)
	if err != nil {
		c.Log.Error(err)
		return fullPokememe, false
	}
	elements := []dbmapping.Element{}
	err = c.Db.Select(&elements, "SELECT * FROM elements")
	if err != nil {
		c.Log.Error(err)
		return fullPokememe, false
	}
	locations := []dbmapping.Location{}
	err = c.Db.Select(&locations, "SELECT * FROM locations")
	if err != nil {
		c.Log.Error(err)
		return fullPokememe, false
	}
	pokememesElements := []dbmapping.PokememeElement{}
	err = c.Db.Select(&pokememesElements, "SELECT * FROM pokememes_elements WHERE pokememe_id='"+strconv.Itoa(pokememe.ID)+"'")
	if err != nil {
		c.Log.Error(err)
		return fullPokememe, false
	}
	pokememesLocations := []dbmapping.PokememeLocation{}
	err = c.Db.Select(&pokememesLocations, "SELECT * FROM pokememes_locations WHERE pokememe_id='"+strconv.Itoa(pokememe.ID)+"'")
	if err != nil {
		c.Log.Error(err)
		return fullPokememe, false
	}

	elementsListed := []dbmapping.Element{}
	locationsListed := []dbmapping.Location{}

	for j := range pokememesLocations {
		if pokememesLocations[j].PokememeID == pokememe.ID {
			for l := range locations {
				if pokememesLocations[j].LocationID == locations[l].ID {
					locationsListed = append(locationsListed, locations[l])
				}
			}
		}
	}

	for k := range pokememesElements {
		if pokememesElements[k].PokememeID == pokememe.ID {
			for e := range elements {
				if pokememesElements[k].ElementID == elements[e].ID {
					elementsListed = append(elementsListed, elements[e])
				}
			}
		}
	}

	fullPokememe.Pokememe = pokememe
	fullPokememe.Elements = elementsListed
	fullPokememe.Locations = locationsListed

	return fullPokememe, true
}
