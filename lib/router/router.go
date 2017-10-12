// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
    // stdlib
    "log"
    "regexp"
    "strings"
    // 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Router struct {}

// This function will route requests to appropriative modules
// It will return "ok" or "fail"
// If command doesn't exist, it's "fail"
func (r *Router) RouteRequest(update tgbotapi.Update) string {
    text := update.Message.Text

    player_raw, ok := c.Getters.GetOrCreatePlayer(update.Message.From.ID)
    if !ok {
        // Silently fail
        return "fail"
    }

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
    var pokememeInfoMsg = regexp.MustCompile("/pk(\\d+)")
    var meMsg = regexp.MustCompile("/me\\z")
    var bestMsg = regexp.MustCompile("/best\\z")

    // Forwards
    var pokememeMsg = regexp.MustCompile("(Уровень)(.+)(Опыт)(.+)\n(Элементы:)(.+)\n(.+)(💙MP)")
    var profileMsg = regexp.MustCompile("(Онлайн: )(\\d+)\n(Турнир Лиг через)(.+)\n\n(.*)\n(Элементы)(.+)\n\n(.+)(Уровень)(.+)\n")

    if update.Message.ForwardFrom != nil {
        if update.Message.ForwardFrom.ID != 360402625 {
            log.Printf("Forward from another user or bot. Ignoring")
        } else {
            log.Printf("Forward from PokememBro bot! Processing...")
            if player_raw.Id != 0 {
                switch {
                case pokememeMsg.MatchString(text):
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
                case profileMsg.MatchString(text):
                    log.Printf("Profile posted!")
                    status := c.Parsers.ParseProfile(update, player_raw)
                    switch status {
                    case "ok":
                        c.Talkers.ProfileAddSuccessMessage(update)
                    case "fail":
                        c.Talkers.ProfileAddFailureMessage(update)
                    }
                default:
                    log.Printf(text)
                }
            } else {
                c.Talkers.AnyMessageUnauthorized(update)
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
        case pokememeInfoMsg.MatchString(text):
            c.Talkers.PokememeInfo(update, player_raw)
        // Profile info
        case meMsg.MatchString(text):
            if player_raw.Id != 0 {
                c.Talkers.ProfileMessage(update, player_raw)
            } else {
                c.Talkers.AnyMessageUnauthorized(update)
            }
        // Suggestions
        case bestMsg.MatchString(text):
            c.Talkers.BestPokememesList(update, player_raw)
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
