// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
    // stdlib
    "log"
    "regexp"
    // 3rd party
	"gopkg.in/telegram-bot-api.v4"
    // local
    "../talkers"
)

type Router struct {}

// This function will route requests to appropriative modules
// It will return "ok" or "fail"
// If command doesn't exist, it's "fail"
func (r *Router) RouteRequest(update tgbotapi.Update) string {
    text := update.Message.Text

    // Regular expressions
    var durakMsg = regexp.MustCompile("(Д|д)(У|у)(Р|р)(А|а|Е|е|О|о)")
    var huMsg = regexp.MustCompile("(Х|х)(У|у)(Й|й|Я|я|Ю|ю|Е|е)")
    var blMsg = regexp.MustCompile("\\s(Б|б)(Л|л)(Я|я)(Т|т|Д|д)")
    var ebMsg = regexp.MustCompile("(Е|е|Ё|ё)(Б|б)(\\s|А|а|Т|т|У|у|Е|е|Ё|ё|И|и)")
    var piMsg = regexp.MustCompile("(П|п)(И|и)(З|з)(Д|д)")
    var helpMsg = regexp.MustCompile("/help\\z")

    switch {
    case helpMsg.MatchString(text):
        talkers.HelpMessage(c.Bot, update)
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

    return "ok"
}
