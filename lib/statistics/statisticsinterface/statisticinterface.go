// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package statisticsinterface

import (
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
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
