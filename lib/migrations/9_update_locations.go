// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
    // stdlib
    "database/sql"
)

func UpdateLocationsUp(tx *sql.Tx) error {
    _, err := tx.Exec("UPDATE `locations` SET symbol='⛪' WHERE symbol=':church:';")
    if err != nil {
        return err
    }
    _, err = tx.Exec("UPDATE `locations` SET symbol='🌲' WHERE symbol=':evergreen_tree:';")
    if err != nil {
        return err
    }
    _, err = tx.Exec("UPDATE `locations` SET symbol='🚣' WHERE symbol=':rowboat:';")
    if err != nil {
        return err
    }

    return nil
}

func UpdateLocationsDown(tx *sql.Tx) error {
    _, err := tx.Exec("UPDATE `locations` SET symbol=':church:' WHERE symbol='⛪'';")
    if err != nil {
        return err
    }
    _, err = tx.Exec("UPDATE `locations` SET symbol=':evergreen_tree:' WHERE symbol='🌲';")
    if err != nil {
        return err
    }
    _, err = tx.Exec("UPDATE `locations` SET symbol=':rowboat:' WHERE symbol='🚣';")
    if err != nil {
        return err
    }

    return nil
}
