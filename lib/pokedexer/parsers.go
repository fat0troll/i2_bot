// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"regexp"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
	// "time"
)

// ParsePokememe parses pokememe, forwarded from PokememeBroBot, to database
func (p *Pokedexer) ParsePokememe(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	pokememeStringsArray := strings.Split(update.Message.Text, "\n")
	pokememeRunesArray := make([][]rune, 0)
	for i := range pokememeStringsArray {
		pokememeRunesArray = append(pokememeRunesArray, []rune(pokememeStringsArray[i]))
	}

	pokememeData := make(map[string]string)
	pokememeLocations := make(map[string]string)
	pokememeElements := make(map[string]string)

	hitPointsRx := regexp.MustCompile("(\\d|\\.)+(K|M)?")

	for i := range pokememeStringsArray {
		c.Log.Debug("Processing string: " + pokememeStringsArray[i])
		if strings.Contains(pokememeStringsArray[i], "⃣") {
			// Strings with name and grade
			pokememeData["grade"] = string(pokememeRunesArray[i][0])
			pokememeData["name"] = string(pokememeRunesArray[i][3:])
		}

		if i == 1 {
			pokememeData["description"] = string(pokememeRunesArray[i])
		}

		if strings.HasPrefix(pokememeStringsArray[i], "Обитает: ") {
			// Elements
			locationsString := strings.TrimPrefix(pokememeStringsArray[i], "Обитает: ")
			locationsArray := strings.Split(locationsString, ", ")
			for i := range locationsArray {
				pokememeLocations[strconv.Itoa(i)] = locationsArray[i]
			}
		}

		if strings.HasPrefix(pokememeStringsArray[i], "Элементы:  ") {
			// Elements
			elementsString := strings.TrimPrefix(pokememeStringsArray[i], "Элементы:  ")
			elementsArray := strings.Split(elementsString, " ")
			for i := range elementsArray {
				pokememeElements[strconv.Itoa(i)] = elementsArray[i]
			}
		}

		if strings.HasPrefix(pokememeStringsArray[i], "⚔Атака: ") {
			// Attack, HP, MP
			hitPoints := hitPointsRx.FindAllString(string(pokememeRunesArray[i]), -1)
			if len(hitPoints) != 3 {
				c.Log.Error("Can't parse hitpoints!")
				c.Log.Debug("Points string was: " + string(pokememeRunesArray[i]))
				p.pokememeAddFailureMessage(update)
				return "fail"
			}
			pokememeData["attack"] = hitPoints[0]
			pokememeData["hp"] = hitPoints[1]
			pokememeData["mp"] = hitPoints[2]
		}

		if strings.HasPrefix(pokememeStringsArray[i], "🛡Защита ") {
			// Defence for top-level pokememes
			defence := hitPointsRx.FindAllString(string(pokememeRunesArray[i]), -1)
			if len(defence) != 1 {
				c.Log.Error("Can't parse defence!")
				c.Log.Debug("Defence string was: " + string(pokememeRunesArray[i]))
				p.pokememeAddFailureMessage(update)
				return "fail"
			}
			pokememeData["defence"] = defence[0]
		}

		if strings.HasPrefix(pokememeStringsArray[i], "Стоимость :") {
			// Price
			price := hitPointsRx.FindAllString(string(pokememeRunesArray[i]), -1)
			if len(price) != 1 {
				c.Log.Error("Can't parse price!")
				c.Log.Debug("Price string was: " + string(pokememeRunesArray[i]))
				p.pokememeAddFailureMessage(update)
				return "fail"
			}
			pokememeData["price"] = price[0]
		}

		if strings.HasPrefix(pokememeStringsArray[i], "Купить: ") {
			// Purchaseability
			pokememeData["purchaseable"] = "false"
			if strings.Contains(pokememeStringsArray[i], "Можно") {
				pokememeData["purchaseable"] = "true"
			}
		}
	}

	// Image
	for _, entity := range *update.Message.Entities {
		if entity.Type == "text_link" && entity.URL != "" {
			pokememeData["image"] = entity.URL
		}
	}

	// Checking grade to be integer
	_, err := strconv.Atoi(pokememeData["grade"])
	if err != nil {
		c.Log.Error(err.Error())
		p.pokememeAddFailureMessage(update)
		return "fail"
	}

	pokememeData["creator_id"] = strconv.Itoa(playerRaw.ID)

	c.Log.Debugln("Pokememe data: ", pokememeData)
	c.Log.Debugln("Elements: ", pokememeElements)
	c.Log.Debugln("Locations: ", pokememeLocations)

	if len(pokememeElements) == 0 {
		p.pokememeAddFailureMessage(update)
		return "fail"
	}

	if len(pokememeLocations) == 0 {
		p.pokememeAddFailureMessage(update)
		return "fail"
	}

	_, err = c.DataCache.GetPokememeByName(pokememeData["name"])
	if err == nil {
		// There is already a pokememe with such name, updating
		pokememeID, err := c.DataCache.UpdatePokememe(pokememeData, pokememeLocations, pokememeElements)
		if err != nil {
			c.Log.Error(err.Error())
			p.pokememeAddFailureMessage(update)
			return "fail"
		}

		p.pokememeAddDuplicateMessage(update, pokememeID)
		return "ok"
	}

	newPokememeID, err := c.DataCache.AddPokememe(pokememeData, pokememeLocations, pokememeElements)
	if err != nil {
		c.Log.Error(err.Error())
		p.pokememeAddFailureMessage(update)
		return "fail"
	}

	p.pokememeAddSuccessMessage(update, newPokememeID)
	return "ok"
}
