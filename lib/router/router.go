// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
    // stdlib
    "fmt"
    "log"
    "regexp"
    "strings"
    "time"
    // 3rd party
	"gopkg.in/telegram-bot-api.v4"
    // local
    "../dbmappings"
)

type Router struct {}

// This function will route requests to appropriative modules
// It will return "ok" or "fail"
// If command doesn't exist, it's "fail"
func (r *Router) RouteRequest(update tgbotapi.Update) string {
    text := update.Message.Text
    user_id := update.Message.From.ID

    player_raw := dbmappings.Players{}
    err := c.Db.Get(&player_raw, c.Db.Rebind("SELECT * FROM players WHERE telegram_id=?"), user_id)
    if err != nil {
        log.Printf("Message user not found in database.")
        log.Printf(err.Error())

        // Create "nobody" user
        player_raw.Telegram_id = user_id
        player_raw.League_id = 0
        player_raw.Squad_id = 0
        player_raw.Status = "nobody"
        player_raw.Created_at = time.Now().UTC()
        player_raw.Updated_at = time.Now().UTC()
        _, erradd := c.Db.NamedExec("INSERT INTO players VALUES(NULL, :telegram_id, :league_id, :squad_id, :status, :created_at, :updated_at)", &player_raw)
        if erradd != nil {
            log.Printf(erradd.Error())
            return "fail"
        }
    } else {
        log.Printf("Message user found in database.")
    }

    fmt.Println(player_raw)

    // Regular expressions
    var durakMsg = regexp.MustCompile("(Д|д)(У|у)(Р|р)(А|а|Е|е|О|о)")
    var huMsg = regexp.MustCompile("(Х|х)(У|у)(Й|й|Я|я|Ю|ю|Е|е)")
    var blMsg = regexp.MustCompile("(\\s|^)(Б|б)(Л|л)(Я|я)(Т|т|Д|д)")
    var ebMsg = regexp.MustCompile("(\\s|^|ЗА|За|зА|за)(Е|е|Ё|ё)(Б|б)(\\s|Л|л|А|а|Т|т|У|у|Е|е|Ё|ё|И|и)")
    var piMsg = regexp.MustCompile("(П|п)(И|и)(З|з)(Д|д)")

    // Commands
    var helpMsg = regexp.MustCompile("/help\\z")
    var helloMsg = regexp.MustCompile("/start\\z")
    var pokedexMsg = regexp.MustCompile("/pokede(x|ks)\\d?\\z")

    // Forwards
    var pokememeMsg = regexp.MustCompile("(Уровень)(.+)(Опыт)(.+)\n(Элементы:)(.+)\n(.+)(💙MP)")

    if update.Message.ForwardFrom != nil {
        if update.Message.ForwardFrom.ID != 360402625 {
            log.Printf("Forward from another user or bot. Ignoring")
        } else {
            log.Printf("Forward from PokememBro bot! Processing...")
            if pokememeMsg.MatchString(text) {
                if player_raw.Id != 0 {
                    log.Printf("Pokememe posted!")
                    status := c.Parsers.ParsePokememe(text, player_raw)
                    switch status {
                    case "ok":
                        c.Talkers.PokememeAddSuccessMessage(update)
                    case "dup":
                        c.Talkers.PokememeAddDuplicateMessage(update)
                    case "fail":
                        c.Talkers.PokememeAddFailureMessage(update)
                    }
                } else {
                    c.Talkers.AnyMessageUnauthorized(update)
                }
            } else {
                log.Printf(text)
            }
        }
    } else {
        // Direct messages from user
        switch {
        case helloMsg.MatchString(text):
            if player_raw.Id != 0 {
                c.Talkers.HelloMessageAuthorized(update, player_raw)
            } else {
                c.Talkers.HelloMessageUnauthorized(update)
            }
        // Help
        case helpMsg.MatchString(text):
            c.Talkers.HelpMessage(update)
        // Pokememes info
        case pokedexMsg.MatchString(text):
            if strings.HasSuffix(text, "1") {
                c.Talkers.PokememesList(update, 1)
            } else if strings.HasSuffix(text, "2") {
                c.Talkers.PokememesList(update, 2)
            } else if strings.HasSuffix(text, "3") {
                c.Talkers.PokememesList(update, 3)
            } else if strings.HasSuffix(text, "4") {
                c.Talkers.PokememesList(update, 4)
            } else if strings.HasSuffix(text, "5") {
                c.Talkers.PokememesList(update, 5)
            } else {
                c.Talkers.PokememesList(update, 1)
            }
        // Easter eggs
        case huMsg.MatchString(text):
            c.Talkers.MatMessage(update)
        case blMsg.MatchString(text):
            c.Talkers.MatMessage(update)
        case ebMsg.MatchString(text):
            c.Talkers.MatMessage(update)
        case piMsg.MatchString(text):
            c.Talkers.MatMessage(update)
        case durakMsg.MatchString(text):
            c.Talkers.DurakMessage(update)
        default:
            log.Printf("User posted unknown command.")
            return "fail"
        }
    }

    return "ok"
}
