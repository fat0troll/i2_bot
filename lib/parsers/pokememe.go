// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsers

import (
    // stdlib
    "log"
    "regexp"
    "strings"
    "strconv"
    "time"
    // local
    "../dbmappings"
)

// Internal functions

func (p *Parsers) getPoints(points_str string) int {
    value := 0
    if strings.HasSuffix(points_str, "K") {
        value_num := strings.Replace(points_str, "K", "", 1)
        value_float, _ := strconv.ParseFloat(value_num, 64)
        value = int(value_float * 1000)
    } else if strings.HasSuffix(points_str, "M") {
        value_num := strings.Replace(points_str, "M", "", 1)
        value_float, _ := strconv.ParseFloat(value_num, 64)
        value = int(value_float * 1000000)
    } else {
        value, _ = strconv.Atoi(points_str)
    }
    return value
}

// External functions

func (p *Parsers) ParsePokememe(text string, player_raw dbmappings.Players) string {
    var defendable_pokememe bool = false
    pokememe_info_strings := strings.Split(text, "\n")
    pokememe_info_runed_strings := make([][]rune, 0)
    for i := range(pokememe_info_strings) {
        pokememe_info_runed_strings = append(pokememe_info_runed_strings, []rune(pokememe_info_strings[i]))
    }

    if len(pokememe_info_runed_strings) == 13 {
        defendable_pokememe = true
    }

    // Getting elements
    elements := []dbmappings.Elements{}
    element_emojis := make([]string, 0)
    element_emojis = append(element_emojis, string(pokememe_info_runed_strings[4][11]))
    if len(pokememe_info_runed_strings[4]) > 12 {
        element_emojis = append(element_emojis, string(pokememe_info_runed_strings[4][13]))
    }
    if len(pokememe_info_runed_strings[4]) > 14 {
        element_emojis = append(element_emojis, string(pokememe_info_runed_strings[4][15]))
    }

    err := c.Db.Select(&elements, "SELECT * FROM elements WHERE symbol IN ('" + strings.Join(element_emojis, "', '") + "')")
    if err != nil {
        log.Printf(err.Error())
        return "fail"
    }

    // Getting hit-points
    hitPointsRx := regexp.MustCompile("(\\d|\\.)+(K|M)?")
    hitPoints := hitPointsRx.FindAllString(string(pokememe_info_runed_strings[5]), -1)
    if len(hitPoints) != 3 {
        log.Printf("Can't parse hitpoints!")
        log.Println(pokememe_info_runed_strings[5])
        return "fail"
    }

    defence := "0"
    price := "0"

    locations := []dbmappings.Locations{}

    purchaseable := false
    image := ""

    if defendable_pokememe {
        // Actions for high-grade pokememes
        defenceMatch := hitPointsRx.FindAllString(string(pokememe_info_runed_strings[6]), -1)
        if len(defenceMatch) < 1 {
            log.Printf("Can't parse defence!")
            log.Println(pokememe_info_runed_strings[6])
            return "fail"
        }
        defence = defenceMatch[0]
        priceMatch := hitPointsRx.FindAllString(string(pokememe_info_runed_strings[7]), -1)
        if len(priceMatch) < 1 {
            log.Printf("Can't parse price!")
            log.Println(pokememe_info_runed_strings[7])
            return "fail"
        }
        price = priceMatch[0]
        locationsPrepare := strings.Split(string(pokememe_info_runed_strings[8]), ": ")
        if len(locationsPrepare) < 2 {
            log.Printf("Can't parse locations!")
            log.Println(pokememe_info_runed_strings[8])
            return "fail"
        }
        locationsNames := strings.Split(locationsPrepare[1], ", ")
        if len(locationsNames) < 1 {
            log.Printf("Can't parse locations!")
            log.Println(locationsPrepare)
            return "fail"
        }

        err2 := c.Db.Select(&locations, "SELECT * FROM locations WHERE name IN ('" + strings.Join(locationsNames, "', '") + "')")
        if err2 != nil {
            log.Printf(err2.Error())
            return "fail"
        }
        if strings.HasSuffix(string(pokememe_info_runed_strings[9]), "Можно") {
            purchaseable = true
        }
        image = strings.Replace(string(pokememe_info_runed_strings[12]), " ", "", -1)
    } else {
        // Actions for low-grade pokememes
        defence = hitPoints[0]
        priceMatch := hitPointsRx.FindAllString(string(pokememe_info_runed_strings[6]), -1)
        if len(priceMatch) < 1 {
            log.Printf("Can't parse price!")
            log.Println(pokememe_info_runed_strings[6])
            return "fail"
        }
        price = priceMatch[0]
        locationsPrepare := strings.Split(string(pokememe_info_runed_strings[7]), ": ")
        if len(locationsPrepare) < 2 {
            log.Printf("Can't parse locations!")
            log.Println(pokememe_info_runed_strings[7])
            return "fail"
        }
        locationsNames := strings.Split(locationsPrepare[1], ", ")
        if len(locationsNames) < 1 {
            log.Printf("Can't parse locations!")
            log.Println(locationsPrepare)
            return "fail"
        }

        err2 := c.Db.Select(&locations, "SELECT * FROM locations WHERE name IN ('" + strings.Join(locationsNames, "', '") + "')")
        if err2 != nil {
            log.Printf(err2.Error())
            return "fail"
        }
        if strings.HasSuffix(string(pokememe_info_runed_strings[8]), "Можно") {
            purchaseable = true
        }
        image = strings.Replace(string(pokememe_info_runed_strings[11]), " ", "", -1)
    }


    grade := string(pokememe_info_runed_strings[0][0])
    name := string(pokememe_info_runed_strings[0][3:])
    description := string(pokememe_info_runed_strings[1])
    log.Printf("Pokememe grade: " + grade)
    log.Printf("Pokememe name: " + name)
    log.Printf("Pokememe description: " + description)
    log.Printf("Elements:")
    for i := range(elements) {
        log.Printf(elements[i].Symbol + " " + elements[i].Name)
    }
    log.Printf("Attack: " + hitPoints[0])
    log.Printf("HP: " + hitPoints[1])
    log.Printf("MP: " + hitPoints[2])
    log.Printf("Defence: " + defence)
    log.Printf("Price: " + price)
    log.Printf("Locations:")
    for i := range(locations) {
        log.Printf(locations[i].Symbol + " " + locations[i].Name)
    }
    if purchaseable {
        log.Printf("Purchaseable")
    } else {
        log.Printf("Non-purchaseable")
    }
    log.Printf("Image: " + image)

    // Building pokememe
    pokememe := dbmappings.Pokememes{}
    // Checking if pokememe exists in database
    err3 := c.Db.Get(&pokememe, c.Db.Rebind("SELECT * FROM pokememes WHERE grade='" + grade + "' AND name='" + name + "';"))
    if err3 != nil {
        log.Printf("Adding new pokememe...")
    } else {
        log.Printf("This pokememe already exist. Return specific error.")
        return "dup"
    }

    grade_int, _ := strconv.Atoi(grade)
    attack_int := p.getPoints(hitPoints[0])
    hp_int := p.getPoints(hitPoints[1])
    mp_int := p.getPoints(hitPoints[2])
    defence_int := p.getPoints(defence)
    price_int := p.getPoints(price)

    pokememe.Grade = grade_int
    pokememe.Name = name
    pokememe.Description = description
    pokememe.Attack = attack_int
    pokememe.HP = hp_int
    pokememe.MP = mp_int
    pokememe.Defence = defence_int
    pokememe.Price = price_int
    if purchaseable {
        pokememe.Purchaseable = true
    } else {
        pokememe.Purchaseable = false
    }
    pokememe.Image_url = image
    pokememe.Player_id = player_raw.Id
    pokememe.Created_at = time.Now().UTC()

    _, err4 := c.Db.NamedExec("INSERT INTO pokememes VALUES(NULL, :grade, :name, :description, :attack, :hp, :mp, :defence, :price, :purchaseable, :image_url, :player_id, :created_at)", &pokememe)
    if err4 != nil {
        log.Printf(err4.Error())
        return "fail"
    }

    // Getting new pokememe
    err5 := c.Db.Get(&pokememe, c.Db.Rebind("SELECT * FROM pokememes WHERE grade='" + grade + "' AND name='" + name + "';"))
    if err5 != nil {
        log.Printf("Pokememe isn't added!")
        return "fail"
    }
    for i := range(elements) {
        link := dbmappings.PokememesElements{}
        link.Pokememe_id = pokememe.Id
        link.Element_id = elements[i].Id
        link.Created_at = time.Now().UTC()

        _, err6 := c.Db.NamedExec("INSERT INTO pokememes_elements VALUES(NULL, :pokememe_id, :element_id, :created_at)", &link)
        if err6 != nil {
            log.Printf(err6.Error())
            return "fail"
        }
    }
    for i := range(locations) {
        link := dbmappings.PokememesLocations{}
        link.Pokememe_id = pokememe.Id
        link.Location_id = locations[i].Id
        link.Created_at = time.Now().UTC()

        _, err7 := c.Db.NamedExec("INSERT INTO pokememes_locations VALUES(NULL, :pokememe_id, :location_id, :created_at)", &link)
        if err7 != nil {
            log.Printf(err7.Error())
            return "fail"
        }
    }


    return "ok"
}

func (p *Parsers) ReturnPoints(points int) string {
    if points < 1000 {
        return strconv.Itoa(points)
    } else if points < 1000000 {
        float_num := float64(points) / 1000.0
        return strconv.FormatFloat(float_num, 'f', -1, 64) + "K"
    } else {
        float_num := float64(points) / 1000000.0
        return strconv.FormatFloat(float_num, 'f', -1, 64) + "M"
    }
}
