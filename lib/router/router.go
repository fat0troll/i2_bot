// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package router

import (
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// RouteRequest decides, what to do with user input
func (r *Router) RouteRequest(update tgbotapi.Update) string {
	c.Log.Debugln(update)
	playerRaw, err := c.DataCache.GetOrCreatePlayerByTelegramID(update.Message.From.ID)
	if err != nil {
		c.Log.Error(err.Error())
		// Silently fail
		return "fail"
	}

	c.Log.Debug("Getting chat...")
	chatRaw, err := c.DataCache.GetOrCreateChat(&update)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	if update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup() {
		c.Log.Debug("Routing group request in group with ID = " + strconv.Itoa(int(chatRaw.TelegramID)))
		return r.routeGroupRequest(update, playerRaw, chatRaw)
	} else if update.Message.Chat.IsPrivate() {
		c.Log.Debug("Routing private request for user with ID = " + strconv.Itoa(int(chatRaw.TelegramID)))
		return r.routePrivateRequest(update, playerRaw, chatRaw)
	}

	return "ok"
}
