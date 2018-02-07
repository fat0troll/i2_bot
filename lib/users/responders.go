// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
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

// ProfileAddEffectsMessage shows when user tries to post profile with effects enabled
func (u *Users) ProfileAddEffectsMessage(update *tgbotapi.Update) string {
	message := "*Наркоман, штоле?*\n\n"
	message += "Бот не принимает профили во время активированных эффектов. Закончи свои дела и принеси чистый профиль через полчаса."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "fail"
}

// ProfileMessage shows current player's profile
func (u *Users) ProfileMessage(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	profileRaw, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return c.Talkers.AnyMessageUnauthorized(update)
	}
	league := dbmapping.League{}
	err = c.Db.Get(&league, c.Db.Rebind("SELECT * FROM leagues WHERE id=?"), playerRaw.LeagueID)
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
	message += "\n⚔Атака: " + c.Statistics.GetPrintablePoints(weapon.Power) + " + " + c.Statistics.GetPrintablePoints(attackPokememes) + "\n"

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
