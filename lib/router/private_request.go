// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"regexp"
)

func (r *Router) routePrivateRequest(update *tgbotapi.Update, playerRaw *dbmapping.Player, chatRaw *dbmapping.Chat) string {
	text := update.Message.Text

	// Commands with regexps
	var pokedexMsg = regexp.MustCompile("/pokede(x|ks)\\d?\\z")
	var pokememeInfoMsg = regexp.MustCompile("/pk(\\d+)")

	if update.Message.ForwardFrom != nil {
		if update.Message.ForwardFrom.ID != 360402625 {
			c.Log.Info("Forward from another user or bot. Ignoring")
		} else {
			c.Log.Info("Forward from PokememBro bot! Processing...")
			if playerRaw.ID != 0 {
				c.Forwarder.ProcessForward(update, playerRaw)
			} else {
				c.Talkers.AnyMessageUnauthorized(update)
				return "fail"
			}
		}
	} else {
		if update.Message.IsCommand() {
			switch {
			case update.Message.Command() == "start":
				if playerRaw.ID != 0 {
					c.Talkers.HelloMessageAuthorized(update, playerRaw)
					return "ok"
				}

				c.Talkers.HelloMessageUnauthorized(update)
				return "ok"
			case update.Message.Command() == "help":
				c.Talkers.HelpMessage(update, playerRaw)
				return "ok"
			// Pokememes info
			case pokedexMsg.MatchString(text):
				c.Talkers.PokememesList(update)
				return "ok"
			case pokememeInfoMsg.MatchString(text):
				c.Talkers.PokememeInfo(update, playerRaw)
				return "ok"
			case update.Message.Command() == "me":
				if playerRaw.ID != 0 {
					c.Talkers.ProfileMessage(update, playerRaw)
					return "ok"
				}

				c.Talkers.AnyMessageUnauthorized(update)
				return "fail"
			case update.Message.Command() == "best":
				c.Talkers.BestPokememesList(update, playerRaw)
				return "ok"
			case update.Message.Command() == "send_all":
				if c.Getters.PlayerBetterThan(playerRaw, "admin") {
					c.Talkers.AdminBroadcastMessageCompose(update, playerRaw)
					return "ok"
				}

				c.Talkers.AnyMessageUnauthorized(update)
				return "fail"
			case update.Message.Command() == "send_confirm":
				if c.Getters.PlayerBetterThan(playerRaw, "admin") {
					c.Talkers.AdminBroadcastMessageSend(update, playerRaw)
					return "ok"
				}

				c.Talkers.AnyMessageUnauthorized(update)
				return "fail"
			case update.Message.Command() == "group_chats":
				if c.Getters.PlayerBetterThan(playerRaw, "admin") {
					c.Talkers.GroupsList(update)
					return "ok"
				}

				c.Talkers.AnyMessageUnauthorized(update)
				return "fail"
			case update.Message.Command() == "pin":
				if c.Getters.PlayerBetterThan(playerRaw, "admin") {
					return c.Pinner.PinMessageToAllChats(update)
				}

				c.Talkers.AnyMessageUnauthorized(update)
				return "fail"
			}
		}
	}

	return "fail"
}
