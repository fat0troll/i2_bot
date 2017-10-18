// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsers

import (
	// stdlib
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
	// 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../dbmapping"
)

// Internal functions

func (p *Parsers) fillProfilePokememe(profileID int, meme string, attack string, rarity string) {
	spkRaw := dbmapping.Pokememe{}
	err := c.Db.Get(&spkRaw, c.Db.Rebind("SELECT * FROM pokememes WHERE name='"+meme+"';"))
	if err != nil {
		log.Println(err)
	} else {
		attackInt := p.getPoints(attack)
		// Improve it. Game's precision is unstable
		origAttack := float64(spkRaw.Attack)
		if rarity == "rare" {
			origAttack = origAttack * 1.1
		}
		level := int(float64(attackInt) / origAttack)

		ppk := dbmapping.ProfilePokememe{}
		ppk.ProfileID = profileID
		ppk.PokememeID = spkRaw.ID
		ppk.PokememeLevel = level
		ppk.PokememeRarity = rarity
		ppk.CreatedAt = time.Now().UTC()
		_, err2 := c.Db.NamedExec("INSERT INTO `profiles_pokememes` VALUES(NULL, :profile_id, :pokememe_id, :pokememe_lvl, :pokememe_rarity, :created_at)", &ppk)
		if err2 != nil {
			log.Println(err2)
		}
	}
}

// External functions

// ParseProfile parses user profile, forwarded from PokememBroBot, to database
func (p *Parsers) ParseProfile(update tgbotapi.Update, playerRaw dbmapping.Player) string {
	text := update.Message.Text
	log.Println(text)

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
				log.Println(err1)
				return "fail"
			}
			for j := range currentRunes {
				if j > 1 {
					nickname += string(currentRunes[j])
				}
			}
		}
		if strings.HasPrefix(currentString, "👤Уровень:") {
			levelRx := regexp.MustCompile("\\d+")
			levelArray := levelRx.FindAllString(currentString, -1)
			if len(levelArray) < 1 {
				log.Println("Level string broken")
				return "fail"
			}
			level = levelArray[0]
			levelInt, _ = strconv.Atoi(level)
		}

		if strings.HasPrefix(currentString, "🎓Опыт") {
			expRx := regexp.MustCompile("\\d+")
			expArray := expRx.FindAllString(currentString, -1)
			if len(expArray) < 4 {
				log.Println("Exp string broken")
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
				log.Println("Pokeballs string broken")
				return "fail"
			}
			pokeballs = pkbArray[1]
			pokeballsInt, _ = strconv.Atoi(pokeballs)
		}

		if strings.HasPrefix(currentString, "💲") {
			wealthRx := regexp.MustCompile("(\\d|\\.|K|M)+")
			wealthArray := wealthRx.FindAllString(currentString, -1)
			if len(wealthArray) < 2 {
				log.Println("Wealth string broken")
				return "fail"
			}
			wealth = wealthArray[0]
			wealthInt = p.getPoints(wealth)
			crystalls = wealthArray[1]
			crystallsInt = p.getPoints(crystalls)
		}

		if strings.HasPrefix(currentString, "🔫") {
			// We need NEXT string!
			weaponType := strings.Replace(currentString, "🔫 ", "", 1)
			wnRx := regexp.MustCompile("(.+)(ита)")
			weapon = wnRx.FindString(weaponType)
		}

		if strings.HasPrefix(currentString, "🐱Покемемы: ") {
			pkmnumRx := regexp.MustCompile("\\d+")
			pkNumArray := pkmnumRx.FindAllString(currentString, -1)
			if len(pkNumArray) < 2 {
				log.Println("Pokememes count broken")
				return "fail"
			}
			pokememesCount, _ := strconv.Atoi(pkNumArray[0])
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
					pokememes[pkName] = pkAttack
					powerInt += p.getPoints(pkAttack)
				}
			}
		}
	}

	log.Printf("Telegram nickname: " + telegramNickname)
	log.Printf("Nickname: " + nickname)
	log.Printf("League: " + league.Name)
	log.Printf("Level: " + level)
	log.Println(levelInt)
	log.Printf("Exp: " + exp)
	log.Println(expInt)
	log.Printf("Egg exp: " + eggexp)
	log.Println(eggexpInt)
	log.Printf("Pokeballs: " + pokeballs)
	log.Println(pokeballsInt)
	log.Printf("Wealth: " + wealth)
	log.Println(wealthInt)
	log.Printf("Crystalls: " + crystalls)
	log.Println(crystallsInt)
	log.Printf("Weapon: " + weapon)
	if len(pokememes) > 0 {
		for meme, attack := range pokememes {
			log.Printf(meme + ": " + attack)
		}
	} else {
		log.Printf("Hand is empty.")
	}

	// Information is gathered, let's create profile in database!
	weaponRaw := dbmapping.Weapon{}
	err2 := c.Db.Get(&weaponRaw, c.Db.Rebind("SELECT * FROM weapons WHERE name='"+weapon+"'"))
	if err2 != nil {
		log.Println(err2)
	}

	if playerRaw.LeagueID == 0 {
		// Updating player with league
		playerRaw.LeagueID = league.ID
		if playerRaw.Status == "nobody" {
			playerRaw.Status = "common"
		}
		_, err4 := c.Db.NamedExec("UPDATE `players` SET league_id=:league_id, status=:status WHERE id=:id", &playerRaw)
		if err4 != nil {
			log.Println(err4)
			return "fail"
		}
	} else if playerRaw.LeagueID != league.ID {
		// Duplicate profile: user changed league, beware!
		playerRaw.LeagueID = league.ID
		playerRaw.SquadID = 0
		playerRaw.Status = "league_changed"
		playerRaw.CreatedAt = time.Now().UTC()
		_, err5 := c.Db.NamedExec("INSERT INTO players VALUES(NULL, :telegram_id, :league_id, :squad_id, :status, :created_at, :updated_at)", &playerRaw)
		if err5 != nil {
			log.Println(err5)
			return "fail"
		}
		err6 := c.Db.Get(&playerRaw, c.Db.Rebind("SELECT * FROM players WHERE telegram_id='"+strconv.Itoa(playerRaw.TelegramID)+"' AND league_id='"+strconv.Itoa(league.ID)+"';"))
		if err6 != nil {
			log.Println(err6)
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
	profileRaw.Exp = expInt
	profileRaw.EggExp = eggexpInt
	profileRaw.Power = powerInt
	profileRaw.WeaponID = weaponRaw.ID
	profileRaw.Crystalls = crystallsInt
	profileRaw.CreatedAt = time.Now().UTC()

	_, err3 := c.Db.NamedExec("INSERT INTO `profiles` VALUES(NULL, :player_id, :nickname, :telegram_nickname, :level_id, :pokeballs, :wealth, :exp, :egg_exp, :power, :weapon_id, :crystalls, :created_at)", &profileRaw)
	if err3 != nil {
		log.Println(err3)
		return "fail"
	}

	err8 := c.Db.Get(&profileRaw, c.Db.Rebind("SELECT * FROM profiles WHERE player_id=? AND created_at=?"), profileRaw.PlayerID, profileRaw.CreatedAt)
	if err8 != nil {
		log.Println(err8)
		log.Printf("Profile isn't added!")
		return "fail"
	}

	playerRaw.UpdatedAt = time.Now().UTC()
	_, err7 := c.Db.NamedExec("UPDATE `players` SET updated_at=:updated_at WHERE id=:id", &playerRaw)
	if err7 != nil {
		log.Println(err7)
		return "fail"
	}

	for meme, attack := range pokememes {
		rarity := "common"
		if strings.HasPrefix(meme, "🔸") {
			rarity = "rare"
			meme = strings.Replace(meme, "🔸", "", 1)
		}
		p.fillProfilePokememe(profileRaw.ID, meme, attack, rarity)
	}

	return "ok"
}
