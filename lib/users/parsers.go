// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Internal functions

func (u *Users) fillProfilePokememe(profileID int, meme string, attack string, rarity string) {
	spkRaw := dbmapping.Pokememe{}
	err := c.Db.Get(&spkRaw, c.Db.Rebind("SELECT * FROM pokememes WHERE name='"+meme+"';"))
	if err != nil {
		c.Log.Error(err.Error())
	} else {
		attackInt := c.Statistics.GetPoints(attack)
		ppk := dbmapping.ProfilePokememe{}
		ppk.ProfileID = profileID
		ppk.PokememeID = spkRaw.ID
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
	powerInt := 1

	// Filling information
	// We don't know how many strings we got, so we iterating each other
	for i := range profileRunesArray {
		currentString := string(profileRunesArray[i])
		currentRunes := profileRunesArray[i]
		if strings.HasPrefix(currentString, "🈸") || strings.HasPrefix(currentString, "🈳 ") || strings.HasPrefix(currentString, "🈵") {
			err1 := c.Db.Get(&league, c.Db.Rebind("SELECT * FROM leagues WHERE symbol='"+string(currentRunes[0])+"'"))
			if err1 != nil {
				c.Log.Error(err1.Error())
				u.profileAddFailureMessage(update)
				return "fail"
			}
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
			pkmnumRx := regexp.MustCompile(`(\d+)(\d|K|M|)`)
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
	weaponRaw := dbmapping.Weapon{}
	err2 := c.Db.Get(&weaponRaw, c.Db.Rebind("SELECT * FROM weapons WHERE name='"+weapon+"'"))
	if err2 != nil {
		c.Log.Error(err2.Error())
	}

	if playerRaw.LeagueID == 0 {
		// Updating player with league
		playerRaw.LeagueID = league.ID
		if playerRaw.Status == "nobody" {
			playerRaw.Status = "common"
		}
		_, err4 := c.Db.NamedExec("UPDATE `players` SET league_id=:league_id, status=:status WHERE id=:id", &playerRaw)
		if err4 != nil {
			c.Log.Error(err4.Error())
			u.profileAddFailureMessage(update)
			return "fail"
		}
	} else if playerRaw.LeagueID != league.ID {
		// Duplicate profile: user changed league, beware!
		playerRaw.LeagueID = league.ID
		playerRaw.Status = "league_changed"
		playerRaw.CreatedAt = time.Now().UTC()
		_, err5 := c.Db.NamedExec("INSERT INTO players VALUES(NULL, :telegram_id, :league_id, :status, :created_at, :updated_at)", &playerRaw)
		if err5 != nil {
			c.Log.Error(err5.Error())
			u.profileAddFailureMessage(update)
			return "fail"
		}
		err6 := c.Db.Get(&playerRaw, c.Db.Rebind("SELECT * FROM players WHERE telegram_id='"+strconv.Itoa(playerRaw.TelegramID)+"' AND league_id='"+strconv.Itoa(league.ID)+"';"))
		if err6 != nil {
			c.Log.Error(err6.Error())
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
	profileRaw.WeaponID = weaponRaw.ID
	profileRaw.Crystalls = crystallsInt
	profileRaw.CreatedAt = time.Now().UTC()

	_, err3 := c.Db.NamedExec("INSERT INTO `profiles` VALUES(NULL, :player_id, :nickname, :telegram_nickname, :level_id, :pokeballs, :wealth, :pokememes_wealth, :exp, :egg_exp, :power, :weapon_id, :crystalls, :created_at)", &profileRaw)
	if err3 != nil {
		c.Log.Error(err3.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}

	err8 := c.Db.Get(&profileRaw, c.Db.Rebind("SELECT * FROM profiles WHERE player_id=? AND created_at=?"), profileRaw.PlayerID, profileRaw.CreatedAt)
	if err8 != nil {
		c.Log.Error(err8.Error())
		c.Log.Error("Profile isn't added!")
		u.profileAddFailureMessage(update)
		return "fail"
	}

	playerRaw.UpdatedAt = time.Now().UTC()
	_, err7 := c.Db.NamedExec("UPDATE `players` SET updated_at=:updated_at WHERE id=:id", &playerRaw)
	if err7 != nil {
		c.Log.Error(err7.Error())
		u.profileAddFailureMessage(update)
		return "fail"
	}

	for n_meme, attack := range pokememes {
		memeAry := strings.Split(n_meme, "_")
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
		u.fillProfilePokememe(profileRaw.ID, meme, attack, rarity)
	}

	u.profileAddSuccessMessage(update, league.ID, profileRaw.LevelID)
	return "ok"
}
