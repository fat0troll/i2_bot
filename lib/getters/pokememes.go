// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package getters

import (
    // stdlib
    "log"
    "strconv"
    // local
    "../dbmapping"
)

// Internal functions

func (g *Getters) formFullPokememes(pokememes []dbmapping.Pokememe) ([]dbmapping.PokememeFull, bool) {
    pokememes_full := []dbmapping.PokememeFull{}
    elements := []dbmapping.Element{}
    err := c.Db.Select(&elements, "SELECT * FROM elements");
    if err != nil {
        log.Println(err)
        return pokememes_full, false
    }
    locations := []dbmapping.Location{}
    err = c.Db.Select(&locations, "SELECT * FROM locations");
    if err != nil {
        log.Println(err)
        return pokememes_full, false
    }
    pokememes_elements := []dbmapping.PokememeElement{}
    err = c.Db.Select(&pokememes_elements, "SELECT * FROM pokememes_elements");
    if err != nil {
        log.Println(err)
        return pokememes_full, false
    }
    pokememes_locations := []dbmapping.PokememeLocation{}
    err = c.Db.Select(&pokememes_locations, "SELECT * FROM pokememes_locations");
    if err != nil {
        log.Println(err)
        return pokememes_full, false
    }

    for i := range(pokememes) {
        full_pokememe := dbmapping.PokememeFull{}
        elements_listed := []dbmapping.Element{}
        locations_listed := []dbmapping.Location{}

        for j := range(pokememes_locations) {
            if pokememes_locations[j].Pokememe_id == pokememes[i].Id {
                for l := range(locations) {
                    if pokememes_locations[j].Location_id == locations[l].Id {
                        locations_listed = append(locations_listed, locations[l])
                    }
                }
            }
        }

        for k := range(pokememes_elements) {
            if pokememes_elements[k].Pokememe_id == pokememes[i].Id {
                for e := range(elements) {
                    if pokememes_elements[k].Element_id == elements[e].Id {
                        elements_listed = append(elements_listed, elements[e])
                    }
                }
            }
        }

        full_pokememe.Pokememe = pokememes[i]
        full_pokememe.Elements = elements_listed
        full_pokememe.Locations = locations_listed

        pokememes_full = append(pokememes_full, full_pokememe)
    }

    return pokememes_full, true
}

// External functions

func (g *Getters) GetPokememes() ([]dbmapping.PokememeFull, bool) {
    pokememes_full := []dbmapping.PokememeFull{}
    pokememes := []dbmapping.Pokememe{}
    err := c.Db.Select(&pokememes, "SELECT * FROM pokememes ORDER BY grade asc, name asc");
    if err != nil {
        log.Println(err)
        return pokememes_full, false
    }

    pokememes_full, ok := g.formFullPokememes(pokememes)
    return pokememes_full, ok
}

func (g *Getters) GetBestPokememes(player_id int) ([]dbmapping.PokememeFull, bool) {
    pokememes_full := []dbmapping.PokememeFull{}
    player_raw, ok := g.GetPlayerByID(player_id)
    if !ok {
        return pokememes_full, ok
    }
    profile_raw, ok := g.GetProfile(player_id)
    if !ok {
        return pokememes_full, ok
    }

    if player_raw.League_id == 0 {
        return pokememes_full, false
    }

    // TODO: make it more complicated
    pokememes := []dbmapping.Pokememe{}
    err := c.Db.Select(&pokememes, c.Db.Rebind("SELECT p.* FROM pokememes p, pokememes_elements pe, elements e WHERE e.league_id = ? AND p.grade = ? AND pe.element_id = e.id AND pe.pokememe_id = p.id ORDER BY p.attack DESC"), player_raw.League_id, profile_raw.Level_id + 1)
    if err != nil {
        log.Println(err)
        return pokememes_full, false
    }

    pokememes_full, ok = g.formFullPokememes(pokememes)
    return pokememes_full, ok
}

func (g *Getters) GetPokememeByID(pokememe_id string) (dbmapping.PokememeFull, bool) {
    pokememe_full := dbmapping.PokememeFull{}
    pokememe := dbmapping.Pokememe{}
    err := c.Db.Get(&pokememe, c.Db.Rebind("SELECT * FROM pokememes WHERE id=?"), pokememe_id)
    if err != nil {
        log.Println(err)
        return pokememe_full, false
    }
    elements := []dbmapping.Element{}
    err = c.Db.Select(&elements, "SELECT * FROM elements");
    if err != nil {
        log.Println(err)
        return pokememe_full, false
    }
    locations := []dbmapping.Location{}
    err = c.Db.Select(&locations, "SELECT * FROM locations");
    if err != nil {
        log.Println(err)
        return pokememe_full, false
    }
    pokememes_elements := []dbmapping.PokememeElement{}
    err = c.Db.Select(&pokememes_elements, "SELECT * FROM pokememes_elements WHERE pokememe_id='" + strconv.Itoa(pokememe.Id) + "'");
    if err != nil {
        log.Println(err)
        return pokememe_full, false
    }
    pokememes_locations := []dbmapping.PokememeLocation{}
    err = c.Db.Select(&pokememes_locations, "SELECT * FROM pokememes_locations WHERE pokememe_id='" + strconv.Itoa(pokememe.Id) + "'");
    if err != nil {
        log.Println(err)
        return pokememe_full, false
    }

    elements_listed := []dbmapping.Element{}
    locations_listed := []dbmapping.Location{}

    for j := range(pokememes_locations) {
        if pokememes_locations[j].Pokememe_id == pokememe.Id {
            for l := range(locations) {
                if pokememes_locations[j].Location_id == locations[l].Id {
                    locations_listed = append(locations_listed, locations[l])
                }
            }
        }
    }

    for k := range(pokememes_elements) {
        if pokememes_elements[k].Pokememe_id == pokememe.Id {
            for e := range(elements) {
                if pokememes_elements[k].Element_id == elements[e].Id {
                    elements_listed = append(elements_listed, elements[e])
                }
            }
        }
    }

    pokememe_full.Pokememe = pokememe
    pokememe_full.Elements = elements_listed
    pokememe_full.Locations = locations_listed

    return pokememe_full, true
}
