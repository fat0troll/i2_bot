// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
)

// Internal functions for Users package

func (u *Users) getUsersWithProfiles() ([]dbmapping.PlayerProfile, bool) {
	usersArray := []dbmapping.PlayerProfile{}
	players := []dbmapping.Player{}
	err := c.Db.Select(&players, "SELECT * FROM players")
	if err != nil {
		c.Log.Error(err)
		return usersArray, false
	}

	for i := range players {
		playerWithProfile := dbmapping.PlayerProfile{}
		profile, ok := u.GetProfile(players[i].ID)
		if !ok {
			playerWithProfile.HaveProfile = false
		} else {
			playerWithProfile.HaveProfile = true
		}
		playerWithProfile.Profile = profile
		playerWithProfile.Player = players[i]

		league := dbmapping.League{}
		if players[i].LeagueID != 0 {
			err = c.Db.Get(&league, c.Db.Rebind("SELECT * FROM leagues WHERE id=?"), players[i].LeagueID)
			if err != nil {
				c.Log.Error(err.Error())
				return usersArray, false
			}
		}
		playerWithProfile.League = league

		usersArray = append(usersArray, playerWithProfile)
	}

	return usersArray, true
}

func (u *Users) findUserByName(pattern string) ([]dbmapping.ProfileWithAddons, bool) {
	selectedUsers := []dbmapping.ProfileWithAddons{}

	err := c.Db.Select(&selectedUsers, c.Db.Rebind("SELECT * FROM (SELECT p.*, l.symbol AS league_symbol, l.id AS league_id, pl.telegram_id FROM players pl, profiles p, leagues l WHERE p.player_id = pl.id AND l.id = pl.league_id AND (p.nickname LIKE ? OR p.telegram_nickname LIKE ?) ORDER BY p.id desc LIMIT 100000) AS find_users_table GROUP BY player_id"), "%"+pattern+"%", "%"+pattern+"%")
	if err != nil {
		c.Log.Error(err.Error())
		return selectedUsers, false
	}

	return selectedUsers, true
}

func (u *Users) profileAddSuccessMessage(update *tgbotapi.Update, leagueID int, level int) {
	message := "*Профиль успешно обновлен.*\n\n"
	message += "Функциональность бота держится на актуальности профилей. Обновляйся почаще, и да пребудет с тобой Рандом!\n"
	message += "Сохраненный профиль ты можешь просмотреть командой /me.\n\n"
	message += "/best – посмотреть лучших покемемов для поимки"

	if leagueID == 1 {
		message += "\n/bastion — получить ссылку на БАСТИОН лиги\n"
		if level < 5 {
			message += "\n/academy — получить ссылку на АКАДЕМИЮ лиги\n"
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

func (u *Users) profileAddFailureMessage(update *tgbotapi.Update) {
	message := "*Неудачно получилось :(*\n\n"
	message += "Случилась жуткая ошибка, и мы не смогли записать профиль в базу. Напиши @fat0troll, он разберется."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

func (u *Users) usersList(update *tgbotapi.Update, page int, usersArray []dbmapping.PlayerProfile) {
	message := "*Зарегистрированные пользователи бота*\n"
	message += "Список отсортирован по ID регистрации.\n"
	message += "Количество зарегистрированных пользователей: " + strconv.Itoa(len(usersArray)) + "\n"
	message += "Отображаем пользователей с " + strconv.Itoa(((page-1)*25)+1) + " по " + strconv.Itoa(page*25) + "\n"
	if len(usersArray) > page*25 {
		message += "Переход на следующую страницу: /users" + strconv.Itoa(page+1)
	}
	if page > 1 {
		message += "\nПереход на предыдущую страницу: /users" + strconv.Itoa(page-1)
	}
	message += "\n\n"

	for i := range usersArray {
		if (i+1 > 25*(page-1)) && (i+1 < (25*page)+1) {
			message += "#" + strconv.Itoa(usersArray[i].Player.ID)
			if usersArray[i].HaveProfile {
				message += " " + usersArray[i].League.Symbol
				message += " " + usersArray[i].Profile.Nickname
				if usersArray[i].Profile.TelegramNickname != "" {
					message += " (@" + u.FormatUsername(usersArray[i].Profile.TelegramNickname) + ")"
				}
				message += " /profile" + strconv.Itoa(usersArray[i].Player.ID) + "\n"
				message += "Telegram ID: " + strconv.Itoa(usersArray[i].Player.TelegramID) + "\n"
				message += "Последнее обновление: " + usersArray[i].Profile.CreatedAt.Format("02.01.2006 15:04:05") + "\n"
			} else {
				if usersArray[i].Player.Status == "special" {
					message += " _суперюзер_\n"
				} else {
					message += " _без профиля_\n"
				}
				message += "Telegram ID: " + strconv.Itoa(usersArray[i].Player.TelegramID) + "\n"
			}
		}
	}

	if len(usersArray) > page*25 {
		message += "\n"
		message += "Переход на следующую страницу: /users" + strconv.Itoa(page+1)
	}
	if page > 1 {
		message += "\nПереход на предыдущую страницу: /users" + strconv.Itoa(page-1)
	}

	message += "\nЧтобы добавить пользователя в отряд, введите команду /squad\\_add\\_user _X Y_ или /squad\\_add\\_commander _X Y_, где _X_ — ID отряда (посмотреть все отряды можно командой /squads), а _Y_ — ID пользователя из списка выше."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
