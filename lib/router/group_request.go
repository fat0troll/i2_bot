// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"regexp"
)

func (r *Router) routeGroupRequest(update *tgbotapi.Update, playerRaw *dbmapping.Player, chatRaw *dbmapping.Chat) string {
	text := update.Message.Text
	// Regular expressions
	var durakMsg = regexp.MustCompile("(Д|д)(У|у)(Р|р)(А|а|Е|е|О|о)")
	var huMsg = regexp.MustCompile("(Х|х)(У|у)(Й|й|Я|я|Ю|ю|Е|е)")
	var blMsg = regexp.MustCompile("(\\s|^)(Б|б)(Л|л)(Я|я)(Т|т|Д|д)")
	var ebMsg = regexp.MustCompile("(\\s|^|ЗА|За|зА|за)(Е|е|Ё|ё)(Б|б)(\\s|Л|л|А|а|Т|т|У|у|Е|е|Ё|ё|И|и)")
	var piMsg = regexp.MustCompile("(П|п)(И|и)(З|з)(Д|д)")

	// Welcomes
	if update.Message.NewChatMembers != nil {
		newUsers := *update.Message.NewChatMembers
		if len(newUsers) > 0 {
			return c.Welcomer.WelcomeMessage(update)
		}
	}
	// New chat names
	if update.Message.NewChatTitle != "" {
		_, ok := c.Getters.UpdateChatTitle(chatRaw, update.Message.NewChatTitle)
		if ok {
			return "ok"
		}

		return "fail"
	}

	switch {
	case huMsg.MatchString(text):
		c.Talkers.MatMessage(update)
	case blMsg.MatchString(text):
		c.Talkers.MatMessage(update)
	case ebMsg.MatchString(text):
		c.Talkers.MatMessage(update)
	case piMsg.MatchString(text):
		c.Talkers.MatMessage(update)
	case durakMsg.MatchString(text):
		c.Talkers.DurakMessage(update)
	}

	return "ok"
}
