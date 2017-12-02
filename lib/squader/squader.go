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

func (s *Squader) getPlayersForSquad(squadID int) ([]dbmapping.SquadPlayerFull, bool) {
	players := []dbmapping.SquadPlayerFull{}
	playersRaw := []dbmapping.Player{}
	squadPlayers := []dbmapping.SquadPlayer{}

	squad, ok := s.GetSquadByID(squadID)
	if !ok {
		return players, false
	}

	err := c.Db.Select(&playersRaw, c.Db.Rebind("SELECT p.* FROM players p, squads_players sp WHERE p.id = sp.player_id AND sp.squad_id=?"), squad.Squad.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return players, false
	}

	err = c.Db.Select(&squadPlayers, c.Db.Rebind("SELECT * FROM squads_players WHERE squad_id=?"), squad.Squad.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return players, false
	}

	for i := range playersRaw {
		for ii := range squadPlayers {
			if squadPlayers[ii].PlayerID == playersRaw[i].ID {
				playerWithProfile := dbmapping.SquadPlayerFull{}
				profile, _ := c.Users.GetProfile(playersRaw[i].ID)
				playerWithProfile.Profile = profile
				playerWithProfile.Player = playersRaw[i]
				playerWithProfile.Squad = squad
				playerWithProfile.UserRole = squadPlayers[ii].UserType

				players = append(players, playerWithProfile)
			}
		}
	}

	return players, true
}

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

		playerRaw, ok := c.Users.GetOrCreatePlayer(update.Message.From.ID)
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

func (s *Squader) getUserRoleForSquad(squadID int, playerID int) string {
	squadPlayer := dbmapping.SquadPlayer{}
	err := c.Db.Get(&squadPlayer, c.Db.Rebind("SELECT * FROM squads_players WHERE squad_id=? AND player_id=?"), squadID, playerID)
	if err != nil {
		c.Log.Debug(err.Error())
		return "nobody"
	}

	return squadPlayer.UserType
}

func (s *Squader) deleteFloodMessage(update *tgbotapi.Update) {
	deleteMessageConfig := tgbotapi.DeleteMessageConfig{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.MessageID,
	}

	_, err := c.Bot.DeleteMessage(deleteMessageConfig)
	if err != nil {
		c.Log.Error(err.Error())
	}
}

func (s *Squader) isUserAnyCommander(playerID int) bool {
	squadPlayers := []dbmapping.SquadPlayer{}
	err := c.Db.Select(&squadPlayers, c.Db.Rebind("SELECT * FROM squads_players WHERE player_id=? AND user_type='commander'"), playerID)
	if err != nil {
		c.Log.Debug(err.Error())
	}

	if len(squadPlayers) > 0 {
		return true
	}

	return false
}

func (s *Squader) getCommandersForSquadViaChat(chatRaw *dbmapping.Chat) ([]dbmapping.Player, bool) {
	commanders := []dbmapping.Player{}
	err := c.Db.Select(&commanders, c.Db.Rebind("SELECT p.* FROM players p, squads_players sp, squads s WHERE (s.chat_id=? OR s.flood_chat_id=?) AND sp.squad_id = s.id AND sp.user_type = 'commander' AND sp.player_id = p.id"), chatRaw.ID, chatRaw.ID)
	if err != nil {
		c.Log.Debug(err.Error())
		return commanders, false
	}

	return commanders, true
}

func (s *Squader) kickUserFromSquadChat(user *tgbotapi.User, chatRaw *dbmapping.Chat) {
	chatUserConfig := tgbotapi.ChatMemberConfig{
		ChatID: chatRaw.TelegramID,
		UserID: user.ID,
	}

	kickConfig := tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: chatUserConfig,
		UntilDate:        1893456000,
	}

	_, err := c.Bot.KickChatMember(kickConfig)
	if err != nil {
		c.Log.Error(err.Error())
	}

	suerName := ""
	if user.UserName != "" {
		suerName = "@" + user.UserName
	} else {
		suerName = user.FirstName
		if user.LastName != "" {
			suerName += " " + user.LastName
		}
	}

	bastionChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
	hqChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
	if chatRaw.TelegramID != bastionChatID {
		// In Bastion notifications are public in default chat
		commanders, ok := s.getCommandersForSquadViaChat(chatRaw)
		if ok {
			for i := range commanders {
				message := "Некто " + c.Users.FormatUsername(suerName) + " попытался зайти в чат _" + chatRaw.Name + "_ и был изгнан ботом, так как не имеет права посещать этот чат."

				msg := tgbotapi.NewMessage(int64(commanders[i].TelegramID), message)
				msg.ParseMode = "Markdown"
				c.Bot.Send(msg)
			}
		}
	} else {
		message := "Некто " + c.Users.FormatUsername(suerName) + " попытался зайти в чат _Бастион Инстинкта_ и был изгнан ботом, так как не имеет права посещать этот чат."

		msg := tgbotapi.NewMessage(hqChatID, message)
		msg.ParseMode = "Markdown"
		c.Bot.Send(msg)
	}
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

	playerRaw, ok := c.Users.GetPlayerByID(playerID)
	if !ok {
		return s.squadUserAdditionFailure(update)
	}
	squadRaw := dbmapping.Squad{}
	err := c.Db.Get(&squadRaw, c.Db.Rebind("SELECT * FROM squads WHERE id=?"), squadID)
	if err != nil {
		c.Log.Error(err.Error())
		return s.squadUserAdditionFailure(update)
	}

	if !c.Users.PlayerBetterThan(&playerRaw, "admin") {
		_, ok = c.Users.GetProfile(playerRaw.ID)
		if !ok {
			return s.squadUserAdditionFailure(update)
		}
	}

	if !c.Users.PlayerBetterThan(adderRaw, "admin") {
		if userType == "commander" {
			return c.Talkers.AnyMessageUnauthorized(update)
		}

		if s.getUserRoleForSquad(squadRaw.ID, adderRaw.ID) != "commander" {
			return c.Talkers.AnyMessageUnauthorized(update)
		}
	}

	if !c.Users.PlayerBetterThan(&playerRaw, "admin") {
		if playerRaw.LeagueID != 1 {
			return s.squadUserAdditionFailure(update)
		}
	}

	// All checks are passed here, creating new item in database
	playerSquad := dbmapping.SquadPlayer{}
	playerSquad.SquadID = squadRaw.ID
	playerSquad.PlayerID = playerRaw.ID
	playerSquad.UserType = userType
	playerSquad.AuthorID = adderRaw.ID
	playerSquad.CreatedAt = time.Now().UTC()

	_, err = c.Db.NamedExec("INSERT INTO squads_players VALUES(NULL, :squad_id, :player_id, :user_type, :author_id, :created_at)", &playerSquad)
	if err != nil {
		c.Log.Error(err.Error())
		return s.squadUserAdditionFailure(update)
	}

	return s.squadUserAdditionSuccess(update)
}

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

// ProcessMessage handles all squad-specified administration actions
func (s *Squader) ProcessMessage(update *tgbotapi.Update, chatRaw *dbmapping.Chat) string {
	// It will pass message or do some extra actions
	// If it returns "ok", we can pass message to router, otherwise we will stop here
	processMain := false
	processFlood := false
	messageProcessed := false
	switch s.IsChatASquadEnabled(chatRaw) {
	case "main":
		processMain = true
	case "flood":
		processFlood = true
	default:
		return "ok"
	}

	// Kicking non-squad members from any chat
	if processMain || processFlood {
		if update.Message.NewChatMembers != nil {
			newUsers := *update.Message.NewChatMembers
			if len(newUsers) > 0 {
				for i := range newUsers {
					playerRaw, ok := c.Users.GetOrCreatePlayer(newUsers[i].ID)
					if !ok {
						s.kickUserFromSquadChat(&newUsers[i], chatRaw)
						messageProcessed = true
					}

					availableChats, ok := s.GetAvailableSquadChatsForUser(&playerRaw)
					if !ok {
						s.kickUserFromSquadChat(&newUsers[i], chatRaw)
						messageProcessed = true
					}

					isChatValid := false
					for i := range availableChats {
						if availableChats[i] == *chatRaw {
							isChatValid = true
						}
					}

					// Dirty hack
					if update.Message.Chat.ID == -1001396321727 || update.Message.Chat.ID == -1001310954317 {
						isChatValid = true
					}

					if !isChatValid {
						switch strings.ToLower(newUsers[i].UserName) {
						case "gantz_yaka":
							messageProcessed = true
						case "agentpb":
							messageProcessed = true
						case "pbhelp":
							messageProcessed = true
						default:
							s.kickUserFromSquadChat(&newUsers[i], chatRaw)
							messageProcessed = true
						}
					}
				}
			}
		}
	}

	if processMain {
		c.Log.Debug("Message found in one of squad's main chats.")
		talker, ok := c.Users.GetOrCreatePlayer(update.Message.From.ID)
		if !ok {
			s.deleteFloodMessage(update)
			messageProcessed = true
		}

		if (update.Message.From.UserName != "i2_bot") && (update.Message.From.UserName != "i2_bot_dev") && !s.isUserAnyCommander(talker.ID) {
			s.deleteFloodMessage(update)
			messageProcessed = true
		}
	}

	if messageProcessed {
		return "fail"
	}

	return "ok"
}

// ProtectBastion avoids spies and no-profile players to join Bastion
func (s *Squader) ProtectBastion(update *tgbotapi.Update, newUser *tgbotapi.User) string {
	defaultChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.DefaultID, 10, 64)
	userName := ""
	if newUser.UserName != "" {
		userName += "@" + newUser.UserName
	} else {
		userName += newUser.FirstName
		if newUser.LastName != "" {
			userName += " " + newUser.LastName
		}
	}

	chatRaw, ok := c.Chatter.GetOrCreateChat(update)
	if !ok {
		return "fail"
	}

	playerRaw, ok := c.Users.GetOrCreatePlayer(newUser.ID)
	if !ok {
		switch newUser.UserName {
		case "gantz_yaka":
			// do nothing
		case "@agentpb":
			// do nothing
		case "@pbhelp":
			// do nothing
		default:
			s.kickUserFromSquadChat(newUser, &chatRaw)
			return "fail"
		}
	}

	if playerRaw.LeagueID != 1 {
		switch newUser.UserName {
		case "gantz_yaka":
			message := "Здравствуй, " + newUser.UserName + "!\n"
			message += "Инстинкт рад приветствовать Бога мира ПокемемБро! Проходите, располагайтесь, чувствуйте себя, как дома.\n"

			msg := tgbotapi.NewMessage(chatRaw.TelegramID, message)
			msg.ParseMode = "Markdown"

			c.Bot.Send(msg)
		case "@agentpb":
			message := "Здравствуй, " + newUser.UserName + "!\n"
			message += "Инстинкт рад приветствовать одного из богов мира ПокемемБро! Проходите, располагайтесь, чувствуйте себя, как дома.\n"

			msg := tgbotapi.NewMessage(chatRaw.TelegramID, message)
			msg.ParseMode = "Markdown"

			c.Bot.Send(msg)
		case "@pbhelp":
			message := "Здравствуй, " + newUser.UserName + "!\n"
			message += "Инстинкт рад приветствовать одного из богов мира ПокемемБро! Проходите, располагайтесь, чувствуйте себя, как дома.\n"

			msg := tgbotapi.NewMessage(chatRaw.TelegramID, message)
			msg.ParseMode = "Markdown"

			c.Bot.Send(msg)
		default:
			// Check for profile
			_, profileOK := c.Users.GetProfile(playerRaw.ID)
			if !profileOK {
				message := "Привет, " + c.Users.FormatUsername(userName) + "! Напиши мне и скинь профиль для доступа в чаты Лиги!"

				msg := tgbotapi.NewMessage(defaultChatID, message)
				msg.ParseMode = "Markdown"

				c.Bot.Send(msg)
			} else {
				message := "Привет, " + c.Users.FormatUsername(userName) + "! Там переход между лигами не завезли случайно? Переходи в нашу Лигу, будем рады тебя видеть... а пока — вход в наши чаты закрыт!"

				msg := tgbotapi.NewMessage(defaultChatID, message)
				msg.ParseMode = "Markdown"

				c.Bot.Send(msg)
			}
			s.kickUserFromSquadChat(newUser, &chatRaw)
			return "fail"
		}
	}

	return "ok"
}

// FilterBastion kicks already joined user if he changed league
func (s *Squader) FilterBastion(update *tgbotapi.Update) string {
	user := update.Message.From
	chatRaw, ok := c.Chatter.GetOrCreateChat(update)
	if !ok {
		return "fail"
	}

	playerRaw, playerOK := c.Users.GetOrCreatePlayer(update.Message.From.ID)
	if !playerOK {
		s.kickUserFromSquadChat(user, &chatRaw)
		return "fail"
	}
	_, profileOK := c.Users.GetProfile(playerRaw.ID)
	if !profileOK {
		s.kickUserFromSquadChat(user, &chatRaw)
		return "fail"
	}
	if playerRaw.LeagueID != 1 {
		s.kickUserFromSquadChat(user, &chatRaw)
		return "fail"
	}

	return "ok"
}
