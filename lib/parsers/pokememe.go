// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsers

import (
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Internal functions

func (p *Parsers) getPoints(pointsStr string) int {
	value := 0
	if strings.HasSuffix(pointsStr, "K") {
		valueNumber := strings.Replace(pointsStr, "K", "", 1)
		valueFloat, _ := strconv.ParseFloat(valueNumber, 64)
		value = int(valueFloat * 1000)
	} else if strings.HasSuffix(pointsStr, "M") {
		valueNumber := strings.Replace(pointsStr, "M", "", 1)
		valueFloat, _ := strconv.ParseFloat(valueNumber, 64)
		value = int(valueFloat * 1000000)
	} else {
		value, _ = strconv.Atoi(pointsStr)
	}
	return value
}

// External functions

// ParsePokememe parses pokememe, forwarded from PokememeBroBot, to database
func (p *Parsers) ParsePokememe(text string, playerRaw *dbmapping.Player) string {
	var defendablePokememe = false
	pokememeStringsArray := strings.Split(text, "\n")
	pokememeRunesArray := make([][]rune, 0)
	for i := range pokememeStringsArray {
		pokememeRunesArray = append(pokememeRunesArray, []rune(pokememeStringsArray[i]))
	}

	if len(pokememeRunesArray) == 13 {
		defendablePokememe = true
	}

	// Getting elements
	elements := []dbmapping.Element{}
	elementEmojis := make([]string, 0)
	elementEmojis = append(elementEmojis, string(pokememeRunesArray[4][11]))
	if len(pokememeRunesArray[4]) > 12 {
		elementEmojis = append(elementEmojis, string(pokememeRunesArray[4][13]))
	}
	if len(pokememeRunesArray[4]) > 14 {
		elementEmojis = append(elementEmojis, string(pokememeRunesArray[4][15]))
	}

	err := c.Db.Select(&elements, "SELECT * FROM elements WHERE symbol IN ('"+strings.Join(elementEmojis, "', '")+"')")
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	// Getting hit-points
	hitPointsRx := regexp.MustCompile("(\\d|\\.)+(K|M)?")
	hitPoints := hitPointsRx.FindAllString(string(pokememeRunesArray[5]), -1)
	if len(hitPoints) != 3 {
		c.Log.Error("Can't parse hitpoints!")
		c.Log.Debug(pokememeRunesArray[5])
		return "fail"
	}

	defence := "0"
	price := "0"

	locations := []dbmapping.Location{}

	purchaseable := false
	image := ""

	if defendablePokememe {
		// Actions for high-grade pokememes
		defenceMatch := hitPointsRx.FindAllString(string(pokememeRunesArray[6]), -1)
		if len(defenceMatch) < 1 {
			c.Log.Error("Can't parse defence!")
			c.Log.Debug(pokememeRunesArray[6])
			return "fail"
		}
		defence = defenceMatch[0]
		priceMatch := hitPointsRx.FindAllString(string(pokememeRunesArray[7]), -1)
		if len(priceMatch) < 1 {
			c.Log.Error("Can't parse price!")
			c.Log.Debug(pokememeRunesArray[7])
			return "fail"
		}
		price = priceMatch[0]
		locationsPrepare := strings.Split(string(pokememeRunesArray[8]), ": ")
		if len(locationsPrepare) < 2 {
			c.Log.Error("Can't parse locations!")
			c.Log.Debug(pokememeRunesArray[8])
			return "fail"
		}
		locationsNames := strings.Split(locationsPrepare[1], ", ")
		if len(locationsNames) < 1 {
			c.Log.Error("Can't parse locations!")
			c.Log.Debug(locationsPrepare)
			return "fail"
		}

		err2 := c.Db.Select(&locations, "SELECT * FROM locations WHERE name IN ('"+strings.Join(locationsNames, "', '")+"')")
		if err2 != nil {
			c.Log.Error(err2.Error())
			return "fail"
		}
		if strings.HasSuffix(string(pokememeRunesArray[9]), "Можно") {
			purchaseable = true
		}
		image = strings.Replace(string(pokememeRunesArray[12]), " ", "", -1)
	} else {
		// Actions for low-grade pokememes
		defence = hitPoints[0]
		priceMatch := hitPointsRx.FindAllString(string(pokememeRunesArray[6]), -1)
		if len(priceMatch) < 1 {
			c.Log.Error("Can't parse price!")
			c.Log.Debug(pokememeRunesArray[6])
			return "fail"
		}
		price = priceMatch[0]
		locationsPrepare := strings.Split(string(pokememeRunesArray[7]), ": ")
		if len(locationsPrepare) < 2 {
			c.Log.Error("Can't parse locations!")
			c.Log.Debug(pokememeRunesArray[7])
			return "fail"
		}
		locationsNames := strings.Split(locationsPrepare[1], ", ")
		if len(locationsNames) < 1 {
			c.Log.Error("Can't parse locations!")
			c.Log.Debug(locationsPrepare)
			return "fail"
		}

		err2 := c.Db.Select(&locations, "SELECT * FROM locations WHERE name IN ('"+strings.Join(locationsNames, "', '")+"')")
		if err2 != nil {
			c.Log.Error(err2.Error())
			return "fail"
		}
		if strings.HasSuffix(string(pokememeRunesArray[8]), "Можно") {
			purchaseable = true
		}
		image = strings.Replace(string(pokememeRunesArray[11]), " ", "", -1)
	}

	grade := string(pokememeRunesArray[0][0])
	name := string(pokememeRunesArray[0][3:])
	description := string(pokememeRunesArray[1])
	c.Log.Debug("Pokememe grade: " + grade)
	c.Log.Debug("Pokememe name: " + name)
	c.Log.Debug("Pokememe description: " + description)
	c.Log.Debug("Elements:")
	for i := range elements {
		c.Log.Debug(elements[i].Symbol + " " + elements[i].Name)
	}
	c.Log.Debug("Attack: " + hitPoints[0])
	c.Log.Debug("HP: " + hitPoints[1])
	c.Log.Debug("MP: " + hitPoints[2])
	c.Log.Debug("Defence: " + defence)
	c.Log.Debug("Price: " + price)
	c.Log.Debug("Locations:")
	for i := range locations {
		c.Log.Debug(locations[i].Symbol + " " + locations[i].Name)
	}
	if purchaseable {
		c.Log.Debug("Purchaseable")
	} else {
		c.Log.Debug("Non-purchaseable")
	}
	c.Log.Debug("Image: " + image)

	// Building pokememe
	pokememe := dbmapping.Pokememe{}
	// Checking if pokememe exists in database
	err3 := c.Db.Get(&pokememe, c.Db.Rebind("SELECT * FROM pokememes WHERE grade='"+grade+"' AND name='"+name+"';"))
	if err3 != nil {
		c.Log.Debug("Adding new pokememe...")
	} else {
		c.Log.Info("This pokememe already exist. Return specific error.")
		return "dup"
	}

	gradeInt, _ := strconv.Atoi(grade)
	attackInt := p.getPoints(hitPoints[0])
	hpInt := p.getPoints(hitPoints[1])
	mpInt := p.getPoints(hitPoints[2])
	defenceInt := p.getPoints(defence)
	priceInt := p.getPoints(price)

	pokememe.Grade = gradeInt
	pokememe.Name = name
	pokememe.Description = description
	pokememe.Attack = attackInt
	pokememe.HP = hpInt
	pokememe.MP = mpInt
	pokememe.Defence = defenceInt
	pokememe.Price = priceInt
	if purchaseable {
		pokememe.Purchaseable = true
	} else {
		pokememe.Purchaseable = false
	}
	pokememe.ImageURL = image
	pokememe.PlayerID = playerRaw.ID
	pokememe.CreatedAt = time.Now().UTC()

	_, err4 := c.Db.NamedExec("INSERT INTO pokememes VALUES(NULL, :grade, :name, :description, :attack, :hp, :mp, :defence, :price, :purchaseable, :image_url, :player_id, :created_at)", &pokememe)
	if err4 != nil {
		c.Log.Error(err4.Error())
		return "fail"
	}

	// Getting new pokememe
	err5 := c.Db.Get(&pokememe, c.Db.Rebind("SELECT * FROM pokememes WHERE grade='"+grade+"' AND name='"+name+"';"))
	if err5 != nil {
		c.Log.Error("Pokememe isn't added!")
		return "fail"
	}
	for i := range elements {
		link := dbmapping.PokememeElement{}
		link.PokememeID = pokememe.ID
		link.ElementID = elements[i].ID
		link.CreatedAt = time.Now().UTC()

		_, err6 := c.Db.NamedExec("INSERT INTO pokememes_elements VALUES(NULL, :pokememe_id, :element_id, :created_at)", &link)
		if err6 != nil {
			c.Log.Error(err6.Error())
			return "fail"
		}
	}
	for i := range locations {
		link := dbmapping.PokememeLocation{}
		link.PokememeID = pokememe.ID
		link.LocationID = locations[i].ID
		link.CreatedAt = time.Now().UTC()

		_, err7 := c.Db.NamedExec("INSERT INTO pokememes_locations VALUES(NULL, :pokememe_id, :location_id, :created_at)", &link)
		if err7 != nil {
			c.Log.Error(err7.Error())
			return "fail"
		}
	}

	return "ok"
}

// ReturnPoints returns to output points (ht, attack, mp...) formatted
// like in PokememBroBot itself.
func (p *Parsers) ReturnPoints(points int) string {
	if points < 1000 {
		return strconv.Itoa(points)
	} else if points < 1000000 {
		floatNum := float64(points) / 1000.0
		return strconv.FormatFloat(floatNum, 'f', -1, 64) + "K"
	} else {
		floatNum := float64(points) / 1000000.0
		return strconv.FormatFloat(floatNum, 'f', -1, 64) + "M"
	}
}
