// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"time"
)

// GetProfile returns last saved profile of player
func (u *Users) GetProfile(playerID int) (dbmapping.Profile, bool) {
	profileRaw := dbmapping.Profile{}
	err := c.Db.Get(&profileRaw, c.Db.Rebind("SELECT * FROM profiles WHERE player_id=? ORDER BY created_at DESC LIMIT 1"), playerID)
	if err != nil {
		c.Log.Error(err)
		return profileRaw, false
	}

	return profileRaw, true
}

// GetPlayerByID returns dbmapping.Player instance with given ID.
func (u *Users) GetPlayerByID(playerID int) (dbmapping.Player, bool) {
	playerRaw := dbmapping.Player{}
	err := c.Db.Get(&playerRaw, c.Db.Rebind("SELECT * FROM players WHERE id=?"), playerID)
	if err != nil {
		c.Log.Error(err.Error())
		return playerRaw, false
	}

	return playerRaw, true
}

// GetOrCreatePlayer seeks for player in database via Telegram ID.
// In case, when there is no player with such ID, new player will be created.
func (u *Users) GetOrCreatePlayer(telegramID int) (dbmapping.Player, bool) {
	playerRaw := dbmapping.Player{}
	err := c.Db.Get(&playerRaw, c.Db.Rebind("SELECT * FROM players WHERE telegram_id=?"), telegramID)
	if err != nil {
		c.Log.Error("Message user not found in database.")
		c.Log.Error(err.Error())

		// Create "nobody" user
		playerRaw.TelegramID = telegramID
		playerRaw.LeagueID = 0
		playerRaw.Status = "nobody"
		playerRaw.CreatedAt = time.Now().UTC()
		playerRaw.UpdatedAt = time.Now().UTC()
		_, err = c.Db.NamedExec("INSERT INTO players VALUES(NULL, :telegram_id, :league_id, :status, :created_at, :updated_at)", &playerRaw)
		if err != nil {
			c.Log.Error(err.Error())
			return playerRaw, false
		}
	} else {
		c.Log.Debug("Message user found in database.")
	}

	return playerRaw, true
}

// GetPrettyName returns "pretty" name of user (first_name + last name or username)
func (u *Users) GetPrettyName(user *tgbotapi.User) string {
	userName := user.FirstName
	if user.LastName != "" {
		userName += " " + user.LastName
	}

	if user.UserName != "" {
		userName += " (@" + user.UserName + ")"
	}
	return c.Users.FormatUsername(userName)
}

// PlayerBetterThan return true, if profile is more or equal powerful than
// provided power level
func (u *Users) PlayerBetterThan(playerRaw *dbmapping.Player, powerLevel string) bool {
	var isBetter = false
	switch playerRaw.Status {
	case "special":
		isBetter = true
	case "owner":
		isBetter = true
	case "admin":
		if powerLevel != "owner" {
			isBetter = true
		}
	case "academic":
		if powerLevel != "ownder" && powerLevel != "admin" {
			isBetter = true
		}
	default:
		isBetter = false
	}

	return isBetter
}
