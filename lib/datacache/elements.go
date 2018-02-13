// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"strconv"
)

func (dc *DataCache) initElements() {
	c.Log.Info("Initializing Elements storage...")
	dc.elements = make(map[int]*dbmapping.Element)
}

func (dc *DataCache) loadElements() {
	c.Log.Info("Load current Elements data from database to DataCache...")
	elements := []dbmapping.Element{}
	err := c.Db.Select(&elements, "SELECT * FROM elements")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}

	dc.elementsMutex.Lock()
	for i := range elements {
		dc.elements[elements[i].ID] = &elements[i]
	}
	c.Log.Info("Loaded elements in DataCache: " + strconv.Itoa(len(dc.elements)))
	dc.elementsMutex.Unlock()
}

func (dc *DataCache) findElementIDBySymbol(symbol string) (int, error) {
	dc.elementsMutex.Lock()
	for i := range dc.elements {
		if dc.elements[i].Symbol == symbol {
			dc.elementsMutex.Unlock()
			return i, nil
		}
	}

	dc.elementsMutex.Unlock()
	return 0, errors.New("There is no element with symbol = " + symbol)
}
