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

func (dc *DataCache) initElements() {
	c.Log.Info("Initializing Elements storage...")
	dc.elements = make(map[int]*datamapping.Element)
}

func (dc *DataCache) loadElements() {
	c.Log.Info("Load current Elements data to DataCache...")
	elements := dc.getElements()

	for i := range elements {
		dc.elements[elements[i].ID] = &elements[i]
	}
	c.Log.Info("Loaded elements in DataCache: " + strconv.Itoa(len(dc.elements)))
}

func (dc *DataCache) getElements() []datamapping.Element {
	elements := []datamapping.Element{}

	yamlFile, err := static.ReadFile("elements.yml")
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't read elements data file")
	}

	err = yaml.Unmarshal(yamlFile, &elements)
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't parse elements data file")
	}

	return elements
}

func (dc *DataCache) findElementIDBySymbol(symbol string) (int, error) {
	for i := range dc.elements {
		if dc.elements[i].Symbol == symbol {
			return i, nil
		}
	}

	return 0, errors.New("There is no element with symbol = " + symbol)
}

// External functions

// GetElementByID returns element with given ID
func (dc *DataCache) GetElementByID(elementID int) (*datamapping.Element, error) {
	if dc.elements[elementID] != nil {
		return dc.elements[elementID], nil
	}

	return nil, errors.New("There is no element with ID = " + strconv.Itoa(elementID))
}
