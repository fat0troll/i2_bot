// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squader

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

func (s *Squader) isUserAnyCommander(playerID int) bool {
	userRoles := c.DataCache.GetUserRolesInSquads(playerID)
	for i := range userRoles {
		c.Log.Debug("Role for squad with ID=" + strconv.Itoa(userRoles[i].Squad.Squad.ID) + ":" + userRoles[i].UserRole)
		if userRoles[i].UserRole == "commander" {
			return true
		}
	}

	return false
}

func (s *Squader) squadUserAdditionFailure(update *tgbotapi.Update) string {
	message := "*Не удалось добавить игрока в отряд*\n"
	message += "Проверьте, правильно ли вы ввели команду, и повторите попытку. Кроме того, возможно, что у пользователя нет профиля в боте."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "fail"
}

func (s *Squader) squadUserAdditionSuccess(update *tgbotapi.Update) string {
	message := "*Игрок добавлен в отряд*\n"
	message += "Теперь вы можете дать ему ссылку для входа в чаты отряда."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

// External functions

// AddUserToSquad adds user to squad
func (s *Squader) AddUserToSquad(update *tgbotapi.Update, adderRaw *dbmapping.Player) string {
	command := update.Message.Command()
	commandArugments := update.Message.CommandArguments()
	userType := "user"
	if command == "squad_add_commander" {
		userType = "commander"
	}
	argumentsRx := regexp.MustCompile(`(\d+)\s(\d+)`)
	if !argumentsRx.MatchString(commandArugments) {
		return s.squadUserAdditionFailure(update)
	}

	argumentNumbers := strings.Split(commandArugments, " ")
	if len(argumentNumbers) < 2 {
		return s.squadUserAdditionFailure(update)
	}
	squadID, _ := strconv.Atoi(argumentNumbers[0])
	if squadID == 0 {
		return s.squadUserAdditionFailure(update)
	}
	playerID, _ := strconv.Atoi(argumentNumbers[1])
	if playerID == 0 {
		return s.squadUserAdditionFailure(update)
	}

	playerRaw, err := c.DataCache.GetPlayerByID(playerID)
	if err != nil {
		c.Log.Error(err.Error())
		return s.squadUserAdditionFailure(update)
	}
	profileRaw, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return s.squadUserAdditionFailure(update)
	}
	squadRaw, err := c.DataCache.GetSquadByID(squadID)
	if err != nil {
		c.Log.Error(err.Error())
		return s.squadUserAdditionFailure(update)
	}

	if !c.Users.PlayerBetterThan(playerRaw, "admin") {
		_, err = c.DataCache.GetProfileByPlayerID(playerRaw.ID)
		if err != nil {
			c.Log.Error(err.Error())
			return s.squadUserAdditionFailure(update)
		}
	}

	if !c.Users.PlayerBetterThan(adderRaw, "admin") {
		userRoles := c.DataCache.GetUserRolesInSquads(adderRaw.ID)
		isCommander := false
		for i := range userRoles {
			if userRoles[i].UserRole == "commander" {
				if userRoles[i].Squad.Squad.ID == squadRaw.Squad.ID {
					isCommander = true
				}
			}
		}

		if !isCommander {
			return c.Talkers.AnyMessageUnauthorized(update)
		}
	}

	if !c.Users.PlayerBetterThan(playerRaw, "admin") {
		if playerRaw.LeagueID != 1 {
			return s.squadUserAdditionFailure(update)
		} else if userType != "commander" {
			if squadRaw.Squad.MinLevel > profileRaw.LevelID {
				c.Log.Debug("Levels mismatch: min"+strconv.Itoa(squadRaw.Squad.MinLevel), ", player: "+strconv.Itoa(profileRaw.LevelID))
				return s.squadUserAdditionFailure(update)
			} else if squadRaw.Squad.MaxLevel-1 < profileRaw.LevelID {
				c.Log.Debug("Levels mismatch: max"+strconv.Itoa(squadRaw.Squad.MaxLevel), ", player: "+strconv.Itoa(profileRaw.LevelID))
				return s.squadUserAdditionFailure(update)
			}
		}
	}

	// All checks are passed here, creating new item in database
	playerSquad := dbmapping.SquadPlayer{}
	playerSquad.SquadID = squadRaw.Squad.ID
	playerSquad.PlayerID = playerRaw.ID
	playerSquad.UserType = userType
	playerSquad.AuthorID = adderRaw.ID
	playerSquad.CreatedAt = time.Now().UTC()

	_, err = c.DataCache.AddPlayerToSquad(&playerSquad)
	if err != nil {
		c.Log.Error(err.Error())
		return s.squadUserAdditionFailure(update)
	}

	message := "Привет! Тебя добавили в отряд «" + squadRaw.Chat.Name + "»\n"
	message += "Присоединиться к чату отряда тут: " + squadRaw.Squad.InviteLink

	msg := tgbotapi.NewMessage(int64(playerRaw.TelegramID), message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return s.squadUserAdditionSuccess(update)
}
