// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"math/rand"
	"regexp"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

func (r *Router) routeGroupRequest(update tgbotapi.Update, playerRaw *dbmapping.Player, chatRaw *dbmapping.Chat) string {
	text := update.Message.Text
	// Regular expressions
	var durakMsg = regexp.MustCompile("(Д|д)(У|у)(Р|р)(А|а|Е|е|О|о)")
	var huMsg = regexp.MustCompile("(Х|х)(У|у)(Й|й|Я|я|Ю|ю|Е|е)")
	var blMsg = regexp.MustCompile("(\\s|^)(Б|б)(Л|л)((Я|я)(Т|т|Д|д))")
	var ebMsg = regexp.MustCompile("(\\s|^|ЗА|За|зА|за)(Е|е|Ё|ё)(Б|б)(\\s|Л|л|А|а|Т|т|У|у|Е|е|Ё|ё|И|и)")
	var piMsg = regexp.MustCompile("(П|п)(И|и)(З|з)(Д|д)")

	restrictionStatus := c.Chatter.ProtectChat(&update, playerRaw, chatRaw)
	if restrictionStatus != "ok" {
		return restrictionStatus
	}

	// Welcomes
	if update.Message.NewChatMembers != nil {
		newUsers := *update.Message.NewChatMembers
		if len(newUsers) > 0 {
			return c.Welcomer.GroupWelcomeMessage(&update)
		}
	}
	// New chat names
	if update.Message.NewChatTitle != "" {
		_, err := c.DataCache.UpdateChatTitle(chatRaw.ID, update.Message.NewChatTitle)
		if err != nil {
			c.Log.Error(err.Error())
			return "fail"
		}

		return "ok"
	}

	// easter eggs
	trigger := rand.Intn(5)
	if trigger == 4 {
		switch {
		case huMsg.MatchString(text):
			return c.Talkers.MatMessage(&update)
		case blMsg.MatchString(text):
			return c.Talkers.MatMessage(&update)
		case ebMsg.MatchString(text):
			return c.Talkers.MatMessage(&update)
		case piMsg.MatchString(text):
			return c.Talkers.MatMessage(&update)
		case durakMsg.MatchString(text):
			return c.Talkers.DurakMessage(&update)
		}
	}

	switch {
	case update.Message.Command() == "long":
		return c.Talkers.LongMessage(&update)
	default:
		if update.Message.IsCommand() {
			return c.Talkers.RulesMessage(&update)
		}
	}

	// Ah, we're still here

	return "ok"
}
