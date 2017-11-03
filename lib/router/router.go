// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	// stdlib
	"log"
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Router is a function-handling struct for router
type Router struct{}

// RouteRequest decides, what to do with user input
func (r *Router) RouteRequest(update tgbotapi.Update) string {
	playerRaw, ok := c.Getters.GetOrCreatePlayer(update.Message.From.ID)
	if !ok {
		// Silently fail
		return "fail"
	}

	chatRaw, ok := c.Getters.GetOrCreateChat(&update)
	if !ok {
		return "fail"
	}

	log.Printf("Received message from chat ")
	log.Println(chatRaw.TelegramID)

	if update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup() {
		return r.routeGroupRequest(update, playerRaw, chatRaw)
	} else if update.Message.Chat.IsPrivate() {
		return r.routePrivateRequest(update, playerRaw, chatRaw)
	}

	return "ok"
}
