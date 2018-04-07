// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package chatter

import (
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

func (ct *Chatter) userPrivilegesCheck(update *tgbotapi.Update, user *tgbotapi.User) bool {
	// There are two special chats, pointed by config, where any member of league may be
	defaultChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.DefaultID, 10, 64)
	bastionChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
	academyChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.AcademyID, 10, 64)
	hqChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
	gamesChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.GamesID, 10, 64)

	if update.Message.Chat.ID == defaultChatID || update.Message.Chat.ID == hqChatID {
		return true
	}

	// There are special users, which will bypass these checks
	specialUsers := []string{"gantz_yaka", "agentpb", "pbhelp", "i2_bot", "i2_dev_bot"}

	for j := range specialUsers {
		if strings.ToLower(user.UserName) == specialUsers[j] {
			// This is for PokememBro admins, they can join any chat at any time
			return true
		}
	}

	switch update.Message.Chat.ID {
	case academyChatID:
		c.Log.Debug("Checking user rights in academy chat...")
	case bastionChatID:
		c.Log.Debug("Checking user rights in bastion chat...")
	case gamesChatID:
		c.Log.Debug("Checking user rights in games chat...")
	case hqChatID:
		c.Log.Debug("Checking user rights in headquarters chat...")
	}

	if update.Message.Chat.ID == gamesChatID && strings.Contains(user.UserName, "bot") {
		c.Log.Debug("Game bot with username @" + update.Message.From.UserName + " passed filtration")
		return true
	}

	playerRaw, err := c.DataCache.GetPlayerByTelegramID(user.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return false
	}

	if c.Users.PlayerBetterThan(playerRaw, "admin") {
		return true
	}

	// So, user is not a PokememBro admin. For Bastion and Academy she needs to be league player
	switch update.Message.Chat.ID {
	case academyChatID, bastionChatID, gamesChatID:
		if playerRaw.LeagueID == 1 && playerRaw.Status != "spy" && playerRaw.Status != "league_changed" && playerRaw.Status != "banned" {
			return true
		}
	default:
		availableChatsForUser := c.DataCache.GetAvailableSquadsChatsForUser(playerRaw.ID)
		for i := range availableChatsForUser {
			if update.Message.Chat.ID == availableChatsForUser[i].TelegramID {
				return true
			}
		}
	}

	c.Log.Debug("User failed to prove identity. Ban sequence arrived.")

	return false
}

// BanUserFromChat removes user from chat
func (ct *Chatter) BanUserFromChat(user *tgbotapi.User, chatRaw *dbmapping.Chat) {
	chatUserConfig := tgbotapi.ChatMemberConfig{
		ChatID: chatRaw.TelegramID,
		UserID: user.ID,
	}

	kickConfig := tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: chatUserConfig,
		UntilDate:        1893456000,
	}

	c.Log.Info("Trying to ban user...")

	_, err := c.Bot.KickChatMember(kickConfig)
	if err != nil {
		c.Log.Error(err.Error())
	}

	bastionChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
	academyChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.AcademyID, 10, 64)
	hqChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
	if (chatRaw.TelegramID != bastionChatID) || (chatRaw.TelegramID != academyChatID) {
		squad, err := c.DataCache.GetSquadByChatID(chatRaw.ID)
		if err != nil {
			c.Log.Error(err.Error())
		} else {
			// In Bastion notifications are public in default chat
			commanders := c.DataCache.GetCommandersForSquad(squad.ID)
			for i := range commanders {
				message := "Некто " + c.Users.GetPrettyName(user) + " попытался зайти в чат _" + chatRaw.Name + "_ и был изгнан ботом, так как не имеет права посещать этот чат."

				msg := tgbotapi.NewMessage(int64(commanders[i].TelegramID), message)
				msg.ParseMode = "Markdown"
				c.Bot.Send(msg)
			}
		}
	} else {
		message := "Некто " + c.Users.GetPrettyName(user) + " попытался зайти в один из общих чатов лиги и был изгнан ботом, так как не имеет права посещать этот чат."

		msg := tgbotapi.NewMessage(hqChatID, message)
		msg.ParseMode = "Markdown"
		c.Bot.Send(msg)
	}
}

// ProtectChat protects chats from unauthorized access
// Returns "protection_passed" if all protection checks passed
func (ct *Chatter) ProtectChat(update *tgbotapi.Update, playerRaw *dbmapping.Player, chatRaw *dbmapping.Chat) string {
	// Check on new user addition
	if update.Message.NewChatMembers != nil {
		newUsers := *update.Message.NewChatMembers
		if len(newUsers) > 0 {
			c.Log.Debug("New users joined/added to chat. Checking rights for them.")
			for i := range newUsers {
				newUserPassed := ct.userPrivilegesCheck(update, &newUsers[i])
				if !newUserPassed {
					c.Log.Debug("This user can't be here: removing from chat...")
					ct.BanUserFromChat(&newUsers[i], chatRaw)
				}
			}
		}
	}

	existingUserPassed := ct.userPrivilegesCheck(update, update.Message.From)
	if !existingUserPassed {
		c.Log.Debug("Existing chat user can't be here. Vanishing...")
		ct.BanUserFromChat(update.Message.From, chatRaw)
		return "fail"
	}

	return "ok"
}
