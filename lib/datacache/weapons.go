// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"gopkg.in/yaml.v2"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
	"source.wtfteam.pro/i2_bot/i2_bot/static"
	"strconv"
	"strings"
)

func (dc *DataCache) initWeapons() {
	c.Log.Info("Initializing Weapons storage...")
	dc.weapons = make(map[int]*datamapping.Weapon)
}

func (dc *DataCache) loadWeapons() {
	c.Log.Info("Load current Weapons data to DataCache...")
	weapons := dc.getWeapons()

	for i := range weapons {
		dc.weapons[weapons[i].ID] = &weapons[i]
	}
	c.Log.Info("Loaded weapon types in DataCache: " + strconv.Itoa(len(dc.weapons)))
}

func (dc *DataCache) getWeapons() []datamapping.Weapon {
	weapons := []datamapping.Weapon{}

	yamlFile, err := static.ReadFile("weapons.yml")
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't read weapons data file")
	}

	err = yaml.Unmarshal(yamlFile, &weapons)
	if err != nil {
		c.Log.Error(err.Error())
		c.Log.Fatal("Can't parse weapons data file")
	}

	return weapons
}

// External functions

// GetWeaponTypeByID returns weapon type from datacache by given ID
func (dc *DataCache) GetWeaponTypeByID(weaponID int) (*datamapping.Weapon, error) {
	if dc.weapons[weaponID] != nil {
		c.Log.Debug("DataCache: found weapon type with ID = " + strconv.Itoa(weaponID))
		return dc.weapons[weaponID], nil
	}
	return nil, errors.New("There is no weapon type with ID = " + strconv.Itoa(weaponID))
}

// GetWeaponTypeByName returns weapon type from datacache by weapon name
func (dc *DataCache) GetWeaponTypeByName(name string) (*datamapping.Weapon, error) {
	for i := range dc.weapons {
		if strings.HasPrefix(dc.weapons[i].Name, name) {
			return dc.weapons[i], nil
		}
	}

	return nil, errors.New("There is no weapon type with name = " + name)
}
