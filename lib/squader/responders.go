// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squader

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
)

// SquadsList lists all squads
func (s *Squader) SquadsList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	if !c.Users.PlayerBetterThan(playerRaw, "admin") {
		if s.isUserAnyCommander(playerRaw.ID) {
			return c.Talkers.AnyMessageUnauthorized(update)
		}
	}
	squads, ok := s.getAllSquadsWithChats()
	if !ok {
		return "fail"
	}

	message := "*Наши отряды:*\n"

	for i := range squads {
		message += "---\n"
		message += "[#" + strconv.Itoa(squads[i].Squad.ID) + "] _" + squads[i].Chat.Name
		message += "_ /show\\_squad" + strconv.Itoa(squads[i].Squad.ID) + "\n"
		message += "Telegram ID: " + strconv.FormatInt(squads[i].Chat.TelegramID, 10) + "\n"
		message += "Флудилка отряда: _" + squads[i].FloodChat.Name + "_\n"
		message += "Статистика отряда:\n"
		message += c.Statistics.SquadStatictics(squads[i].Squad.ID)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

// SquadInfo returns statistic and list of squad players
func (s *Squader) SquadInfo(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	squadNumber := strings.Replace(update.Message.Text, "/show_squad", "", 1)
	squadID, _ := strconv.Atoi(squadNumber)
	if squadID == 0 {
		squadID = 1
	}

	if !c.Users.PlayerBetterThan(playerRaw, "admin") {
		if s.getUserRoleForSquad(squadID, playerRaw.ID) != "commander" {
			return c.Talkers.AnyMessageUnauthorized(update)
		}
	}

	squad, ok := s.GetSquadByID(squadID)
	if !ok {
		return c.Talkers.BotError(update)
	}

	message := "*Информация об отряде* _" + squad.Chat.Name + "_*:*\n"
	message += c.Statistics.SquadStatictics(squad.Squad.ID)
	message += "\n"

	squadMembers, ok := s.getPlayersForSquad(squad.Squad.ID)
	if !ok {
		return "fail"
	}
	if len(squadMembers) > 0 {
		message += "Участники отряда:\n"
		for i := range squadMembers {
			message += "#" + strconv.Itoa(squadMembers[i].Player.ID)
			if squadMembers[i].UserRole == "commander" {
				message += " \\[К]"
			}
			if squadMembers[i].Player.Status == "special" {
				message += " _суперюзер_"
			} else {
				message += " " + squadMembers[i].Profile.Nickname + " "
				if squadMembers[i].Profile.TelegramNickname != "" {
					message += "(@" + c.Users.FormatUsername(squadMembers[i].Profile.TelegramNickname) + ")"
				}
			}
			message += " ⚔" + strconv.Itoa(squadMembers[i].Profile.Power)
			message += "\n"
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
