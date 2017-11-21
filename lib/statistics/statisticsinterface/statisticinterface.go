// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package statisticsinterface

// StatisticsInterface implements Statistics for importing via appcontext.
type StatisticsInterface interface {
	Init()

	GetPoints(pointsStr string) int
	GetPrintablePoints(points int) string

	PossibilityRequiredPokeballs(location int, grade int, lvl int) (float64, int)
}
