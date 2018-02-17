// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"sort"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
	"time"
)

// Internal functions for Users package

func (u *Users) findUsersByLevel(levelID int) map[int]*dbmapping.PlayerProfile {
	selectedUsers := make(map[int]*dbmapping.PlayerProfile)
	allUsers := c.DataCache.GetPlayersWithCurrentProfiles()

	for i := range allUsers {
		if allUsers[i].Profile.LevelID == levelID {
			if allUsers[i].Player.UpdatedAt.After(time.Now().UTC().Add(-72 * time.Hour)) {
				selectedUsers[i] = allUsers[i]
			}
		}
	}

	return selectedUsers
}

func (u *Users) findUserByName(pattern string) map[int]*dbmapping.PlayerProfile {
	selectedUsers := make(map[int]*dbmapping.PlayerProfile)
	allUsers := c.DataCache.GetPlayersWithCurrentProfiles()

	for i := range allUsers {
		matchedPattern := false
		if strings.Contains(strings.ToLower(allUsers[i].Profile.Nickname), strings.ToLower(pattern)) {
			matchedPattern = true
		}
		if strings.Contains(strings.ToLower(allUsers[i].Profile.TelegramNickname), strings.ToLower(pattern)) {
			matchedPattern = true
		}
		if matchedPattern {
			selectedUsers[i] = allUsers[i]
		}
	}

	return selectedUsers
}

func (u *Users) findUserByTopAttack(power int) map[int]*dbmapping.PlayerProfile {
	selectedUsers := make(map[int]*dbmapping.PlayerProfile)
	allPlayers := c.DataCache.GetPlayersWithCurrentProfiles()

	profiles := make([]*dbmapping.PlayerProfile, 0)

	for i := range allPlayers {
		if allPlayers[i].Player.LeagueID == 1 {
			profiles = append(profiles, allPlayers[i])
		}
	}

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Profile.Power > profiles[j].Profile.Power
	})

	for i := (power - 1); i < (power + 2); i++ {
		if profiles[i] != nil {
			selectedUsers[i] = profiles[i]
		}
	}

	return selectedUsers
}

func (u *Users) foundUsersMessage(update *tgbotapi.Update, users map[int]*dbmapping.PlayerProfile) {
	var keys []int
	for i := range users {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	message := "*Найденные игроки:*\n"

	for _, i := range keys {
		message += "#" + strconv.Itoa(users[i].Player.ID)
		if users[i].HaveProfile {
			message += " " + users[i].League.Symbol
			message += " " + users[i].Profile.Nickname
			if users[i].Profile.TelegramNickname != "" {
				message += " (@" + u.FormatUsername(users[i].Profile.TelegramNickname) + ")"
			}
		}
		message += " /profile" + strconv.Itoa(users[i].Player.ID) + "\n"
		message += "Telegram ID: " + strconv.Itoa(users[i].Player.TelegramID) + "\n"
		message += "Последнее обновление: " + users[i].Profile.CreatedAt.Format("02.01.2006 15:04:05") + "\n"

		if len(message) > 2000 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			msg.ParseMode = "Markdown"

			c.Bot.Send(msg)

			message = ""
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}

func (u *Users) profileAddSuccessMessage(update *tgbotapi.Update, leagueID int, level int) {
	message := "*Профиль успешно обновлен.*\n\n"
	message += "Функциональность бота держится на актуальности профилей. Обновляйся почаще, и да пребудет с тобой Рандом!\n"
	message += "Сохраненный профиль ты можешь просмотреть командой /me.\n\n"
	message += "/best – посмотреть лучших покемемов для поимки\n"
	message += "/advice – посмотреть самых дорогих покемемов для поимки\n"
	message += "/top — посмотреть лучших представителей лиги\n"
	message += "/top\\_my — посмотреть лучших представителей лиги твоего уровня\n"

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

func (u *Users) usersList(update *tgbotapi.Update, page int, users map[int]*dbmapping.PlayerProfile) {
	message := "*Зарегистрированные пользователи бота*\n"
	message += "Список отсортирован по ID регистрации.\n"
	message += "Количество зарегистрированных пользователей: " + strconv.Itoa(len(users)) + "\n"
	message += "Отображаем пользователей с " + strconv.Itoa(((page-1)*25)+1) + " по " + strconv.Itoa(page*25) + "\n"
	if len(users) > page*25 {
		message += "Переход на следующую страницу: /users" + strconv.Itoa(page+1)
	}
	if page > 1 {
		message += "\nПереход на предыдущую страницу: /users" + strconv.Itoa(page-1)
	}
	message += "\n\n"

	var keys []int
	for i := range users {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, i := range keys {
		if (i+1 > 25*(page-1)) && (i+1 < (25*page)+1) {
			message += "#" + strconv.Itoa(users[i].Player.ID)
			if users[i].HaveProfile {
				message += " " + users[i].League.Symbol
				message += " " + users[i].Profile.Nickname
				if users[i].Profile.TelegramNickname != "" {
					message += " (@" + u.FormatUsername(users[i].Profile.TelegramNickname) + ")"
				}
				message += " /profile" + strconv.Itoa(users[i].Player.ID) + "\n"
				message += "Telegram ID: " + strconv.Itoa(users[i].Player.TelegramID) + "\n"
				message += "Последнее обновление: " + users[i].Profile.CreatedAt.Format("02.01.2006 15:04:05") + "\n"
			} else {
				if users[i].Player.Status == "special" {
					message += " _суперюзер_\n"
				} else {
					message += " _без профиля_\n"
				}
				message += "Telegram ID: " + strconv.Itoa(users[i].Player.TelegramID) + "\n"
			}
		}
	}

	if len(users) > page*25 {
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
