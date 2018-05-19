// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package users

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Internal functions

func (u *Users) fillProfilePokememe(profileID int, pokememeID int, attack string, rarity string) {
	spkRaw, err := c.DataCache.GetPokememeByID(pokememeID)
	if err != nil {
		c.Log.Error(err.Error())
	} else {
		attackInt := c.Statistics.GetPoints(attack)
		ppk := dbmapping.ProfilePokememe{}
		ppk.ProfileID = profileID
		ppk.PokememeID = spkRaw.Pokememe.ID
		ppk.PokememeAttack = attackInt
		ppk.PokememeRarity = rarity
		ppk.CreatedAt = time.Now().UTC()
		_, err2 := c.Db.NamedExec("INSERT INTO `profiles_pokememes` VALUES(NULL, :profile_id, :pokememe_id, :pokememe_attack, :pokememe_rarity, :created_at)", &ppk)
		if err2 != nil {
			c.Log.Error(err2.Error())
		}
	}
}

// External functions

// ParseProfile parses user profile, forwarded from PokememBroBot, to database
func (u *Users) ParseProfile(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	text := update.Message.Text
	c.Log.Info(text)

	profileStringsArray := strings.Split(text, "\n")

	telegramNickname := update.Message.From.UserName
	rawProfileData := make(map[string]string)
	rawCurrentHandData := make(map[string]string)
	currentHand := 0
	profilePower := 0

	// Not using range here, because range is picking elements in random order
	for i := 0; i < len(profileStringsArray); i++ {
		infoString := profileStringsArray[i]
		c.Log.Debug("Processing string: " + infoString)
		if strings.HasPrefix(infoString, "CurrentHand ") {
			currentHandNumber := strings.Split(infoString, " ")[1]
			currentHandInt, err := strconv.Atoi(currentHandNumber)
			if err != nil {
				c.Log.Error(err.Error())
				u.profileAddFailureMessage(update)
				return "fail"
			}
			currentHand = currentHandInt
		} else if strings.HasPrefix(infoString, ("Hand" + strconv.Itoa(currentHand))) {
			if strings.HasPrefix(infoString, ("Hand"+strconv.Itoa(currentHand))+" Attack") {
				rawProfileData["currentHandAttack"] = strings.Join(strings.Split(infoString, " ")[2:], " ")
			} else {
				rawCurrentHandData[strconv.Itoa(len(rawCurrentHandData))] = strings.Join(strings.Split(infoString, " ")[1:], " ")
			}
		} else {
			rawProfileData[strings.Split(infoString, " ")[0]] = strings.Join(strings.Split(infoString, " ")[1:], " ")
		}
	}

	fmt.Println(rawProfileData)

	nickname := rawProfileData["Name"]
	league, err := c.DataCache.GetLeagueByEnglishName(rawProfileData["Team"])
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}
	level := rawProfileData["Lvl"]
	levelInt, err := strconv.Atoi(level)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}
	exp := rawProfileData["Exp"]
	expInt, err := strconv.Atoi(exp)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}
	eggexp := strings.Split(rawProfileData["Eggs"], "/")[0]
	eggexpInt, err := strconv.Atoi(eggexp)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}
	pokeballs := rawProfileData["BallsMax"]
	pokeballsInt, err := strconv.Atoi(pokeballs)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}
	wealth := rawProfileData["Money"]
	wealthInt, err := strconv.Atoi(wealth)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}
	crystals := rawProfileData["Crystalls"]
	crystalsInt, err := strconv.Atoi(crystals)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}
	weapon := rawProfileData["Weapon"]
	weaponAttack, err := strconv.Atoi(weapon)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}
	currentHandAttack := rawProfileData["currentHandAttack"]
	currentHandAttackInt, err := strconv.Atoi(currentHandAttack)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}
	profilePower = weaponAttack + currentHandAttackInt

	c.Log.Debug("Telegram nickname: " + telegramNickname)
	c.Log.Debug("Nickname: " + nickname)
	c.Log.Debug("League: " + league.Name)
	c.Log.Debug("Level: " + level)
	c.Log.Debugln(levelInt)
	c.Log.Debug("Exp: " + exp)
	c.Log.Debugln(expInt)
	c.Log.Debug("Egg exp: " + eggexp)
	c.Log.Debugln(eggexpInt)
	c.Log.Debug("Pokeballs: " + pokeballs)
	c.Log.Debugln(pokeballsInt)
	c.Log.Debug("Wealth: " + wealth)
	c.Log.Debugln(wealthInt)
	c.Log.Debug("Crystals: " + crystals)
	c.Log.Debugln(crystalsInt)
	c.Log.Debug("Weapon attack: " + weapon)
	c.Log.Debugln(weaponAttack)
	c.Log.Debug("Current hand attack: " + currentHandAttack)
	c.Log.Debugln(currentHandAttackInt)
	if len(rawCurrentHandData) > 0 {
		for i := range rawCurrentHandData {
			c.Log.Debug(rawCurrentHandData[i])
		}
	} else {
		c.Log.Debug("Hand is empty.")
	}

	// Information is gathered, let's create profile in database!
	weaponRaw, err := c.DataCache.GetWeaponTypeByAttack(weaponAttack)
	if err != nil {
		c.Log.Error(err.Error())
	}

	if weaponRaw != nil {
		c.Log.Debug("Got weapon: " + weaponRaw.Name)
	} else {
		c.Log.Debug("This profile contains no weapon")
	}

	if playerRaw.LeagueID == 0 {
		// Updating player with league
		playerRaw.LeagueID = league.ID
		if playerRaw.Status == "nobody" {
			playerRaw.Status = "common"
		}
		_, err = c.DataCache.UpdatePlayerFields(playerRaw)
		if err != nil {
			u.profileAddFailureMessage(update)
			return "fail"
		}
	} else if playerRaw.LeagueID != league.ID {
		// User changed league, beware!
		playerRaw.LeagueID = league.ID
		playerRaw.Status = "league_changed"
		_, err = c.DataCache.UpdatePlayerFields(playerRaw)
		if err != nil {
			u.profileAddFailureMessage(update)
			return "fail"
		}
	}

	profileRaw := dbmapping.Profile{}
	profileRaw.PlayerID = playerRaw.ID
	profileRaw.Nickname = nickname
	profileRaw.TelegramNickname = telegramNickname
	profileRaw.LevelID = levelInt
	profileRaw.Pokeballs = pokeballsInt
	profileRaw.Wealth = wealthInt
	// TODO: count pokememes wealth
	profileRaw.PokememesWealth = 0
	profileRaw.Exp = expInt
	profileRaw.EggExp = eggexpInt
	profileRaw.Power = profilePower
	if weaponRaw != nil {
		profileRaw.WeaponID = weaponRaw.ID
	} else {
		profileRaw.WeaponID = 0
	}
	profileRaw.Crystals = crystalsInt
	profileRaw.CreatedAt = time.Now().UTC()

	newProfileID, err := c.DataCache.AddProfile(&profileRaw)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}

	_, err = c.DataCache.GetProfileByID(newProfileID)
	if err != nil {
		c.Log.Error(err.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}

	err = c.DataCache.UpdatePlayerTimestamp(playerRaw.ID)
	if err != nil {
		u.profileAddFailureMessage(update)
		return "fail"
	}

	for i := range rawCurrentHandData {
		pokememeInfoArray := strings.Split(rawCurrentHandData[i], " ")
		pokememePokedexNumber := pokememeInfoArray[1]
		pokememePokedexID, err := strconv.Atoi(pokememePokedexNumber)
		if err != nil {
			c.Log.Error(err.Error())
		}
		pokememeRarityName := "common"
		pokememeRarity := pokememeInfoArray[3]
		pokememeAttack := pokememeInfoArray[7]
		switch pokememeRarity {
		case "1":
			pokememeRarityName = "rare"
		case "2":
			pokememeRarityName = "super rare"
		case "3":
			pokememeRarityName = "liber"
		case "4":
			pokememeRarityName = "super liber"
		case "7":
			pokememeRarityName = "new year"
		case "8":
			pokememeRarityName = "valentine"
		}
		u.fillProfilePokememe(newProfileID, pokememePokedexID, pokememeAttack, pokememeRarityName)
	}

	u.profileAddSuccessMessage(update, league.ID, profileRaw.LevelID)
	return "ok"
}
