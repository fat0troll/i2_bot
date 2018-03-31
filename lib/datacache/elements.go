// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
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

	elements = append(elements, datamapping.Element{1, "👊", "Боевой", 1})
	elements = append(elements, datamapping.Element{2, "🌀", "Летающий", 1})
	elements = append(elements, datamapping.Element{3, "💀", "Ядовитый", 1})
	elements = append(elements, datamapping.Element{4, "🗿", "Каменный", 1})
	elements = append(elements, datamapping.Element{5, "🔥", "Огненный", 2})
	elements = append(elements, datamapping.Element{6, "⚡", "Электрический", 2})
	elements = append(elements, datamapping.Element{7, "💧", "Водяной", 2})
	elements = append(elements, datamapping.Element{8, "🍀", "Травяной", 2})
	elements = append(elements, datamapping.Element{9, "💩", "Отважный", 3})
	elements = append(elements, datamapping.Element{10, "👁", "Психический", 3})
	elements = append(elements, datamapping.Element{11, "👿", "Тёмный", 3})
	elements = append(elements, datamapping.Element{12, "⌛", "Времени", 3})

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
