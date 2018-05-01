// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"gopkg.in/yaml.v2"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
	"source.wtfteam.pro/i2_bot/i2_bot/static"
	"strconv"
)

func (dc *DataCache) initLocations() {
	c.Log.Info("Initializing Locations storage...")
	dc.locations = make(map[int]*datamapping.Location)
}

func (dc *DataCache) loadLocations() {
	c.Log.Info("Load current Locations data to DataCache...")
	locations := dc.getLocations()

	for i := range locations {
		dc.locations[locations[i].ID] = &locations[i]
	}
	c.Log.Info("Loaded locations in DataCache: " + strconv.Itoa(len(dc.locations)))
}

func (dc *DataCache) getLocations() []datamapping.Location {
	locations := []datamapping.Location{}

	yamlFile, err := static.ReadFile("locations.yml")
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't read locations data file")
	}

	err = yaml.Unmarshal(yamlFile, &locations)
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't parse locations data file")
	}

	return locations
}

func (dc *DataCache) getLocationByID(locationID int) (*datamapping.Location, error) {
	if dc.locations[locationID] != nil {
		return dc.locations[locationID], nil
	}

	return nil, errors.New("There is no localtion with ID = " + strconv.Itoa(locationID))
}

// External functions

// FindLocationIDByName returns location ID for given location name
func (dc *DataCache) FindLocationIDByName(name string) (int, error) {
	for i := range dc.locations {
		if dc.locations[i].Name == name {
			return i, nil
		}
	}

	return 0, errors.New("There is no location with name = " + name)
}
