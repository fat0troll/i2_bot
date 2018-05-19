// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package statistics

import (
	"github.com/fat0troll/i2_bot/lib/appcontext"
	"github.com/fat0troll/i2_bot/lib/statistics/statisticsinterface"
)

var (
	c *appcontext.Context
)

// Statistics is a function-handling struct for package statistics.
type Statistics struct{}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	s := &Statistics{}
	c.RegisterStatisticsInterface(statisticsinterface.StatisticsInterface(s))
}

// Init is a initialization function for package
func (s *Statistics) Init() {
	c.Log.Info("Initializing Statistics...")
}
