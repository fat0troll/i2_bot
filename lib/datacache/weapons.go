// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/datamapping"
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

	weapons = append(weapons, datamapping.Weapon{1, "Бита", 2, 5})
	weapons = append(weapons, datamapping.Weapon{2, "Стальная бита", 10, 40})
	weapons = append(weapons, datamapping.Weapon{3, "Чугунная бита", 200, 500})
	weapons = append(weapons, datamapping.Weapon{4, "Титановая бита", 2000, 10000})
	weapons = append(weapons, datamapping.Weapon{5, "Алмазная бита", 10000, 100000})
	weapons = append(weapons, datamapping.Weapon{6, "Криптонитовая бита", 100000, 500000})
	weapons = append(weapons, datamapping.Weapon{7, "Буханка из пятёры", 1000000, 5000000})

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
