// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsers

import (
    // stdlib
    "fmt"
    "log"
    "strings"
    // local
    "../dbmappings"
)

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
    fmt.Println(defendable_pokememe)

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

    return "ok"
}
