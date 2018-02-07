// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
)

func (dc *DataCache) initWeapons() {
	c.Log.Info("Initializing Weapons storage...")
	dc.weapons = make(map[int]*dbmapping.Weapon)
}

func (dc *DataCache) loadWeapons() {
	c.Log.Info("Load current Weapons data from database to DataCache...")
	weapons := []dbmapping.Weapon{}
	err := c.Db.Select(&weapons, "SELECT * FROM weapons")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}

	dc.weaponsMutex.Lock()
	for i := range weapons {
		dc.weapons[weapons[i].ID] = &weapons[i]
	}
	c.Log.Info("Loaded weapon types in DataCache: " + strconv.Itoa(len(dc.weapons)))
	dc.weaponsMutex.Unlock()
}

// External functions

// GetWeaponTypeByID returns weapon type from datacache by given ID
func (dc *DataCache) GetWeaponTypeByID(weaponID int) (*dbmapping.Weapon, error) {
	dc.weaponsMutex.Lock()
	if dc.weapons[weaponID] != nil {
		c.Log.Debug("DataCache: found weapon type with ID = " + strconv.Itoa(weaponID))
		dc.weaponsMutex.Unlock()
		return dc.weapons[weaponID], nil
	}
	dc.weaponsMutex.Unlock()
	return nil, errors.New("There is no weapon type with ID = " + strconv.Itoa(weaponID))
}

// GetWeaponTypeByName returns weapon type from datacache by weapon name
func (dc *DataCache) GetWeaponTypeByName(name string) (*dbmapping.Weapon, error) {
	dc.weaponsMutex.Lock()
	for i := range dc.weapons {
		if strings.HasPrefix(dc.weapons[i].Name, name) {
			dc.weaponsMutex.Unlock()
			return dc.weapons[i], nil
		}
	}

	dc.weaponsMutex.Unlock()
	return nil, errors.New("There is no weapon type with name = " + name)
}
