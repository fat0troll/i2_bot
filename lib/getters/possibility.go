// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package getters

import (
	// stdlib
	"log"
)

func (g *Getters) PossibilityRequiredPokeballs(location int, grade int, lvl int) (float64, int) {
	var base_possibility float64 = 0.00
	var required_pokeballs int = 0
	var percentile = 0.00

	if lvl > 3 {
		switch {
		case grade == (lvl + 1):
			base_possibility = 0.05
		case grade == lvl:
			base_possibility = 0.5
		case grade == (lvl - 1):
			base_possibility = 0.3
		case grade == (lvl - 2):
			base_possibility = 0.1
		case grade == (lvl - 3):
			base_possibility = 0.05
		default:
			base_possibility = 0.00
		}
	} else if lvl == 3 {
		switch grade {
		case 4:
			base_possibility = 0.05
		case 3:
			base_possibility = 0.5
		case 2:
			base_possibility = 0.3
		case 1:
			base_possibility = 0.15
		default:
			base_possibility = 0.00
		}
	} else if lvl == 2 {
		switch grade {
		case 3:
			base_possibility = 0.05
		case 2:
			base_possibility = 0.70
		case 1:
			base_possibility = 0.25
		default:
			base_possibility = 0.00
		}
	} else if lvl == 1 {
		switch grade {
		case 2:
			base_possibility = 0.80
		case 1:
			base_possibility = 0.20
		default:
			base_possibility = 0.00
		}
	}

	var number_of_pokememes int = 0

	err := c.Db.Get(&number_of_pokememes, c.Db.Rebind("SELECT count(*) FROM pokememes p, pokememes_locations pl WHERE p.grade = ? AND pl.location_id = ? AND pl.pokememe_id = p.id;"), grade, location)
	if err != nil {
		log.Println(err)
	}

	if base_possibility != 0 && number_of_pokememes != 0 {
		percentile = base_possibility * 100.0 / float64(number_of_pokememes)
		required_pokeballs = int(100.0 / percentile)
	}

	return percentile, required_pokeballs
}
