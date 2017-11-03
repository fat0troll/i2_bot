// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	// stdlib
	"log"
	"regexp"
	"strings"
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

func (r *Router) routePrivateRequest(update tgbotapi.Update, playerRaw dbmapping.Player, chatRaw dbmapping.Chat) string {
	text := update.Message.Text
	// Forwards
	var pokememeMsg = regexp.MustCompile("(Уровень)(.+)(Опыт)(.+)\n(Элементы:)(.+)\n(.+)(💙MP)")
	var profileMsg = regexp.MustCompile(`(Онлайн: )(\d+)\n(Турнир через)(.+)\n\n((.*)\n|(.*)\n(.*)\n)(Элементы)(.+)\n(.*)\n\n(.+)(Уровень)(.+)\n`)

	// Commands with regexps
	var pokedexMsg = regexp.MustCompile("/pokede(x|ks)\\d?\\z")
	var pokememeInfoMsg = regexp.MustCompile("/pk(\\d+)")

	if update.Message.ForwardFrom != nil {
		if update.Message.ForwardFrom.ID != 360402625 {
			log.Printf("Forward from another user or bot. Ignoring")
		} else {
			log.Printf("Forward from PokememBro bot! Processing...")
			if playerRaw.ID != 0 {
				switch {
				case pokememeMsg.MatchString(text):
					log.Printf("Pokememe posted!")
					if playerRaw.LeagueID == 1 {
						status := c.Parsers.ParsePokememe(text, playerRaw)
						switch status {
						case "ok":
							c.Talkers.PokememeAddSuccessMessage(update)
							return "ok"
						case "dup":
							c.Talkers.PokememeAddDuplicateMessage(update)
							return "ok"
						case "fail":
							c.Talkers.PokememeAddFailureMessage(update)
							return "fail"
						}
					} else {
						c.Talkers.AnyMessageUnauthorized(update)
						return "fail"
					}
				case profileMsg.MatchString(text):
					log.Printf("Profile posted!")
					status := c.Parsers.ParseProfile(update, playerRaw)
					switch status {
					case "ok":
						c.Talkers.ProfileAddSuccessMessage(update)
						return "ok"
					case "fail":
						c.Talkers.ProfileAddFailureMessage(update)
						return "fail"
					}
				default:
					log.Printf(text)
					return "fail"
				}
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
				c.Talkers.HelpMessage(update, &playerRaw)
				return "ok"
			// Pokememes info
			case pokedexMsg.MatchString(text):
				if strings.HasSuffix(text, "1") {
					c.Talkers.PokememesList(update, 1)
					return "ok"
				} else if strings.HasSuffix(text, "2") {
					c.Talkers.PokememesList(update, 2)
					return "ok"
				} else if strings.HasSuffix(text, "3") {
					c.Talkers.PokememesList(update, 3)
					return "ok"
				} else if strings.HasSuffix(text, "4") {
					c.Talkers.PokememesList(update, 4)
					return "ok"
				} else if strings.HasSuffix(text, "5") {
					c.Talkers.PokememesList(update, 5)
					return "ok"
				}

				c.Talkers.PokememesList(update, 1)
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
				if c.Getters.PlayerBetterThan(&playerRaw, "admin") {
					c.Talkers.AdminBroadcastMessageCompose(update, &playerRaw)
					return "ok"
				}

				c.Talkers.AnyMessageUnauthorized(update)
				return "fail"
			case update.Message.Command() == "send_confirm":
				if c.Getters.PlayerBetterThan(&playerRaw, "admin") {
					c.Talkers.AdminBroadcastMessageSend(update, &playerRaw)
					return "ok"
				}

				c.Talkers.AnyMessageUnauthorized(update)
				return "fail"
			}
		}
	}

	return "fail"
}
