// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// stdlib
	"strings"
    "strconv"
    // 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../dbmapping"
)

// Internal functions

func (t *Talkers) pokememesListing(update tgbotapi.Update, page int, pokememes_full []dbmapping.PokememeFull) {
    message := "*Известные боту покемемы*\n"
    message += "Список отсортирован по грейду и алфавиту.\n"
    message += "Покедекс: " + strconv.Itoa(len(pokememes_full)) + " / 206\n"
    message += "Отображаем покемемов с " + strconv.Itoa(((page - 1)*50)+1) + " по " + strconv.Itoa(page*50) + "\n"
    if len(pokememes_full) > page*50 {
        message += "Переход на следующую страницу: /pokedeks" + strconv.Itoa(page + 1)
    }
    if page > 1 {
        message += "\nПереход на предыдущую страницу: /pokedeks" + strconv.Itoa(page - 1)
    }
    message += "\n\n"

    for i := range(pokememes_full) {
        if (i + 1 > 50*(page - 1)) && (i + 1 < (50*page) + 1) {
            pk := pokememes_full[i].Pokememe
            pk_e := pokememes_full[i].Elements
            message += strconv.Itoa(i + 1) + ". " + strconv.Itoa(pk.Grade)
            message += "⃣ *" + pk.Name
            message += "* (" + c.Parsers.ReturnPoints(pk.HP) + "-" + c.Parsers.ReturnPoints(pk.MP) + ") ⚔️ *"
            message += c.Parsers.ReturnPoints(pk.Attack) + "* \\["
            for j := range(pk_e) {
                message += pk_e[j].Symbol
            }
            message += "] " + c.Parsers.ReturnPoints(pk.Price) + "$ /pk" + strconv.Itoa(pk.Id)
            message += "\n"
        }
    }

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
    msg.ParseMode = "Markdown"

    c.Bot.Send(msg)

}

// External functions

func (t *Talkers) PokememesList(update tgbotapi.Update, page int) {
    pokememes_full, ok := c.Getters.GetPokememes()
    if !ok {
        t.GetterError(update)
    } else {
        t.pokememesListing(update, page, pokememes_full)
    }
}

func (t *Talkers) PokememeInfo(update tgbotapi.Update, player_raw dbmapping.Player) string {
    pokememe_number := strings.Replace(update.Message.Text, "/pk", "", 1)
    var calculate_possibilites bool = true
    profile_raw, ok := c.Getters.GetProfile(player_raw.Id)
    if !ok {
        calculate_possibilites = false
    }

    pokememe, ok := c.Getters.GetPokememeByID(pokememe_number)
    if !ok {
        return "fail"
    }

    pk := pokememe.Pokememe

    message := strconv.Itoa(pk.Grade) + "⃣ *" + pk.Name + "*\n"
    message += pk.Description + "\n\n"
    message += "Элементы:"
    for i := range(pokememe.Elements) {
        message += " " + pokememe.Elements[i].Symbol
    }
    message += "\n⚔ Атака: *" + c.Parsers.ReturnPoints(pk.Attack)
    message += "*\n❤️ HP: *" + c.Parsers.ReturnPoints(pk.HP)
    message += "*\n💙 MP: *" + c.Parsers.ReturnPoints(pk.MP)
    if (pk.Defence != pk.Attack) {
        message += "*\n🛡Защита: *" + c.Parsers.ReturnPoints(pk.Defence) + "* _(сопротивляемость покемема к поимке)_"
    } else {
        message += "*"
    }
    message += "\nСтоимость: *" + c.Parsers.ReturnPoints(pk.Price)
    message += "*\nКупить: *"
    if pk.Purchaseable {
        message += "Можно"
    } else {
        message += "Нельзя"
    }
    message += "*\nОбитает:"
    for i := range(pokememe.Locations) {
        message += " *" + pokememe.Locations[i].Name + "*"
        if (i + 1) < len(pokememe.Locations) {
            message += ","
        }
    }

    if calculate_possibilites {
        if (pk.Grade < profile_raw.Level_id + 2) || (pk.Grade > profile_raw.Level_id - 3) {
            message += "\nВероятность поимки:"
            for i := range(pokememe.Locations) {
                percentile, pokeballs := c.Getters.PossibilityRequiredPokeballs(pokememe.Locations[i].Id, pk.Grade, profile_raw.Level_id)
                message += "\n" + pokememe.Locations[i].Name + " – "
                message += strconv.FormatFloat(percentile, 'f', 2, 64) + "% или "
                message += strconv.Itoa(pokeballs) + "⭕"
            }
        }
    }

    message += "\n" + pk.Image_url

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
    keyboard := tgbotapi.InlineKeyboardMarkup{}
    for i := range(pokememe.Locations) {
       var row []tgbotapi.InlineKeyboardButton
       btn := tgbotapi.NewInlineKeyboardButtonSwitch(pokememe.Locations[i].Symbol + pokememe.Locations[i].Name, pokememe.Locations[i].Symbol + pokememe.Locations[i].Name)
       row = append(row, btn)
       keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
    }

    msg.ReplyMarkup = keyboard
    msg.ParseMode = "Markdown"

    c.Bot.Send(msg)

    return "ok"
}
