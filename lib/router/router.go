// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// RouteRequest decides, what to do with user input
func (r *Router) RouteRequest(update *tgbotapi.Update) string {
	playerRaw, ok := c.Users.GetOrCreatePlayer(update.Message.From.ID)
	if !ok {
		// Silently fail
		return "fail"
	}

	chatRaw, ok := c.Chatter.GetOrCreateChat(update)
	if !ok {
		return "fail"
	}

	c.Log.Debug("Received message from chat ")
	c.Log.Debugln(chatRaw.TelegramID)

	if update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup() {
		return r.routeGroupRequest(update, &playerRaw, &chatRaw)
	} else if update.Message.Chat.IsPrivate() {
		return r.routePrivateRequest(update, &playerRaw, &chatRaw)
	}

	return "ok"
}
