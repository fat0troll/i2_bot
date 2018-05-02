// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package router

import (
	"regexp"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

func (r *Router) routePrivateRequest(update tgbotapi.Update, playerRaw *dbmapping.Player, chatRaw *dbmapping.Chat) string {
	text := update.Message.Text

	// Commands with regexps
	var pokedexMsg = regexp.MustCompile("/pokede(x|ks)\\d?\\z")
	var pokememeInfoMsg = regexp.MustCompile("/pk(\\d+)")
	var usersMsg = regexp.MustCompile("/users(\\d+|)\\z")
	var profileMsg = regexp.MustCompile("/profile(\\d+)\\z")
	var squadInfoMsg = regexp.MustCompile("/show_squad(\\d+)\\z")
	var orderSendMsg = regexp.MustCompile("/send_order(\\d+)\\z")

	if playerRaw.Status == "banned" {
		return c.Talkers.BanError(&update)
	}

	if update.Message.ForwardFrom != nil {
		if update.Message.ForwardFrom.ID == 360402625 {
			c.Log.Info("Forward from PokememBro bot! Processing...")
			if playerRaw.ID != 0 {
				c.Forwarder.ProcessForward(&update, playerRaw)
			} else {
				return c.Talkers.AnyMessageUnauthorized(&update)
			}
		} else if update.Message.ForwardFrom.ID == 392622454 {
			// Pokememes test bot with actual pokedeks
			c.Log.Info("Forward from PokememBro test bot! Processing...")
			if playerRaw.ID != 0 {
				c.Forwarder.ProcessForward(&update, playerRaw)
			} else {
				return c.Talkers.AnyMessageUnauthorized(&update)
			}
		} else {
			c.Log.Info("Forward from another user or bot (" + strconv.Itoa(update.Message.ForwardFrom.ID) + "). Ignoring")
		}
	} else {
		if update.Message.IsCommand() {
			switch {
			case update.Message.Command() == "start":
				if playerRaw.LeagueID != 0 {
					if playerRaw.Status == "special" {
						c.Welcomer.PrivateWelcomeMessageSpecial(&update, playerRaw)
						return "ok"
					}

					c.Welcomer.PrivateWelcomeMessageAuthorized(&update, playerRaw)
					return "ok"
				}

				c.Welcomer.PrivateWelcomeMessageUnauthorized(&update)
				return "ok"

			case update.Message.Command() == "help":
				c.Talkers.HelpMessage(&update, playerRaw)
				return "ok"
			case update.Message.Command() == "faq":
				return c.Talkers.FAQMessage(&update)
			case update.Message.Command() == "academy":
				c.Talkers.AcademyMessage(&update, playerRaw)
				return "ok"
			case update.Message.Command() == "bastion":
				c.Talkers.BastionMessage(&update, playerRaw)
				return "ok"
			case update.Message.Command() == "games_chat":
				c.Talkers.GamesMessage(&update, playerRaw)
				return "ok"

			case pokedexMsg.MatchString(text):
				c.Pokedexer.PokememesList(&update)
				return "ok"
			case pokememeInfoMsg.MatchString(text):
				c.Pokedexer.PokememeInfo(&update, playerRaw)
				return "ok"

			case update.Message.Command() == "me":
				if playerRaw.ID != 0 {
					c.Users.ProfileMessage(&update, playerRaw)
					return "ok"
				}
			case update.Message.Command() == "top":
				if playerRaw.ID != 0 {
					return c.Statistics.TopList(&update, playerRaw)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)
			case update.Message.Command() == "top_my":
				if playerRaw.ID != 0 {
					return c.Statistics.TopList(&update, playerRaw)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)

			case update.Message.Command() == "best":
				c.Pokedexer.AdvicePokememesList(&update, playerRaw)
				return "ok"
			case update.Message.Command() == "advice":
				c.Pokedexer.AdvicePokememesList(&update, playerRaw)
				return "ok"
			case update.Message.Command() == "best_all":
				c.Pokedexer.AdvicePokememesList(&update, playerRaw)
				return "ok"
			case update.Message.Command() == "advice_all":
				c.Pokedexer.AdvicePokememesList(&update, playerRaw)
				return "ok"
			case update.Message.Command() == "best_nofilter":
				c.Pokedexer.AdvicePokememesList(&update, playerRaw)
				return "ok"
			case update.Message.Command() == "reminders":
				return c.Reminder.AlarmsList(&update, playerRaw)

			case update.Message.Command() == "send_all":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					c.Broadcaster.AdminBroadcastMessageCompose(&update, playerRaw)
					return "ok"
				}

				return c.Talkers.AnyMessageUnauthorized(&update)
			case update.Message.Command() == "send_league":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					c.Broadcaster.AdminBroadcastMessageCompose(&update, playerRaw)
					return "ok"
				}

				return c.Talkers.AnyMessageUnauthorized(&update)
			case update.Message.Command() == "send_confirm":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					go c.Broadcaster.AdminBroadcastMessageSend(&update, playerRaw)
					return "ok"
				}

				return c.Talkers.AnyMessageUnauthorized(&update)
			case update.Message.Command() == "chats":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					c.Chatter.GroupsList(&update)
					return "ok"
				}

				return c.Talkers.AnyMessageUnauthorized(&update)
			case update.Message.Command() == "squads":
				return c.Squader.SquadsList(&update, playerRaw)

			case update.Message.Command() == "pin":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Pinner.PinMessageToSomeChats(&update)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)
			case update.Message.Command() == "pin_all":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Pinner.PinMessageToAllChats(&update)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)

			case update.Message.Command() == "orders":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Orders.ListAllOrders(&update)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)
			case orderSendMsg.MatchString(text):
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Orders.SendOrder(&update)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)

			case usersMsg.MatchString(text):
				if c.Users.PlayerBetterThan(playerRaw, "academic") {
					return c.Users.UsersList(&update)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)

			case profileMsg.MatchString(text):
				if c.Users.PlayerBetterThan(playerRaw, "academic") {
					return c.Users.ForeignProfileMessage(&update)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)

			case update.Message.Command() == "find_level":
				if c.Users.PlayerBetterThan(playerRaw, "academic") {
					return c.Users.FindByLevel(&update)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)
			case update.Message.Command() == "find_user":
				if c.Users.PlayerBetterThan(playerRaw, "academic") {
					return c.Users.FindByName(&update)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)
			case update.Message.Command() == "find_top_attack":
				if c.Users.PlayerBetterThan(playerRaw, "academic") {
					return c.Users.FindByTopAttack(&update)
				}

				return c.Talkers.AnyMessageUnauthorized(&update)

			case update.Message.Command() == "squad_add_user":
				return c.Squader.AddUserToSquad(&update, playerRaw)
			case update.Message.Command() == "squad_add_commander":
				return c.Squader.AddUserToSquad(&update, playerRaw)

			case squadInfoMsg.MatchString(text):
				return c.Squader.SquadInfo(&update, playerRaw)
			}
		}
	}

	return "fail"
}
