// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package statistics

import (
	"sort"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

func (s *Statistics) renderPosition(profilesRaw *[]*dbmapping.PlayerProfile, playerRaw *dbmapping.Player) string {
	render := "_…а ты в этом топе на "
	profiles := *profilesRaw
	c.Log.Debugln(len(profiles))
	for i := range profiles {
		if profiles[i].Player.ID == playerRaw.ID {
			render += strconv.Itoa(i + 1)
			// Russian numericals...
			if ((i+1)%10) == 3 && ((i+1)%100 == 13) {
				render += "-ем"
			} else {
				render += "-ом"
			}
			render += " месте_\n"
		}
	}

	return render
}

// TopList returns list of top users by level, money ans so on
func (s *Statistics) TopList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	allPlayers := c.DataCache.GetPlayersWithCurrentProfiles()
	myProfile, err := c.DataCache.GetProfileByPlayerID(playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return c.Talkers.AnyMessageUnauthorized(update)
	}

	profiles := make([]*dbmapping.PlayerProfile, 0)

	for i := range allPlayers {
		if allPlayers[i].Player.LeagueID == 1 {
			if update.Message.Command() == "top" {
				profiles = append(profiles, allPlayers[i])
			} else {
				// Local top of level
				if allPlayers[i].Profile.LevelID == myProfile.LevelID {
					profiles = append(profiles, allPlayers[i])
				}
			}
		}
	}

	topLimit := 5
	if len(profiles) < 5 {
		topLimit = len(profiles)
	}

	message := "*Топ-5 по атаке (без биты)*\n"

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Profile.Power > profiles[j].Profile.Power
	})

	for i := 0; i < topLimit; i++ {
		message += "*" + strconv.Itoa(i+1) + "*: " + c.Users.FormatUsername(profiles[i].Profile.Nickname) + " (⚔️" + s.GetPrintablePoints(profiles[i].Profile.Power) + ")\n"
	}

	message += s.renderPosition(&profiles, playerRaw)

	message += "\n*Топ-5 по богатству*\n"

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Profile.Wealth > profiles[j].Profile.Wealth
	})

	for i := 0; i < topLimit; i++ {
		message += "*" + strconv.Itoa(i+1) + "*: " + c.Users.FormatUsername(profiles[i].Profile.Nickname) + " (💲" + s.GetPrintablePoints(profiles[i].Profile.Wealth) + ")\n"
	}

	message += s.renderPosition(&profiles, playerRaw)

	message += "\n*Топ-5 по стоимости покемемов в руке*\n"

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Profile.PokememesWealth > profiles[j].Profile.PokememesWealth
	})

	for i := 0; i < topLimit; i++ {
		message += "*" + strconv.Itoa(i+1) + "*: " + c.Users.FormatUsername(profiles[i].Profile.Nickname) + " (💲" + s.GetPrintablePoints(profiles[i].Profile.PokememesWealth) + ")\n"
	}

	message += s.renderPosition(&profiles, playerRaw)

	message += "\n*Топ-5 по лимиту покеболов*\n"

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Profile.Pokeballs > profiles[j].Profile.Pokeballs
	})

	for i := 0; i < topLimit; i++ {
		message += "*" + strconv.Itoa(i+1) + "*: " + c.Users.FormatUsername(profiles[i].Profile.Nickname) + " (⭕️" + s.GetPrintablePoints(profiles[i].Profile.Pokeballs) + ")\n"
	}

	message += s.renderPosition(&profiles, playerRaw)

	message += "\n*Топ-5 по опыту*\n"

	sort.Slice(profiles, func(i, j int) bool {
		firstProfileLevel, err := c.DataCache.GetLevelByID(profiles[i].Profile.LevelID)
		if err != nil {
			c.Log.Error(err.Error())
			return false
		}
		secondProfileLevel, err := c.DataCache.GetLevelByID(profiles[j].Profile.LevelID)
		if err != nil {
			c.Log.Error(err.Error())
			return false
		}
		firstExp := firstProfileLevel.LevelStart + profiles[i].Profile.Exp
		secondExp := secondProfileLevel.LevelStart + profiles[j].Profile.Exp
		return firstExp > secondExp
	})

	for i := 0; i < topLimit; i++ {
		if profiles[i].Profile.LevelID != 0 {
			profileLevel, err := c.DataCache.GetLevelByID(profiles[i].Profile.LevelID)
			if err != nil {
				c.Log.Error(err.Error())
				return c.Talkers.BotError(update)
			}
			message += "*" + strconv.Itoa(i+1) + "*: " + c.Users.FormatUsername(profiles[i].Profile.Nickname) + " (" + strconv.Itoa(profiles[i].Profile.Exp+profileLevel.LevelStart) + " очков)\n"
		}
	}

	message += s.renderPosition(&profiles, playerRaw)

	message += "\nИгроков, принявших участие в статистике: " + strconv.Itoa(len(profiles))

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
