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

func (dc *DataCache) initLevels() {
	c.Log.Info("Initializing Levels storage...")
	dc.levels = make(map[int]*datamapping.Level)
}

func (dc *DataCache) loadLevels() {
	c.Log.Info("Load current Levels data to DataCache...")
	levels := dc.getLevels()

	for i := range levels {
		dc.levels[levels[i].ID] = &levels[i]
	}
	c.Log.Info("Loaded levels in DataCache: " + strconv.Itoa(len(dc.levels)))
}

func (dc *DataCache) getLevels() []datamapping.Level {
	levels := []datamapping.Level{}

	yamlFile, err := static.ReadFile("levels.yml")
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't read levels data file")
	}

	err = yaml.Unmarshal(yamlFile, &levels)
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't parse levels data file")
	}

	return levels
}

// External functions

// GetLevelByID returns level data by ID
func (dc *DataCache) GetLevelByID(levelID int) (*datamapping.Level, error) {
	if dc.levels[levelID] != nil {
		return dc.levels[levelID], nil
	}

	return nil, errors.New("There is no level with ID = " + strconv.Itoa(levelID))
}
