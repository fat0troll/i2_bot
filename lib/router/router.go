// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// RouteRequest decides, what to do with user input
func (r *Router) RouteRequest(update *tgbotapi.Update) string {
	playerRaw, err := c.DataCache.GetOrCreatePlayerByTelegramID(update.Message.From.ID)
	if err != nil {
		c.Log.Error(err.Error())
		// Silently fail
		return "fail"
	}

	chatRaw, err := c.DataCache.GetOrCreateChat(update)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	if update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup() {
		return r.routeGroupRequest(update, playerRaw, chatRaw)
	} else if update.Message.Chat.IsPrivate() {
		return r.routePrivateRequest(update, playerRaw, chatRaw)
	}

	return "ok"
}
