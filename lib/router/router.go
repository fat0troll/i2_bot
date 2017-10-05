// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
    // stdlib
    "fmt"
    "log"
    "regexp"
    // 3rd party
	"gopkg.in/telegram-bot-api.v4"
    // local
    "../dbmappings"
    "../talkers"
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
    } else {
        log.Printf("Message user found in database.")
    }

    fmt.Println(player_raw)

    // Regular expressions
    var durakMsg = regexp.MustCompile("(Д|д)(У|у)(Р|р)(А|а|Е|е|О|о)")
    var huMsg = regexp.MustCompile("(Х|х)(У|у)(Й|й|Я|я|Ю|ю|Е|е)")
    var blMsg = regexp.MustCompile("\\s(Б|б)(Л|л)(Я|я)(Т|т|Д|д)")
    var ebMsg = regexp.MustCompile("(Е|е|Ё|ё)(Б|б)(\\s|А|а|Т|т|У|у|Е|е|Ё|ё|И|и)")
    var piMsg = regexp.MustCompile("(П|п)(И|и)(З|з)(Д|д)")
    var helpMsg = regexp.MustCompile("/help\\z")
    var helloMsg = regexp.MustCompile("/start\\z")

    var pokememeMsg = regexp.MustCompile("(Уровень)(.+)(Опыт)(.+)\n(Элементы:)(.+)\n(.+)(💙MP)")

    if update.Message.ForwardFrom != nil {
        if update.Message.ForwardFrom.ID != 360402625 {
            log.Printf("Forward from another user or bot. Ignoring")
        } else {
            log.Printf("Forward from PokememBro bot! Processing...")
            if pokememeMsg.MatchString(text) {
                if player_raw.Id != 0 {
                    log.Printf("Pokememe posted!")
                    c.Parsers.ParsePokememe(text, player_raw)
                } else {
                    talkers.AnyMessageUnauthorized(c.Bot, update)
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
                talkers.HelloMessageAuthorized(c.Bot, update, player_raw)
            } else {
                talkers.HelloMessageUnauthorized(c.Bot, update)
            }
        // Help
        case helpMsg.MatchString(text):
            talkers.HelpMessage(c.Bot, update)
        // Easter eggs
        case huMsg.MatchString(text):
            talkers.MatMessage(c.Bot, update)
        case blMsg.MatchString(text):
            talkers.MatMessage(c.Bot, update)
        case ebMsg.MatchString(text):
            talkers.MatMessage(c.Bot, update)
        case piMsg.MatchString(text):
            talkers.MatMessage(c.Bot, update)
        case durakMsg.MatchString(text):
            talkers.DurakMessage(c.Bot, update)
        default:
            log.Printf("User posted unknown command.")
            return "fail"
        }
    }

    return "ok"
}
