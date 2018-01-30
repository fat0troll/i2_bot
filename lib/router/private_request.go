// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"regexp"
)

func (r *Router) routePrivateRequest(update *tgbotapi.Update, playerRaw *dbmapping.Player, chatRaw *dbmapping.Chat) string {
	text := update.Message.Text

	// Commands with regexps
	var pokedexMsg = regexp.MustCompile("/pokede(x|ks)\\d?\\z")
	var pokememeInfoMsg = regexp.MustCompile("/pk(\\d+)")
	var usersMsg = regexp.MustCompile("/users(\\d+|)\\z")
	var profileMsg = regexp.MustCompile("/profile(\\d+)\\z")
	var squadInfoMsg = regexp.MustCompile("/show_squad(\\d+)\\z")
	var orderSendMsg = regexp.MustCompile("/send_order(\\d+)\\z")

	if update.Message.ForwardFrom != nil {
		if update.Message.ForwardFrom.ID != 360402625 {
			c.Log.Info("Forward from another user or bot. Ignoring")
		} else {
			c.Log.Info("Forward from PokememBro bot! Processing...")
			if playerRaw.ID != 0 {
				c.Forwarder.ProcessForward(update, playerRaw)
			} else {
				return c.Talkers.AnyMessageUnauthorized(update)
			}
		}
	} else {
		if update.Message.IsCommand() {
			switch {
			case update.Message.Command() == "start":
				if playerRaw.LeagueID != 0 {
					if playerRaw.Status == "special" {
						c.Welcomer.PrivateWelcomeMessageSpecial(update, playerRaw)
						return "ok"
					}

					c.Welcomer.PrivateWelcomeMessageAuthorized(update, playerRaw)
					return "ok"
				}

				c.Welcomer.PrivateWelcomeMessageUnauthorized(update)
				return "ok"

			case update.Message.Command() == "help":
				c.Talkers.HelpMessage(update, playerRaw)
				return "ok"
			case update.Message.Command() == "academy":
				c.Talkers.AcademyMessage(update, playerRaw)
				return "ok"
			case update.Message.Command() == "bastion":
				c.Talkers.BastionMessage(update, playerRaw)
				return "ok"

			case pokedexMsg.MatchString(text):
				c.Pokedexer.PokememesList(update)
				return "ok"
			case pokememeInfoMsg.MatchString(text):
				c.Pokedexer.PokememeInfo(update, playerRaw)
				return "ok"
			case update.Message.Command() == "delete_pokememe":
				if c.Users.PlayerBetterThan(playerRaw, "owner") {
					return c.Pokedexer.DeletePokememe(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case update.Message.Command() == "me":
				if playerRaw.ID != 0 {
					c.Users.ProfileMessage(update, playerRaw)
					return "ok"
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case update.Message.Command() == "best":
				c.Pokedexer.BestPokememesList(update, playerRaw)
				return "ok"
			case update.Message.Command() == "reminders":
				return c.Reminder.AlarmsList(update, playerRaw)

			case update.Message.Command() == "send_all":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					c.Broadcaster.AdminBroadcastMessageCompose(update, playerRaw)
					return "ok"
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case update.Message.Command() == "send_league":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					c.Broadcaster.AdminBroadcastMessageCompose(update, playerRaw)
					return "ok"
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case update.Message.Command() == "send_confirm":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					go c.Broadcaster.AdminBroadcastMessageSend(update, playerRaw)
					return "ok"
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case update.Message.Command() == "chats":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					c.Chatter.GroupsList(update)
					return "ok"
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case update.Message.Command() == "squads":
				return c.Squader.SquadsList(update, playerRaw)
			case update.Message.Command() == "make_squad":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Squader.CreateSquad(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case update.Message.Command() == "pin":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Pinner.PinMessageToSomeChats(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case update.Message.Command() == "pin_all":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Pinner.PinMessageToAllChats(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)

			case update.Message.Command() == "orders":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Orders.ListAllOrders(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case orderSendMsg.MatchString(text):
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Orders.SendOrder(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)

			case usersMsg.MatchString(text):
				if c.Users.PlayerBetterThan(playerRaw, "academic") {
					return c.Users.UsersList(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)

			case profileMsg.MatchString(text):
				if c.Users.PlayerBetterThan(playerRaw, "academic") {
					return c.Users.ForeignProfileMessage(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)

			case update.Message.Command() == "find_level":
				if c.Users.PlayerBetterThan(playerRaw, "academic") {
					return c.Users.FindByLevel(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			case update.Message.Command() == "find_user":
				if c.Users.PlayerBetterThan(playerRaw, "academic") {
					return c.Users.FindByName(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)

			case update.Message.Command() == "squad_add_user":
				return c.Squader.AddUserToSquad(update, playerRaw)
			case update.Message.Command() == "squad_add_commander":
				return c.Squader.AddUserToSquad(update, playerRaw)

			case squadInfoMsg.MatchString(text):
				return c.Squader.SquadInfo(update, playerRaw)

			case update.Message.Command() == "five_offer":
				if c.Users.PlayerBetterThan(playerRaw, "admin") {
					return c.Talkers.FiveOffer(update)
				}

				return c.Talkers.AnyMessageUnauthorized(update)
			}
		}
	}

	return "fail"
}
