// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package chatter

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
)

func (ct *Chatter) userPrivilegesCheck(update *tgbotapi.Update, user *tgbotapi.User) bool {
	// There are two special chats, pointed by config, where any member of league may be
	defaultChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.DefaultID, 10, 64)
	bastionChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
	academyChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.AcademyID, 10, 64)
	hqChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)

	if update.Message.Chat.ID == defaultChatID || update.Message.Chat.ID == hqChatID {
		return true
	}

	// There are special users, which will bypass these checks
	specialUsers := []string{"gantz_yaka", "agentpb", "pbhelp"}

	for j := range specialUsers {
		if strings.ToLower(user.UserName) == specialUsers[j] {
			// This is for PokememBro admins, they can join any chat at any time
			return true
		}
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
	case academyChatID:
		if playerRaw.LeagueID == 1 && playerRaw.Status != "spy" && playerRaw.Status != "league_changed" && playerRaw.Status != "banned" {
			return true
		}
	case bastionChatID:
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
			for i := range newUsers {
				newUserPassed := ct.userPrivilegesCheck(update, &newUsers[i])
				if !newUserPassed {
					ct.BanUserFromChat(&newUsers[i], chatRaw)
				}
			}
		}
	}

	existingUserPassed := ct.userPrivilegesCheck(update, update.Message.From)
	if !existingUserPassed {
		ct.BanUserFromChat(update.Message.From, chatRaw)
		return "fail"
	}

	return "ok"
}
