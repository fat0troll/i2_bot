// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	// stdlib
	"log"
    "strconv"
    // 3rd party
	"gopkg.in/telegram-bot-api.v4"
	// local
	"../dbmappings"
)

type PokememeFull struct {
    Pokememe    dbmappings.Pokememes
    Elements    []dbmappings.Elements
    Locations   []dbmappings.Locations
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
