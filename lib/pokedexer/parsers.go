// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

// ParsePokememe parses pokememe, forwarded from PokememeBroBot, to database
func (p *Pokedexer) ParsePokememe(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	c.Log.Info("Starting pokememe parsing...")
	pokememeStringsArray := strings.Split(update.Message.Text, "\n")
	pokememeRunesArray := make([][]rune, 0)
	for i := range pokememeStringsArray {
		pokememeRunesArray = append(pokememeRunesArray, []rune(pokememeStringsArray[i]))
	}

	pokememeData := make(map[string]string)

	for i := range pokememeStringsArray {
		infoString := pokememeStringsArray[i]
		c.Log.Debug("Processing string: " + infoString)
		if strings.HasPrefix(infoString, "Elements ") {
			infoString = strings.Replace(infoString, "Elements  ", "", -1)
			elements := strings.Split(infoString, " ")
			elementsIDs := make([]string, 0)
			for ii := range elements {
				element, _ := c.DataCache.FindElementIDBySymbol(elements[ii])
				elementsIDs = append(elementsIDs, strconv.Itoa(element))
			}
			pokememeData["elements"] = "\\[" + strings.Join(elementsIDs, ", ") + "]"
		} else if strings.HasPrefix(infoString, "Place ") {
			infoString = strings.Replace(infoString, "Place ", "", -1)
			places := strings.Split(infoString, ",")
			locationIDs := make([]string, 0)
			for ii := range places {
				locationID, _ := c.DataCache.FindLocationIDByName(places[ii])
				locationIDs = append(locationIDs, strconv.Itoa(locationID))
			}
			pokememeData["locations"] = "\\[" + strings.Join(locationIDs, ", ") + "]"
		} else if strings.HasPrefix(infoString, "Buyable ") {
			pokememeData["purchaseable"] = "false"
			if strings.HasSuffix(infoString, "Yes") {
				pokememeData["purchaseable"] = "true"
			}
		} else {
			pokememeData[strings.Split(infoString, " ")[0]] = strings.Join(strings.Split(infoString, " ")[1:], " ")
		}
	}

	c.Log.Debugln("Pokememe data: ", pokememeData)

	message := "- id: " + pokememeData["Dex"]
	message += "\n  grade: " + pokememeData["Grade"]
	message += "\n  name: \"" + pokememeData["Name"] + "\""
	message += "\n  description: \"" + pokememeData["Description"] + "\""
	message += "\n  attack: " + pokememeData["Attack"]
	message += "\n  defence: " + pokememeData["Def"]
	message += "\n  health: " + pokememeData["HP"]
	message += "\n  mana: " + pokememeData["MP"]
	message += "\n  cost: " + pokememeData["Cost"]
	message += "\n  purchaseable: " + pokememeData["purchaseable"]
	message += "\n  image: \"" + pokememeData["Image"] + "\""
	message += "\n  elements: " + pokememeData["elements"]
	message += "\n  locations: " + pokememeData["locations"]

	c.Sender.SendMarkdownAnswer(update, message)

	return "ok"
}
