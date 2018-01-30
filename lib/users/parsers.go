// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Internal functions

func (u *Users) fillProfilePokememe(profileID int, meme string, attack string, rarity string) {
	spkRaw, err := c.DataCache.GetPokememeByName(meme)
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
	profileRunesArray := make([][]rune, 0)
	for i := range profileStringsArray {
		profileRunesArray = append(profileRunesArray, []rune(profileStringsArray[i]))
	}

	league := dbmapping.League{}

	telegramNickname := update.Message.From.UserName
	nickname := ""
	level := ""
	levelInt := 0
	exp := ""
	expInt := 0
	eggexp := ""
	eggexpInt := 0
	pokeballs := ""
	pokeballsInt := 0
	wealth := ""
	wealthInt := 0
	pokememesWealth := ""
	pokememesWealthInt := 0
	crystalls := ""
	crystallsInt := 0
	weapon := ""
	pokememes := make(map[string]string)
	powerInt := 0

	// Filling information
	// We don't know how many strings we got, so we iterating each other
	for i := range profileRunesArray {
		currentString := string(profileRunesArray[i])
		currentRunes := profileRunesArray[i]
		if strings.HasPrefix(currentString, "🈸") || strings.HasPrefix(currentString, "🈳 ") || strings.HasPrefix(currentString, "🈵") {
			leagueRaw, err := c.DataCache.GetLeagueBySymbol(string(currentRunes[0]))
			if err != nil {
				c.Log.Error(err.Error())
				u.profileAddFailureMessage(update)
				return "fail"
			}
			league = *leagueRaw
			for j := range currentRunes {
				if j > 1 {
					nickname += string(currentRunes[j])
				}
			}
		}
		if strings.HasPrefix(currentString, "id: ") {
			realUserID := strings.TrimPrefix(currentString, "id: ")
			c.Log.Debug("Profile user ID: " + realUserID)
			realUID, _ := strconv.Atoi(realUserID)
			if realUID != playerRaw.TelegramID {
				return "fail"
			}
		}
		if strings.HasPrefix(currentString, "👤Уровень:") {
			levelRx := regexp.MustCompile("\\d+")
			levelArray := levelRx.FindAllString(currentString, -1)
			if len(levelArray) < 1 {
				c.Log.Error("Level string broken")
				u.profileAddFailureMessage(update)
				return "fail"
			}
			level = levelArray[0]
			levelInt, _ = strconv.Atoi(level)
		}

		if strings.HasPrefix(currentString, "🎓Опыт") {
			expRx := regexp.MustCompile("\\d+")
			expArray := expRx.FindAllString(currentString, -1)
			if len(expArray) < 4 {
				c.Log.Error("Exp string broken")
				u.profileAddFailureMessage(update)
				return "fail"
			}
			exp = expArray[0]
			expInt, _ = strconv.Atoi(exp)
			eggexp = expArray[2]
			eggexpInt, _ = strconv.Atoi(eggexp)
		}

		if strings.HasPrefix(currentString, "⭕Покеболы") {
			pkbRx := regexp.MustCompile("\\d+")
			pkbArray := pkbRx.FindAllString(currentString, -1)
			if len(pkbArray) < 2 {
				c.Log.Error("Pokeballs string broken")
				u.profileAddFailureMessage(update)
				return "fail"
			}
			pokeballs = pkbArray[1]
			pokeballsInt, _ = strconv.Atoi(pokeballs)
		}

		if strings.HasPrefix(currentString, "💲") {
			wealthRx := regexp.MustCompile("(\\d|\\.|K|M)+")
			wealthArray := wealthRx.FindAllString(currentString, -1)
			if len(wealthArray) < 2 {
				c.Log.Error("Wealth string broken")
				u.profileAddFailureMessage(update)
				return "fail"
			}
			wealth = wealthArray[0]
			wealthInt = c.Statistics.GetPoints(wealth)
			crystalls = wealthArray[1]
			crystallsInt = c.Statistics.GetPoints(crystalls)
		}

		if strings.HasPrefix(currentString, "🔫") {
			// We need NEXT string!
			weaponType := strings.Replace(currentString, "🔫 ", "", 1)
			wnRx := regexp.MustCompile("(.+)(ита|ёры)")
			weapon = wnRx.FindString(weaponType)
		}

		if strings.HasPrefix(currentString, "🐱Покемемы: ") {
			pkmnumRx := regexp.MustCompile(`(\d|\.|K|M)+`)
			pkNumArray := pkmnumRx.FindAllString(currentString, -1)
			if len(pkNumArray) < 3 {
				c.Log.Error("Pokememes count broken")
				u.profileAddFailureMessage(update)
				return "fail"
			}
			pokememesCount, _ := strconv.Atoi(pkNumArray[0])
			pokememesWealth = pkNumArray[2]
			pokememesWealthInt = c.Statistics.GetPoints(pokememesWealth)
			if pokememesCount > 0 {
				for pi := 0; pi < pokememesCount; pi++ {
					pokememeString := string(profileRunesArray[i+1+pi])
					attackRx := regexp.MustCompile("(\\d|\\.|K|M)+")
					pkPointsArray := attackRx.FindAllString(pokememeString, -1)
					pkAttack := pkPointsArray[1]
					pkName := strings.Split(pokememeString, "+")[0]
					pkName = strings.Replace(pkName, " ⭐", "", 1)
					pkName = strings.TrimSuffix(pkName, " ")
					pkName = strings.Split(pkName, "⃣ ")[1]
					pokememes[strconv.Itoa(pi)+"_"+pkName] = pkAttack
					powerInt += c.Statistics.GetPoints(pkAttack)
				}
			}
		}
	}

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
	c.Log.Debug("Crystalls: " + crystalls)
	c.Log.Debugln(crystallsInt)
	c.Log.Debug("Weapon: " + weapon)
	if len(pokememes) > 0 {
		c.Log.Debug("Hand cost: " + pokememesWealth)
		c.Log.Debugln(pokememesWealthInt)
		for meme, attack := range pokememes {
			c.Log.Debug(meme + ": " + attack)
		}
	} else {
		c.Log.Debug("Hand is empty.")
	}

	// Information is gathered, let's create profile in database!
	weaponRaw, err := c.DataCache.GetWeaponTypeByName(weapon)
	if err != nil {
		c.Log.Error(err.Error())
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
	profileRaw.PokememesWealth = pokememesWealthInt
	profileRaw.Exp = expInt
	profileRaw.EggExp = eggexpInt
	profileRaw.Power = powerInt
	if weaponRaw != nil {
		profileRaw.WeaponID = weaponRaw.ID
	} else {
		profileRaw.WeaponID = 0
	}
	profileRaw.Crystalls = crystallsInt
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

	for nMeme, attack := range pokememes {
		memeAry := strings.Split(nMeme, "_")
		meme := memeAry[1]
		rarity := "common"
		if strings.HasPrefix(meme, "🔸") {
			rarity = "rare"
			meme = strings.Replace(meme, "🔸", "", 1)
		}
		if strings.HasPrefix(meme, "🔶") {
			rarity = "super rare"
			meme = strings.Replace(meme, "🔶", "", 1)
		}
		if strings.HasPrefix(meme, "🔹") {
			rarity = "liber"
			meme = strings.Replace(meme, "🔹", "", 1)
		}
		if strings.HasPrefix(meme, "🔷") {
			rarity = "super liber"
			meme = strings.Replace(meme, "🔷", "", 1)
		}
		if strings.HasPrefix(meme, "❄") {
			rarity = "new year"
			meme = strings.Replace(meme, "❄", "", 1)
		}
		u.fillProfilePokememe(newProfileID, meme, attack, rarity)
	}

	u.profileAddSuccessMessage(update, league.ID, profileRaw.LevelID)
	return "ok"
}
