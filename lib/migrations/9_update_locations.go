// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package migrations

import (
	"database/sql"
)

// UpdateLocationsUp fixes some fuckup with locations' emoji
func UpdateLocationsUp(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE `locations` SET symbol='⛪' WHERE symbol=':church:'")
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE `locations` SET symbol='🌲' WHERE symbol=':evergreen_tree:'")
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE `locations` SET symbol='🚣' WHERE symbol=':rowboat:'")

	return err
}

// UpdateLocationsDown returns location emoji fuckup for sanity purposes
func UpdateLocationsDown(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE `locations` SET symbol=':church:' WHERE symbol='⛪''")
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE `locations` SET symbol=':evergreen_tree:' WHERE symbol='🌲'")
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE `locations` SET symbol=':rowboat:' WHERE symbol='🚣'")

	return err
}
