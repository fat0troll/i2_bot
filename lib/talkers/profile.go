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

func (t *Talkers) ProfileMessage(update tgbotapi.Update, player_raw dbmapping.Player) string {
    profile_raw, ok := c.Getters.GetProfile(player_raw.Id)
    if !ok {
        c.Talkers.AnyMessageUnauthorized(update)
        return "fail"
    }
    league := dbmapping.League{}
    err := c.Db.Get(&league, c.Db.Rebind("SELECT * FROM leagues WHERE id=?"), player_raw.League_id)
    if err != nil {
        log.Println(err)
    }
    level := dbmapping.Level{}
    err = c.Db.Get(&level, c.Db.Rebind("SELECT * FROM levels WHERE id=?"), profile_raw.Level_id)
    if err != nil {
        log.Println(err)
    }
    weapon := dbmapping.Weapon{}
    if profile_raw.Weapon_id != 0 {
        err = c.Db.Get(&weapon, c.Db.Rebind("SELECT * FROM weapons WHERE id=?"), profile_raw.Weapon_id)
        if err != nil {
            log.Println(err)
        }
    }
    p_pk := []dbmapping.ProfilePokememe{}
    err = c.Db.Select(&p_pk, c.Db.Rebind("SELECT * FROM profiles_pokememes WHERE profile_id=?"), profile_raw.Id)
    if err != nil {
        log.Println(err)
    }
    pokememes := []dbmapping.Pokememe{}
    err = c.Db.Select(&pokememes, c.Db.Rebind("SELECT * FROM pokememes"))
    if err != nil {
        log.Println(err)
    }

    attack_pm := 0
    for i := range(p_pk) {
        for j := range(pokememes) {
            if p_pk[i].Pokememe_id == pokememes[j].Id {
                single_attack := float64(pokememes[j].Attack)
                single_attack = single_attack * float64(p_pk[i].Pokememe_lvl)
                if p_pk[i].Pokememe_rarity == "rare" {
                    single_attack = single_attack * 1.15
                }
                attack_pm += int(single_attack)
            }
        }
    }


    message := "*Профиль игрока "
    message += profile_raw.Nickname + "* (@" + profile_raw.TelegramNickname + ")\n"
    message += "\nЛига: " + league.Symbol + league.Name
    message += "\n👤 " + strconv.Itoa(profile_raw.Level_id)
    message += " | 🎓 " + strconv.Itoa(profile_raw.Exp) + "/" + strconv.Itoa(level.Max_exp)
    message += " | 🥚 " + strconv.Itoa(profile_raw.Egg_exp) + "/" + strconv.Itoa(level.Max_egg)
    message += "\n💲" + c.Parsers.ReturnPoints(profile_raw.Wealth)
    message += " |💎" + strconv.Itoa(profile_raw.Crystalls)
    message += " |⭕" + strconv.Itoa(profile_raw.Pokeballs)
    message += "\n⚔Атака: 1 + " + c.Parsers.ReturnPoints(weapon.Power) + " + " + c.Parsers.ReturnPoints(attack_pm) + "\n"

    if profile_raw.Weapon_id != 0 {
        message += "\n🔫Оружие: " + weapon.Name + " " + c.Parsers.ReturnPoints(weapon.Power) + "⚔"
    }

    message += "\n🐱Покемемы:"
    for i := range(p_pk) {
        for j := range(pokememes) {
            if p_pk[i].Pokememe_id == pokememes[j].Id {
                single_attack := float64(pokememes[j].Attack)
                single_attack = single_attack * float64(p_pk[i].Pokememe_lvl)
                if p_pk[i].Pokememe_rarity == "rare" {
                    single_attack = single_attack * 1.15
                }

                message += "\n" + strconv.Itoa(pokememes[j].Grade)
                message += "⃣ " + pokememes[j].Name
                message += " +" + c.Parsers.ReturnPoints(int(single_attack)) + "⚔"
            }
        }
    }
    message += "\n\n💳" + strconv.Itoa(player_raw.Telegram_id)
    message += "\n⏰Последнее обновление профиля: " + profile_raw.Created_at.Format("02.01.2006 15:04:05")
    message += "\n\nНе забывай обновляться, это важно для получения актуальной информации.\n\n"
    message += "/best – посмотреть лучших покемемов для поимки"

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
    msg.ParseMode = "Markdown"

    c.Bot.Send(msg)

    return "ok"
}
