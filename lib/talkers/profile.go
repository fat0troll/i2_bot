// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// stdlib
	"log"
	"strconv"
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../dbmapping"
)

// ProfileMessage shows current player's profile
func (t *Talkers) ProfileMessage(update tgbotapi.Update, playerRaw dbmapping.Player) string {
	profileRaw, ok := c.Getters.GetProfile(playerRaw.ID)
	if !ok {
		c.Talkers.AnyMessageUnauthorized(update)
		return "fail"
	}
	league := dbmapping.League{}
	err := c.Db.Get(&league, c.Db.Rebind("SELECT * FROM leagues WHERE id=?"), playerRaw.LeagueID)
	if err != nil {
		log.Println(err)
	}
	level := dbmapping.Level{}
	err = c.Db.Get(&level, c.Db.Rebind("SELECT * FROM levels WHERE id=?"), profileRaw.LevelID)
	if err != nil {
		log.Println(err)
	}
	weapon := dbmapping.Weapon{}
	if profileRaw.WeaponID != 0 {
		err = c.Db.Get(&weapon, c.Db.Rebind("SELECT * FROM weapons WHERE id=?"), profileRaw.WeaponID)
		if err != nil {
			log.Println(err)
		}
	}
	profilePokememes := []dbmapping.ProfilePokememe{}
	err = c.Db.Select(&profilePokememes, c.Db.Rebind("SELECT * FROM profiles_pokememes WHERE profile_id=?"), profileRaw.ID)
	if err != nil {
		log.Println(err)
	}
	pokememes := []dbmapping.Pokememe{}
	err = c.Db.Select(&pokememes, c.Db.Rebind("SELECT * FROM pokememes"))
	if err != nil {
		log.Println(err)
	}

	attackPokememes := 0
	for i := range profilePokememes {
		for j := range pokememes {
			if profilePokememes[i].PokememeID == pokememes[j].ID {
				singleAttack := profilePokememes[i].PokememeAttack
				attackPokememes += singleAttack
			}
		}
	}

	message := "*Профиль игрока "
	message += profileRaw.Nickname + "* (@" + profileRaw.TelegramNickname + ")\n"
	message += "\nЛига: " + league.Symbol + league.Name
	message += "\n👤 " + strconv.Itoa(profileRaw.LevelID)
	message += " | 🎓 " + strconv.Itoa(profileRaw.Exp) + "/" + strconv.Itoa(level.MaxExp)
	message += " | 🥚 " + strconv.Itoa(profileRaw.EggExp) + "/" + strconv.Itoa(level.MaxEgg)
	message += "\n💲" + c.Parsers.ReturnPoints(profileRaw.Wealth)
	message += " |💎" + strconv.Itoa(profileRaw.Crystalls)
	message += " |⭕" + strconv.Itoa(profileRaw.Pokeballs)
	message += "\n⚔Атака: 1 + " + c.Parsers.ReturnPoints(weapon.Power) + " + " + c.Parsers.ReturnPoints(attackPokememes) + "\n"

	if profileRaw.WeaponID != 0 {
		message += "\n🔫Оружие: " + weapon.Name + " " + c.Parsers.ReturnPoints(weapon.Power) + "⚔"
	}

	message += "\n🐱Покемемы:"
	for i := range profilePokememes {
		for j := range pokememes {
			if profilePokememes[i].PokememeID == pokememes[j].ID {
				message += "\n" + strconv.Itoa(pokememes[j].Grade)
				message += "⃣ " + pokememes[j].Name
				message += " +" + c.Parsers.ReturnPoints(profilePokememes[i].PokememeAttack) + "⚔"
			}
		}
	}
	message += "\nСтоимость покемемов на руках: " + c.Parsers.ReturnPoints(profileRaw.PokememesWealth) + "$"
	message += "\n\n💳" + strconv.Itoa(playerRaw.TelegramID)
	message += "\n⏰Последнее обновление профиля: " + profileRaw.CreatedAt.Format("02.01.2006 15:04:05")
	message += "\n\nНе забывай обновляться, это важно для получения актуальной информации.\n\n"
	message += "/best – посмотреть лучших покемемов для поимки"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
