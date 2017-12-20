// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squader

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
)

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

	bastionChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
	hqChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.HeadquartersID, 10, 64)
	if chatRaw.TelegramID != bastionChatID {
		// In Bastion notifications are public in default chat
		commanders, ok := s.getCommandersForSquadViaChat(chatRaw)
		if ok {
			for i := range commanders {
				message := "Некто " + c.Users.GetPrettyName(user) + " попытался зайти в чат _" + chatRaw.Name + "_ и был изгнан ботом, так как не имеет права посещать этот чат."

				msg := tgbotapi.NewMessage(int64(commanders[i].TelegramID), message)
				msg.ParseMode = "Markdown"
				c.Bot.Send(msg)
			}
		}
	} else {
		message := "Некто " + c.Users.GetPrettyName(user) + " попытался зайти в чат _Бастион Инстинкта_ и был изгнан ботом, так как не имеет права посещать этот чат."

		msg := tgbotapi.NewMessage(hqChatID, message)
		msg.ParseMode = "Markdown"
		c.Bot.Send(msg)
	}
}
