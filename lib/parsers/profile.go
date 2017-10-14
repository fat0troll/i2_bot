// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsers

import (
    // stdlib
    "log"
    "regexp"
    "strings"
    "strconv"
    "time"
    // 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
    // local
    "../dbmapping"
)

// Internal functions

func (p *Parsers) fillProfilePokememe(profile_id int, meme string, attack string, rarity string) {
    spk_raw := dbmapping.Pokememe{}
    err := c.Db.Get(&spk_raw, c.Db.Rebind("SELECT * FROM pokememes WHERE name='" + meme + "';"))
    if err != nil {
        log.Println(err)
    } else {
        attack_int := p.getPoints(attack)
        // Improve it. Game's precision is unstable
        orig_attack := float64(spk_raw.Attack)
        if rarity == "rare" {
            orig_attack = orig_attack * 1.1
        }
        level := int(float64(attack_int) / orig_attack)

        ppk := dbmapping.ProfilePokememe{}
        ppk.Profile_id = profile_id
        ppk.Pokememe_id = spk_raw.Id
        ppk.Pokememe_lvl = level
        ppk.Pokememe_rarity = rarity
        ppk.Created_at = time.Now().UTC()
        _, err2 := c.Db.NamedExec("INSERT INTO `profiles_pokememes` VALUES(NULL, :profile_id, :pokememe_id, :pokememe_lvl, :pokememe_rarity, :created_at)", &ppk)
        if err2 != nil {
            log.Println(err2)
        }
    }
}

// External functions

func (p *Parsers) ParseProfile(update tgbotapi.Update, player_raw dbmapping.Player) string {
    text := update.Message.Text
    log.Println(text)

    profile_info_strings := strings.Split(text, "\n")
    profile_info_runed_strings := make([][]rune, 0)
    for i := range(profile_info_strings) {
        profile_info_runed_strings = append(profile_info_runed_strings, []rune(profile_info_strings[i]))
    }

    league := dbmapping.League{}

    telegram_nickname := update.Message.From.UserName
    nickname := ""
    level := ""
    level_int := 0
    exp := ""
    exp_int := 0
    egg_exp := ""
    egg_exp_int := 0
    pokeballs := ""
    pokeballs_int := 0
    wealth := ""
    wealth_int := 0
    crystalls := ""
    crystalls_int := 0
    weapon := ""
    pokememes := make(map[string]string)
    power_int := 1

    // Filling information
    // We don't know how many strings we got, so we iterating each other
    for i := range(profile_info_runed_strings) {
        current_string := string(profile_info_runed_strings[i])
        current_runes := profile_info_runed_strings[i]
        if strings.HasPrefix(current_string, "🈸") || strings.HasPrefix(current_string, "🈳 ") || strings.HasPrefix(current_string, "🈵") {
            err1 := c.Db.Get(&league, c.Db.Rebind("SELECT * FROM leagues WHERE symbol='" + string(current_runes[0]) + "'"))
            if err1 != nil {
                log.Println(err1)
                return "fail"
            }
            for j := range(current_runes) {
                if j > 1 {
                    nickname += string(current_runes[j])
                }
            }
        }
        if strings.HasPrefix(current_string, "👤Уровень:") {
            levelRx := regexp.MustCompile("\\d+")
            level_array := levelRx.FindAllString(current_string, -1)
            if len(level_array) < 1 {
                log.Println("Level string broken")
                return "fail"
            }
            level = level_array[0]
            level_int, _ = strconv.Atoi(level)
        }

        if strings.HasPrefix(current_string, "🎓Опыт") {
            expRx := regexp.MustCompile("\\d+")
            exp_array := expRx.FindAllString(current_string, -1)
            if len(exp_array) < 4 {
                log.Println("Exp string broken")
                return "fail"
            }
            exp = exp_array[0]
            exp_int, _ = strconv.Atoi(exp)
            egg_exp = exp_array[2]
            egg_exp_int, _ = strconv.Atoi(egg_exp)
        }

        if strings.HasPrefix(current_string, "⭕Покеболы") {
            pkbRx := regexp.MustCompile("\\d+")
            pkb_array := pkbRx.FindAllString(current_string, -1)
            if len(pkb_array) < 2 {
                log.Println("Pokeballs string broken")
                return "fail"
            }
            pokeballs = pkb_array[1]
            pokeballs_int, _ = strconv.Atoi(pokeballs)
        }

        if strings.HasPrefix(current_string, "💲") {
            wealthRx := regexp.MustCompile("(\\d|\\.|K|M)+")
            wealth_array := wealthRx.FindAllString(current_string, -1)
            if len(wealth_array) < 2 {
                log.Println("Wealth string broken")
                return "fail"
            }
            wealth = wealth_array[0]
            wealth_int = p.getPoints(wealth)
            crystalls  =wealth_array[1]
            crystalls_int = p.getPoints(crystalls)
        }

        if strings.HasPrefix(current_string, "🔫") {
            // We need NEXT string!
            weapon_type_string := strings.Replace(current_string, "🔫 ", "", 1)
            wnRx := regexp.MustCompile("(.+)(ита)")
            weapon = wnRx.FindString(weapon_type_string)
        }

        if strings.HasPrefix(current_string, "🐱Покемемы: ") {
            pkmnumRx := regexp.MustCompile("\\d+")
            pk_num_array := pkmnumRx.FindAllString(current_string, -1)
            if len(pk_num_array) < 2 {
                log.Println("Pokememes count broken")
                return "fail"
            }
            pokememes_count, _ := strconv.Atoi(pk_num_array[0])
            if pokememes_count > 0 {
                for pi := 0; pi < pokememes_count; pi++ {
                    pokememe_string := string(profile_info_runed_strings[i + 1 + pi])
                    attackRx := regexp.MustCompile("(\\d|\\.|K|M)+")
                    pk_points_array := attackRx.FindAllString(pokememe_string, -1)
                    pk_attack := pk_points_array[1]
                    pk_name := strings.Split(pokememe_string, "+")[0]
                    pk_name = strings.Replace(pk_name, " ⭐", "", 1)
                    pk_name = strings.TrimSuffix(pk_name, " ")
                    pk_name = strings.Split(pk_name, "⃣ ")[1]
                    pokememes[pk_name] = pk_attack
                    power_int += p.getPoints(pk_attack)
                }
            }
        }
    }

    log.Printf("Telegram nickname: " + telegram_nickname)
    log.Printf("Nickname: " + nickname)
    log.Printf("League: " + league.Name)
    log.Printf("Level: " + level)
    log.Println(level_int)
    log.Printf("Exp: " + exp)
    log.Println(exp_int)
    log.Printf("Egg exp: " + egg_exp)
    log.Println(egg_exp_int)
    log.Printf("Pokeballs: " + pokeballs)
    log.Println(pokeballs_int)
    log.Printf("Wealth: " + wealth)
    log.Println(wealth_int)
    log.Printf("Crystalls: " + crystalls)
    log.Println(crystalls_int)
    log.Printf("Weapon: " + weapon)
    if len(pokememes) > 0 {
        for meme, attack := range(pokememes) {
            log.Printf(meme + ": " + attack)
        }
    } else {
        log.Printf("Hand is empty.")
    }

    // Information is gathered, let's create profile in database!
    weapon_raw := dbmapping.Weapon{}
    err2 := c.Db.Get(&weapon_raw, c.Db.Rebind("SELECT * FROM weapons WHERE name='" + weapon + "'"))
    if err2 != nil {
        log.Println(err2)
    }

    if player_raw.League_id == 0 {
        // Updating player with league
        player_raw.League_id = league.Id
        if player_raw.Status == "nobody" {
            player_raw.Status = "common"
        }
        _, err4 := c.Db.NamedExec("UPDATE `players` SET league_id=:league_id, status=:status WHERE id=:id", &player_raw)
        if err4 != nil {
            log.Println(err4)
            return "fail"
        }
    } else if player_raw.League_id != league.Id {
        // Duplicate profile: user changed league, beware!
        player_raw.League_id = league.Id
        player_raw.Squad_id = 0
        player_raw.Status = "league_changed"
        player_raw.Created_at = time.Now().UTC()
        _, err5 := c.Db.NamedExec("INSERT INTO players VALUES(NULL, :telegram_id, :league_id, :squad_id, :status, :created_at, :updated_at)", &player_raw)
        if err5 != nil {
            log.Println(err5)
            return "fail"
        }
        err6 := c.Db.Get(&player_raw, c.Db.Rebind("SELECT * FROM players WHERE telegram_id='" + strconv.Itoa(player_raw.Telegram_id) + "' AND league_id='" + strconv.Itoa(league.Id) + "';"))
        if err6 != nil {
            log.Println(err6)
            return "fail"
        }
    }

    profile_raw := dbmapping.Profile{}
    profile_raw.Player_id = player_raw.Id
    profile_raw.Nickname = nickname
    profile_raw.TelegramNickname = telegram_nickname
    profile_raw.Level_id = level_int
    profile_raw.Pokeballs = pokeballs_int
    profile_raw.Wealth = wealth_int
    profile_raw.Exp = exp_int
    profile_raw.Egg_exp = egg_exp_int
    profile_raw.Power = power_int
    profile_raw.Weapon_id = weapon_raw.Id
    profile_raw.Crystalls = crystalls_int
    profile_raw.Created_at = time.Now().UTC()

    _, err3 := c.Db.NamedExec("INSERT INTO `profiles` VALUES(NULL, :player_id, :nickname, :telegram_nickname, :level_id, :pokeballs, :wealth, :exp, :egg_exp, :power, :weapon_id, :crystalls, :created_at)", &profile_raw)
    if err3 != nil {
        log.Println(err3)
        return "fail"
    }

    err8 := c.Db.Get(&profile_raw, c.Db.Rebind("SELECT * FROM profiles WHERE player_id=? AND created_at=?"), profile_raw.Player_id, profile_raw.Created_at)
    if err8 != nil {
        log.Println(err8)
        log.Printf("Profile isn't added!")
        return "fail"
    }

    player_raw.Updated_at = time.Now().UTC()
    _, err7 := c.Db.NamedExec("UPDATE `players` SET updated_at=:updated_at WHERE id=:id", &player_raw)
    if err7 != nil {
        log.Println(err7)
        return "fail"
    }


    for meme, attack := range(pokememes) {
        rarity := "common"
        if strings.HasPrefix(meme, "🔸") {
            rarity = "rare"
            meme = strings.Replace(meme, "🔸", "", 1)
        }
        p.fillProfilePokememe(profile_raw.Id, meme, attack, rarity)
    }

    return "ok"
}
