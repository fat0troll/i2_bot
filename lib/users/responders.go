// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
)

// FormatUsername formats Telegram username for posting
func (u *Users) FormatUsername(userName string) string {
	return strings.Replace(userName, "_", `\_`, -1)
}

// FindByName finds user with such username or nickname
func (u *Users) FindByName(update *tgbotapi.Update) string {
	commandArgs := update.Message.CommandArguments()
	if commandArgs == "" {
		c.Talkers.BotError(update)
		return "fail"
	}

	usersArray, ok := u.findUserByName(commandArgs)
	if !ok {
		return "fail"
	}

	message := "*Найденные игроки:*\n"

	for i := range usersArray {
		message += "#" + strconv.Itoa(usersArray[i].Player.ID)
		message += " " + usersArray[i].League.Symbol
		message += " " + usersArray[i].Profile.Nickname
		if usersArray[i].Profile.TelegramNickname != "" {
			message += " (@" + u.FormatUsername(usersArray[i].Profile.TelegramNickname) + ")"
		}
		message += " /profile" + strconv.Itoa(usersArray[i].Player.ID) + "\n"
		message += "Telegram ID: " + strconv.Itoa(usersArray[i].Player.TelegramID) + "\n"
		message += "Последнее обновление: " + usersArray[i].Profile.CreatedAt.Format("02.01.2006 15:04:05") + "\n"
	}

	c.Log.Debug(message)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

// ForeignProfileMessage shows profile of another user
func (u *Users) ForeignProfileMessage(update *tgbotapi.Update) string {
	userNum := strings.TrimPrefix(update.Message.Command(), "profile")
	userID, err := strconv.Atoi(userNum)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	playerRaw, ok := u.GetPlayerByID(userID)
	if !ok {
		return "fail"
	}

	_, ok = u.GetProfile(playerRaw.ID)
	if !ok {
		return c.Talkers.BotError(update)
	}

	return u.ProfileMessage(update, &playerRaw)
}

// ProfileMessage shows current player's profile
func (u *Users) ProfileMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	profileRaw, ok := u.GetProfile(playerRaw.ID)
	if !ok {
		return c.Talkers.AnyMessageUnauthorized(update)
	}
	league := dbmapping.League{}
	err := c.Db.Get(&league, c.Db.Rebind("SELECT * FROM leagues WHERE id=?"), playerRaw.LeagueID)
	if err != nil {
		c.Log.Error(err)
	}
	level := dbmapping.Level{}
	err = c.Db.Get(&level, c.Db.Rebind("SELECT * FROM levels WHERE id=?"), profileRaw.LevelID)
	if err != nil {
		c.Log.Error(err)
	}
	weapon := dbmapping.Weapon{}
	if profileRaw.WeaponID != 0 {
		err = c.Db.Get(&weapon, c.Db.Rebind("SELECT * FROM weapons WHERE id=?"), profileRaw.WeaponID)
		if err != nil {
			c.Log.Error(err)
		}
	}
	profilePokememes := []dbmapping.ProfilePokememe{}
	err = c.Db.Select(&profilePokememes, c.Db.Rebind("SELECT * FROM profiles_pokememes WHERE profile_id=?"), profileRaw.ID)
	if err != nil {
		c.Log.Error(err)
	}
	pokememes := []dbmapping.Pokememe{}
	err = c.Db.Select(&pokememes, c.Db.Rebind("SELECT * FROM pokememes"))
	if err != nil {
		c.Log.Error(err)
	}

	attackPokememes := 0
	for i := range profilePokememes {
		for j := range pokememes {
			if profilePokememes[i].PokememeID == pokememes[j].ID {
				singleAttack := profilePokememes[i].PokememeAttack
				attackPokememes += singleAttack
			}
		}
	}

	message := "*Профиль игрока "
	message += profileRaw.Nickname + "*"
	if profileRaw.TelegramNickname != "" {
		message += " (@" + u.FormatUsername(profileRaw.TelegramNickname) + ")"
	}
	message += "\nЛига: " + league.Symbol + league.Name
	message += "\n👤 " + strconv.Itoa(profileRaw.LevelID)
	message += " | 🎓 " + strconv.Itoa(profileRaw.Exp) + "/" + strconv.Itoa(level.MaxExp)
	message += " | 🥚 " + strconv.Itoa(profileRaw.EggExp) + "/" + strconv.Itoa(level.MaxEgg)
	message += "\n💲" + c.Statistics.GetPrintablePoints(profileRaw.Wealth)
	message += " |💎" + strconv.Itoa(profileRaw.Crystalls)
	message += " |⭕" + strconv.Itoa(profileRaw.Pokeballs)
	message += "\n⚔Атака: 1 + " + c.Statistics.GetPrintablePoints(weapon.Power) + " + " + c.Statistics.GetPrintablePoints(attackPokememes) + "\n"

	if profileRaw.WeaponID != 0 {
		message += "\n🔫Оружие: " + weapon.Name + " " + c.Statistics.GetPrintablePoints(weapon.Power) + "⚔"
	}

	message += "\n🐱Покемемы:"
	for i := range profilePokememes {
		for j := range pokememes {
			if profilePokememes[i].PokememeID == pokememes[j].ID {
				message += "\n" + strconv.Itoa(pokememes[j].Grade)
				message += "⃣ " + pokememes[j].Name
				message += " +" + c.Statistics.GetPrintablePoints(profilePokememes[i].PokememeAttack) + "⚔"
			}
		}
	}
	message += "\nСтоимость покемемов на руках: " + c.Statistics.GetPrintablePoints(profileRaw.PokememesWealth) + "$"
	message += "\n\n💳" + strconv.Itoa(playerRaw.TelegramID)

	if playerRaw.Status == "owner" {
		message += "\n\nСтатус в боте: _владелец_"
	} else if playerRaw.Status == "admin" {
		message += "\n\nСтатус в боте: _администратор_"
	} else if playerRaw.Status == "academic" {
		message += "\n\nСтатус в боте: _академик_"
	} else {
		message += "\n\nСтатус в боте: _игрок_"
	}

	squadRoles, ok := c.Squader.GetUserRolesInSquads(playerRaw)
	if ok && len(squadRoles) > 0 {
		for i := range squadRoles {
			if squadRoles[i].UserRole == "commander" {
				message += "\nКомандир отряда " + squadRoles[i].Squad.Chat.Name
			} else {
				message += "\nУчастник отряда " + squadRoles[i].Squad.Chat.Name
			}
		}
	}

	message += "\n\n⏰Последнее обновление профиля: " + profileRaw.CreatedAt.Format("02.01.2006 15:04:05")
	message += "\nНе забывай обновляться, это важно для получения актуальной информации.\n\n"
	message += "/best – посмотреть лучших покемемов для поимки"

	c.Log.Debug(message)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

// UsersList lists all known users
func (u *Users) UsersList(update *tgbotapi.Update) string {
	pageNumber := strings.Replace(update.Message.Text, "/users", "", 1)
	pageNumber = strings.Replace(pageNumber, "/users", "", 1)
	page, _ := strconv.Atoi(pageNumber)
	if page == 0 {
		page = 1
	}
	usersArray, ok := u.getUsersWithProfiles()
	if !ok {
		return c.Talkers.BotError(update)
	}

	u.usersList(update, page, usersArray)
	return "ok"
}
