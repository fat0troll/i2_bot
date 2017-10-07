// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// stdlib
	"log"
	"strings"
    "strconv"
    // 3rd party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"../dbmappings"
)

type PokememeFull struct {
    Pokememe    dbmappings.Pokememes
    Elements    []dbmappings.Elements
    Locations   []dbmappings.Locations
}

func (t *Talkers) PokememeInfo(update tgbotapi.Update) string {
	pokememe_number := strings.Replace(update.Message.Text, "/pk", "", 1)

	// Building pokememe
	pk := dbmappings.Pokememes{}
	// Checking if pokememe exists in database
	err := c.Db.Get(&pk, c.Db.Rebind("SELECT * FROM pokememes WHERE id='" + pokememe_number + "'"))
	if err != nil {
		log.Println(err)
		return "fail"
	}

	elements := []dbmappings.Elements{}
	err = c.Db.Select(&elements, "SELECT * FROM elements");
	if err != nil {
		log.Println(err)
	}
	locations := []dbmappings.Locations{}
	err = c.Db.Select(&locations, "SELECT * FROM locations");
	if err != nil {
		log.Println(err)
	}

	pokememes_elements := []dbmappings.PokememesElements{}
	err = c.Db.Select(&pokememes_elements, "SELECT * FROM pokememes_elements WHERE pokememe_id='" + pokememe_number + "'");
	if err != nil {
		log.Println(err)
	}
	pokememes_locations := []dbmappings.PokememesLocations{}
	err = c.Db.Select(&pokememes_locations, "SELECT * FROM pokememes_locations WHERE pokememe_id='" + pokememe_number + "'");
	if err != nil {
		log.Println(err)
	}

	message := strconv.Itoa(pk.Grade) + "⃣ *" + pk.Name + "*\n"
	message += pk.Description + "\n\n"
	message += "Элементы:"
	for i := range(pokememes_elements) {
		for j := range(elements) {
			if pokememes_elements[i].Element_id == elements[j].Id {
				message += " " + elements[j].Symbol
			}
		}
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
	for i := range(pokememes_locations) {
		for j := range(locations) {
			if pokememes_locations[i].Location_id == locations[j].Id {
				message += " *" + locations[j].Name + "*"
				if (i + 1) < len(pokememes_locations) {
					message += ","
				}
			}
		}
	}

	message += "\n" + pk.Image_url

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for i := range(pokememes_locations) {
		for j := range(locations) {
			if pokememes_locations[i].Location_id == locations[j].Id {
			   var row []tgbotapi.InlineKeyboardButton
			   btn := tgbotapi.NewInlineKeyboardButtonSwitch(locations[j].Symbol + locations[j].Name, locations[j].Symbol + locations[j].Name)
			   row = append(row, btn)
			   keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
			}
		}
	}

	msg.ReplyMarkup = keyboard
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

func (t *Talkers) PokememesList(update tgbotapi.Update, page int) {
	pokememes := []dbmappings.Pokememes{}
	err := c.Db.Select(&pokememes, "SELECT * FROM pokememes");
	if err != nil {
		log.Println(err)
	}
    pokememes_limited := []dbmappings.Pokememes{}
    err = c.Db.Select(&pokememes_limited, "SELECT * FROM pokememes ORDER BY grade asc, name asc LIMIT 50 OFFSET " + strconv.Itoa(50*(page-1)));
    if err != nil {
        log.Println(err)
    }
	elements := []dbmappings.Elements{}
	err = c.Db.Select(&elements, "SELECT * FROM elements");
	if err != nil {
		log.Println(err)
	}
	locations := []dbmappings.Locations{}
	err = c.Db.Select(&locations, "SELECT * FROM locations");
	if err != nil {
		log.Println(err)
	}
	pokememes_elements := []dbmappings.PokememesElements{}
	err = c.Db.Select(&pokememes_elements, "SELECT * FROM pokememes_elements");
	if err != nil {
		log.Println(err)
	}
	pokememes_locations := []dbmappings.PokememesLocations{}
	err = c.Db.Select(&pokememes_locations, "SELECT * FROM pokememes_locations");
	if err != nil {
		log.Println(err)
	}

    pokememes_full := []PokememeFull{}

    for i := range(pokememes_limited) {
        full_pokememe := PokememeFull{}
        elements_listed := []dbmappings.Elements{}
        locations_listed := []dbmappings.Locations{}

        for j := range(pokememes_locations) {
            if pokememes_locations[j].Pokememe_id == pokememes_limited[i].Id {
                for l := range(locations) {
                    if pokememes_locations[j].Location_id == locations[l].Id {
                        locations_listed = append(locations_listed, locations[l])
                    }
                }
            }
        }

        for k := range(pokememes_elements) {
            if pokememes_elements[k].Pokememe_id == pokememes_limited[i].Id {
                for e := range(elements) {
                    if pokememes_elements[k].Element_id == elements[e].Id {
                        elements_listed = append(elements_listed, elements[e])
                    }
                }
            }
        }

        full_pokememe.Pokememe = pokememes_limited[i]
        full_pokememe.Elements = elements_listed
        full_pokememe.Locations = locations_listed

        pokememes_full = append(pokememes_full, full_pokememe)
    }

	message := "*Известные боту покемемы*\n"
	message += "Список отсортирован по грейду и алфавиту.\n"
    message += "Покедекс: " + strconv.Itoa(len(pokememes)) + " / 206\n"
    message += "Отображаем покемемов с " + strconv.Itoa(((page - 1)*50)+1) + " по " + strconv.Itoa(page*50) + "\n"
    if len(pokememes) > page*50 {
        message += "Переход на следующую страницу: /pokedeks" + strconv.Itoa(page + 1)
    }
    if page > 1 {
        message += "\nПереход на предыдущую страницу: /pokedeks" + strconv.Itoa(page - 1)
    }
    message += "\n\n"

    for i := range(pokememes_full) {
        pk := pokememes_full[i].Pokememe
        pk_e := pokememes_full[i].Elements
        message += strconv.Itoa(i + 1 + (50*(page-1))) + ". " + strconv.Itoa(pk.Grade)
        message += "⃣ *" + pk.Name
        message += "* (" + c.Parsers.ReturnPoints(pk.HP) + "-" + c.Parsers.ReturnPoints(pk.MP) + ") ⚔️ *"
        message += c.Parsers.ReturnPoints(pk.Attack) + "* \\["
        for j := range(pk_e) {
            message += pk_e[j].Symbol
        }
        message += "] " + c.Parsers.ReturnPoints(pk.Price) + "$ /pk" + strconv.Itoa(pk.Id)
        message += "\n"
    }

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)
}
