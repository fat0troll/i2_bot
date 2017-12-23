// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"math/rand"
	"regexp"
)

func (r *Router) routeGroupRequest(update *tgbotapi.Update, playerRaw *dbmapping.Player, chatRaw *dbmapping.Chat) string {
	text := update.Message.Text
	// Regular expressions
	var durakMsg = regexp.MustCompile("(Д|д)(У|у)(Р|р)(А|а|Е|е|О|о)")
	var huMsg = regexp.MustCompile("(Х|х)(У|у)(Й|й|Я|я|Ю|ю|Е|е)")
	var blMsg = regexp.MustCompile("(\\s|^)(Б|б)(Л|л)((Я|я)(Т|т|Д|д))")
	var ebMsg = regexp.MustCompile("(\\s|^|ЗА|За|зА|за)(Е|е|Ё|ё)(Б|б)(\\s|Л|л|А|а|Т|т|У|у|Е|е|Ё|ё|И|и)")
	var piMsg = regexp.MustCompile("(П|п)(И|и)(З|з)(Д|д)")

	restrictionStatus := c.Chatter.ProtectChat(update, playerRaw, chatRaw)
	if restrictionStatus != "protection_passed" {
		return restrictionStatus
	}

	// Welcomes
	if update.Message.NewChatMembers != nil {
		newUsers := *update.Message.NewChatMembers
		if len(newUsers) > 0 {
			return c.Welcomer.GroupWelcomeMessage(update)
		}
	}
	// New chat names
	if update.Message.NewChatTitle != "" {
		_, ok := c.Chatter.UpdateChatTitle(chatRaw, update.Message.NewChatTitle)
		if ok {
			return "ok"
		}

		return "fail"
	}

	// New chat IDs (usually on supergroup creation)
	if (update.Message.MigrateToChatID != 0) && (update.Message.MigrateFromChatID != 0) {
		_, ok := c.Chatter.UpdateChatTelegramID(update)
		if ok {
			return "ok"
		}

		return "fail"
	}

	// easter eggs
	trigger := rand.Intn(5)
	if trigger == 4 {
		switch {
		case huMsg.MatchString(text):
			return c.Talkers.MatMessage(update)
		case blMsg.MatchString(text):
			return c.Talkers.MatMessage(update)
		case ebMsg.MatchString(text):
			return c.Talkers.MatMessage(update)
		case piMsg.MatchString(text):
			return c.Talkers.MatMessage(update)
		case durakMsg.MatchString(text):
			return c.Talkers.DurakMessage(update)
		}
	}

	switch {
	case update.Message.Command() == "long":
		return c.Talkers.LongMessage(update)
	}

	// Ah, we're still here

	return "ok"
}
