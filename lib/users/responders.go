// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

// FormatUsername formats Telegram username for posting
func (u *Users) FormatUsername(userName string) string {
	return strings.Replace(userName, "_", `\_`, -1)
}

// FindByLevel finds user with level and recent profile update
func (u *Users) FindByLevel(update *tgbotapi.Update) string {
	commandArgs := update.Message.CommandArguments()
	if commandArgs == "" {
		c.Talkers.BotError(update)
		return "fail"
	}

	levelID, err := strconv.Atoi(commandArgs)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	users := u.findUsersByLevel(levelID)

	u.foundUsersMessage(update, users)

	return "ok"
}

// FindByName finds user with such username or nickname
func (u *Users) FindByName(update *tgbotapi.Update) string {
	commandArgs := update.Message.CommandArguments()
	if commandArgs == "" {
		c.Talkers.BotError(update)
		return "fail"
	}

	users := u.findUserByName(commandArgs)

	u.foundUsersMessage(update, users)

	return "ok"
}

// FindByTopAttack finds user by top-attack rating
func (u *Users) FindByTopAttack(update *tgbotapi.Update) string {
	commandArgs := update.Message.CommandArguments()
	if commandArgs == "" {
		c.Talkers.BotError(update)
		return "fail"
	}

	attackInt, err := strconv.Atoi(commandArgs)
	if err != nil {
		c.Log.Error(err.Error())
		c.Talkers.BotError(update)
		return "fail"
	}

	users := u.findUserByTopAttack(attackInt)

	u.foundUsersMessage(update, users)

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

	playerRaw, err := c.DataCache.GetPlayerByID(userID)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	_, err = c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return c.Talkers.BotError(update)
	}

	return u.ProfileMessage(update, playerRaw)
}

// ProfileMessage shows current player's profile
func (u *Users) ProfileMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	profileRaw, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return c.Talkers.AnyMessageUnauthorized(update)
	}
	league, err := c.DataCache.GetLeagueByID(playerRaw.LeagueID)
	if err != nil {
		c.Log.Error(err.Error())
		return c.Talkers.BotError(update)
	}
	level, err := c.DataCache.GetLevelByID(profileRaw.LevelID)
	if err != nil {
		c.Log.Error(err)
	}
	weapon, err := c.DataCache.GetWeaponTypeByID(profileRaw.WeaponID)
	if err != nil {
		// It's non critical
		c.Log.Debug(err.Error())
	}
	profilePokememes := []dbmapping.ProfilePokememe{}
	err = c.Db.Select(&profilePokememes, c.Db.Rebind("SELECT * FROM profiles_pokememes WHERE profile_id=?"), profileRaw.ID)
	if err != nil {
		c.Log.Error(err)
	}
	pokememes := c.DataCache.GetAllPokememes()

	attackPokememes := 0
	for i := range profilePokememes {
		for j := range pokememes {
			if profilePokememes[i].PokememeID == pokememes[j].Pokememe.ID {
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
	message += " |💎" + strconv.Itoa(profileRaw.Crystals)
	message += " |⭕" + strconv.Itoa(profileRaw.Pokeballs)
	if weapon != nil {
		message += "\n⚔Атака: " + c.Statistics.GetPrintablePoints(weapon.Power) + " + " + c.Statistics.GetPrintablePoints(attackPokememes) + "\n"
	} else {
		message += "\n⚔Атака: " + c.Statistics.GetPrintablePoints(attackPokememes) + "\n"
	}

	if profileRaw.WeaponID != 0 {
		message += "\n🔫Оружие: " + weapon.Name + " " + c.Statistics.GetPrintablePoints(weapon.Power) + "⚔"
	}

	message += "\n🐱Покемемы:"
	for i := range profilePokememes {
		for j := range pokememes {
			if profilePokememes[i].PokememeID == pokememes[j].Pokememe.ID {
				message += "\n *[" + strconv.Itoa(pokememes[j].Pokememe.Grade)
				message += "]* " + pokememes[j].Pokememe.Name
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

	squadRoles := c.DataCache.GetUserRolesInSquads(playerRaw.ID)
	if len(squadRoles) > 0 {
		for i := range squadRoles {
			if squadRoles[i].UserRole == "commander" {
				message += "\nКомандир отряда " + squadRoles[i].Squad.Chat.Name
			} else {
				message += "\nУчастник отряда " + squadRoles[i].Squad.Chat.Name
			}
		}
	}

	message += "\nКарма: " + strconv.Itoa(playerRaw.Karma)

	message += "\n\n⏰Последнее обновление профиля: " + profileRaw.CreatedAt.Format("02.01.2006 15:04:05")
	message += "\nНе забывай обновляться, это важно для получения актуальной информации.\n\n"
	message += "/best – посмотреть лучших покемемов для поимки\n"
	message += "/advice – посмотреть самых дорогих покемемов для поимки\n"
	message += "/top — посмотреть лучших игроков лиги\n"
	message += "/top\\_my — посмотреть лучших игроков лиги твоего уровня\n"

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
	users := c.DataCache.GetPlayersWithCurrentProfiles()

	u.usersList(update, page, users)
	return "ok"
}
