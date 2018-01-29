// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
)

func (dc *DataCache) initLocations() {
	c.Log.Info("Initializing Locations storage...")
	dc.locations = make(map[int]*dbmapping.Location)
}

func (dc *DataCache) loadLocations() {
	c.Log.Info("Load current Locations data from database to DataCache...")
	locations := []dbmapping.Location{}
	err := c.Db.Select(&locations, "SELECT * FROM locations")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}

	dc.locationsMutex.Lock()
	for i := range locations {
		dc.locations[locations[i].ID] = &locations[i]
	}
	c.Log.Info("Loaded locations in DataCache: " + strconv.Itoa(len(dc.locations)))
	dc.locationsMutex.Unlock()
}

func (dc *DataCache) findLocationIDByName(name string) (int, error) {
	dc.locationsMutex.Lock()
	for i := range dc.locations {
		if dc.locations[i].Name == name {
			dc.locationsMutex.Unlock()
			return i, nil
		}
	}

	dc.locationsMutex.Unlock()
	return 0, errors.New("There is no location with name = " + name)
}
