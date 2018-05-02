// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package statisticsinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

// StatisticsInterface implements Statistics for importing via appcontext.
type StatisticsInterface interface {
	Init()

	SquadStatictics(squadID int) string

	GetPoints(pointsStr string) int
	GetPrintablePoints(points int) string

	PossibilityRequiredPokeballs(location int, grade int, lvl int) (float64, int)

	TopList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
}
