// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

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
