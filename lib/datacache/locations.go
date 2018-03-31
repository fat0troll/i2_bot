// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
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

	locations = append(locations, datamapping.Location{1, "🌲", "Лес"})
	locations = append(locations, datamapping.Location{2, "⛰", "Горы"})
	locations = append(locations, datamapping.Location{3, "🚣", "Озеро"})
	locations = append(locations, datamapping.Location{4, "🏙", "Город"})
	locations = append(locations, datamapping.Location{5, "🏛", "Катакомбы"})
	locations = append(locations, datamapping.Location{6, "⛪️", "Кладбище"})

	return locations
}

func (dc *DataCache) findLocationIDByName(name string) (int, error) {
	for i := range dc.locations {
		if dc.locations[i].Name == name {
			return i, nil
		}
	}

	return 0, errors.New("There is no location with name = " + name)
}
