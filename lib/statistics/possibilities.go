// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package statistics

// PossibilityRequiredPokeballs returns possibility of catching pokememe
// It's based on location, grade of pokememe and current level of player
func (s *Statistics) PossibilityRequiredPokeballs(location int, grade int, lvl int) (float64, int) {
	var basePossibility float64
	var requiredPokeballs int
	var percentile float64

	if lvl > 3 {
		switch {
		case grade == (lvl + 1):
			basePossibility = 0.05
		case grade == lvl:
			basePossibility = 0.5
		case grade == (lvl - 1):
			basePossibility = 0.3
		case grade == (lvl - 2):
			basePossibility = 0.1
		case grade == (lvl - 3):
			basePossibility = 0.05
		default:
			basePossibility = 0.00
		}
	} else if lvl == 3 {
		switch grade {
		case 4:
			basePossibility = 0.05
		case 3:
			basePossibility = 0.5
		case 2:
			basePossibility = 0.3
		case 1:
			basePossibility = 0.15
		default:
			basePossibility = 0.00
		}
	} else if lvl == 2 {
		switch grade {
		case 3:
			basePossibility = 0.05
		case 2:
			basePossibility = 0.70
		case 1:
			basePossibility = 0.25
		default:
			basePossibility = 0.00
		}
	} else if lvl == 1 {
		switch grade {
		case 2:
			basePossibility = 0.80
		case 1:
			basePossibility = 0.20
		default:
			basePossibility = 0.00
		}
	}

	pokememesCount := c.DataCache.GetPokememesCountByGradeAndLocation(grade, location)

	if basePossibility != 0 && pokememesCount != 0 {
		percentile = basePossibility * 100.0 / float64(pokememesCount)
		requiredPokeballs = int(100.0 / percentile)
	}

	return percentile, requiredPokeballs
}
