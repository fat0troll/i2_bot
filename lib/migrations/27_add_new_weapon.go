// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

// AddNewWeaponUp adds new weapon, presented on 26.11.2017
func AddNewWeaponUp(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO `weapons` VALUES(NULL, 'Буханка из пятёры', 1000000, 5000000, NOW())")
	if err != nil {
		return err
	}

	return nil
}

// AddNewWeaponDown does nothing
func AddNewWeaponDown(tx *sql.Tx) error {
	return nil
}
