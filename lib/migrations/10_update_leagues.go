// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"database/sql"
)

func UpdateLeaguesUp(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE `leagues` SET symbol='🈸' WHERE symbol=':u7533:';")
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE `leagues` SET symbol='🈳 ' WHERE symbol=':u6e80';")
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE `leagues` SET symbol='🈵' WHERE symbol=':u7a7a:';")
	if err != nil {
		return err
	}

	return nil
}

func UpdateLeaguesDown(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE `leagues` SET symbol=':u7533:' WHERE symbol='🈸';")
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE `leagues` SET symbol=':u6e80' WHERE symbol='🈳 ';")
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE `leagues` SET symbol=':u7a7a:' WHERE symbol='🈵';")
	if err != nil {
		return err
	}

	return nil
}
