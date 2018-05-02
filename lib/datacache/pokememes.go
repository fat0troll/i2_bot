// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
	"source.wtfteam.pro/i2_bot/i2_bot/static"
)

func (dc *DataCache) initPokememes() {
	c.Log.Info("Initializing Pokememes storage...")
	dc.pokememes = make(map[int]*datamapping.Pokememe)
	dc.fullPokememes = make(map[int]*datamapping.PokememeFull)
	dc.pokememesGradeLocation = make(map[int]map[int]int)
}

func (dc *DataCache) loadPokememes() {
	c.Log.Info("Load current Pokememes data to DataCache...")

	pokememes := dc.getPokememes()

	for i := range pokememes {
		dc.pokememes[pokememes[i].ID] = &pokememes[i]

		if dc.pokememesGradeLocation[pokememes[i].Grade] == nil {
			dc.pokememesGradeLocation[pokememes[i].Grade] = make(map[int]int)
		}

		pokememeFull := datamapping.PokememeFull{}
		pokememeFullElements := []datamapping.Element{}
		pokememeFullLocations := []datamapping.Location{}
		pokememeFull.Pokememe = pokememes[i]
		for ii := range pokememes[i].Elements {
			element, err := dc.GetElementByID(pokememes[i].Elements[ii])
			if err != nil {
				// This is critical
				c.Log.Fatal(err.Error())
			}
			pokememeFullElements = append(pokememeFullElements, *element)
		}
		for ii := range pokememes[i].Locations {
			location, err := dc.getLocationByID(pokememes[i].Locations[ii])
			if err != nil {
				// This is critical
				c.Log.Fatal(err.Error())
			}
			pokememeFullLocations = append(pokememeFullLocations, *location)
			dc.pokememesGradeLocation[pokememes[i].Grade][location.ID]++
		}

		pokememeFull.Elements = pokememeFullElements
		pokememeFull.Locations = pokememeFullLocations
		dc.fullPokememes[pokememes[i].ID] = &pokememeFull
	}

	c.Log.Info("Loaded pokememes in DataCache: " + strconv.Itoa(len(dc.fullPokememes)))
}

func (dc *DataCache) getPokememes() []datamapping.Pokememe {
	pokememes := []datamapping.Pokememe{}

	allPokememesFiles, err := static.WalkDirs("pokememes", false)
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't read directory with pokememes information")
	}

	var pokememesData []byte

	for i := range allPokememesFiles {
		yamlFile, err := static.ReadFile(allPokememesFiles[i])
		if err != nil {
			c.Log.Error(err.Error())
			c.Log.Fatal("Can't read pokememes data file: " + allPokememesFiles[i])
		}

		for ii := range yamlFile {
			pokememesData = append(pokememesData, yamlFile[ii])
		}
		pokememesData = append(pokememesData, '\n')
	}

	err = yaml.Unmarshal(pokememesData, &pokememes)
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't parse merged pokememes data")
	}

	return pokememes
}

// External functions

// GetAllPokememes returns all pokememes
func (dc *DataCache) GetAllPokememes() map[int]*datamapping.PokememeFull {
	pokememes := make(map[int]*datamapping.PokememeFull)

	for i := range dc.fullPokememes {
		pokememes[dc.fullPokememes[i].Pokememe.ID] = dc.fullPokememes[i]
	}

	return pokememes
}

// GetPokememeByID returns pokememe with additional information by ID
func (dc *DataCache) GetPokememeByID(pokememeID int) (*datamapping.PokememeFull, error) {
	if dc.fullPokememes[pokememeID] != nil {
		return dc.fullPokememes[pokememeID], nil
	}

	return nil, errors.New("There is no pokememe with ID = " + strconv.Itoa(pokememeID))
}

// GetPokememeByName returns pokememe with additional information by name
func (dc *DataCache) GetPokememeByName(pokememeName string) (*datamapping.PokememeFull, error) {
	for i := range dc.fullPokememes {
		if strings.Contains(dc.fullPokememes[i].Pokememe.Name, pokememeName) {
			return dc.fullPokememes[i], nil
		}
	}

	return nil, errors.New("There is no pokememe with name = " + pokememeName)
}

// GetPokememesCountByGradeAndLocation returns pokememes count with given grade on given location
func (dc *DataCache) GetPokememesCountByGradeAndLocation(grade int, locationID int) int {
	if dc.pokememesGradeLocation[grade] == nil {
		return 0
	}

	return dc.pokememesGradeLocation[grade][locationID]
}
