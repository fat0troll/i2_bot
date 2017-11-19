// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squader

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (s *Squader) getAllSquadsWithChats() ([]dbmapping.SquadChat, bool) {
	squadsWithChats := []dbmapping.SquadChat{}
	squads := []dbmapping.Squad{}

	err := c.Db.Select(&squads, "SELECT * FROM squads")
	if err != nil {
		c.Log.Error(err)
		return squadsWithChats, false
	}

	for i := range squads {
		chatSquad := dbmapping.SquadChat{}
		chat := dbmapping.Chat{}
		floodChat := dbmapping.Chat{}
		err = c.Db.Get(&chat, c.Db.Rebind("SELECT * FROM chats WHERE id=?"), squads[i].ChatID)
		if err != nil {
			c.Log.Error(err)
			return squadsWithChats, false
		}
		err = c.Db.Get(&floodChat, c.Db.Rebind("SELECT * FROM chats WHERE id=?"), squads[i].FloodChatID)
		if err != nil {
			c.Log.Error(err)
			return squadsWithChats, false
		}

		chatSquad.Squad = squads[i]
		chatSquad.Chat = chat
		chatSquad.FloodChat = floodChat

		squadsWithChats = append(squadsWithChats, chatSquad)
	}

	return squadsWithChats, true
}

func (s *Squader) createSquad(update *tgbotapi.Update, chatID int, floodChatID int) (dbmapping.Squad, string) {
	squad := dbmapping.Squad{}
	chat := dbmapping.Chat{}
	floodChat := dbmapping.Chat{}

	// Checking if chats in database exist
	err := c.Db.Get(&chat, c.Db.Rebind("SELECT * FROM chats WHERE id=?"), chatID)
	if err != nil {
		c.Log.Error(err)
		return squad, "fail"
	}
	err = c.Db.Get(&floodChat, c.Db.Rebind("SELECT * FROM chats WHERE id=?"), floodChatID)
	if err != nil {
		c.Log.Error(err)
		return squad, "fail"
	}

	err2 := c.Db.Get(&squad, c.Db.Rebind("SELECT * FROM squads WHERE chat_id IN (?, ?) OR flood_chat_id IN (?, ?)"), chat.ID, floodChat.ID, chat.ID, floodChat.ID)
	if err2 == nil {
		return squad, "dup"
	}
	c.Log.Debug(err2)

	err = c.Db.Get(&squad, c.Db.Rebind("SELECT * FROM squads WHERE chat_id=? AND flood_chat_id=?"), chatID, floodChatID)
	if err != nil {
		c.Log.Debug(err)

		playerRaw, ok := c.Getters.GetOrCreatePlayer(update.Message.From.ID)
		if !ok {
			return squad, "fail"
		}

		squad.AuthorID = playerRaw.ID
		squad.ChatID = chatID
		squad.FloodChatID = floodChatID
		squad.CreatedAt = time.Now().UTC()

		_, err = c.Db.NamedExec("INSERT INTO `squads` VALUES(NULL, :chat_id, :flood_chat_id, :author_id, :created_at)", &squad)
		if err != nil {
			c.Log.Error(err)
			return squad, "fail"
		}

		err = c.Db.Get(&squad, c.Db.Rebind("SELECT * FROM squads WHERE chat_id=? AND flood_chat_id=?"), chatID, floodChatID)
		if err != nil {
			c.Log.Error(err)
			return squad, "fail"
		}

		return squad, "ok"
	}

	return squad, "dup"
}

func (s *Squader) getSquadByChatID(update *tgbotapi.Update, chatID int) (dbmapping.Squad, string) {
	squad := dbmapping.Squad{}
	chat := dbmapping.Chat{}

	// Checking if chat in database exist
	err := c.Db.Get(&chat, c.Db.Rebind("SELECT * FROM chats WHERE id=?"), chatID)
	if err != nil {
		c.Log.Error(err)
		return squad, "fail"
	}

	err = c.Db.Get(&squad, c.Db.Rebind("SELECT * FROM squads WHERE chat_id=?"), chat.ID)
	if err != nil {
		c.Log.Error(err)
		return squad, "fail"
	}

	return squad, "ok"
}

func (s *Squader) squadCreationDuplicate(update *tgbotapi.Update) string {
	message := "*Отряд уже существует*\n"
	message += "Проверьте, правильно ли вы ввели команду, и повторите попытку."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "fail"
}

func (s *Squader) squadCreationFailure(update *tgbotapi.Update) string {
	message := "*Не удалось добавить отряд в базу*\n"
	message += "Проверьте, правильно ли вы ввели команду, и повторите попытку."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "fail"
}

func (s *Squader) squadCreationSuccess(update *tgbotapi.Update) string {
	message := "*Отряд успешно добавлен в базу*\n"
	message += "Просмотреть список отрядов можно командой /squads."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "fail"
}

// External functions

// CreateSquad creates new squad from chat if not already exist
func (s *Squader) CreateSquad(update *tgbotapi.Update) string {
	commandArugments := update.Message.CommandArguments()
	argumentsRx := regexp.MustCompile(`(\d+)\s(\d+)`)

	if !argumentsRx.MatchString(commandArugments) {
		return s.squadCreationFailure(update)
	}

	chatNumbers := strings.Split(commandArugments, " ")
	if len(chatNumbers) < 2 {
		return s.squadCreationFailure(update)
	}
	chatID, _ := strconv.Atoi(chatNumbers[0])
	if chatID == 0 {
		return s.squadCreationFailure(update)
	}
	floodChatID, _ := strconv.Atoi(chatNumbers[1])
	if floodChatID == 0 {
		return s.squadCreationFailure(update)
	}

	_, ok := s.createSquad(update, chatID, floodChatID)
	if ok == "fail" {
		return s.squadCreationFailure(update)
	} else if ok == "dup" {
		return s.squadCreationDuplicate(update)
	}

	return s.squadCreationSuccess(update)
}

// SquadsList lists all squads
func (s *Squader) SquadsList(update *tgbotapi.Update) string {
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
		message += s.SquadStatictics(squads[i].Squad.ID)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

// SquadStatictics generates statistics message snippet. Public due to usage in chats list
func (s *Squader) SquadStatictics(squadID int) string {
	squadMembersWithInformation := []dbmapping.SquadPlayerFull{}
	squadMembers := []dbmapping.SquadPlayer{}
	squad := dbmapping.Squad{}

	err := c.Db.Get(&squad, c.Db.Rebind("SELECT * FROM squads WHERE id=?"), squadID)
	if err != nil {
		c.Log.Error(err.Error())
		return "Отряда не существует!"
	}

	err = c.Db.Select(&squadMembers, c.Db.Rebind("SELECT * FROM squads_players WHERE squad_id=?"), squadID)
	if err != nil {
		c.Log.Error(err.Error())
		return "Невозможно получить информацию о данном отряде. Возможно, он пуст или произошла ошибка."
	}

	for i := range squadMembers {
		fullInfo := dbmapping.SquadPlayerFull{}

		playerRaw, _ := c.Getters.GetPlayerByID(squadMembers[i].PlayerID)
		profileRaw, _ := c.Getters.GetProfile(playerRaw.ID)

		fullInfo.Squad = squad
		fullInfo.Player = playerRaw
		fullInfo.Profile = profileRaw

		squadMembersWithInformation = append(squadMembersWithInformation, fullInfo)
	}

	message := "Количество человек в отряде: " + strconv.Itoa(len(squadMembersWithInformation))
	message += "\n"

	return message
}
