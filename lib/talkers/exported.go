// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/appcontext"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/talkers/talkersinterface"
)

var (
	c *appcontext.Context
)

// Talkers is a function-handling struct for talkers
type Talkers struct{}

// New is a appcontext initialization function
func New(ac *appcontext.Context) {
	c = ac
	m := &Talkers{}
	c.RegisterTalkersInterface(talkersinterface.TalkersInterface(m))
}

// Init is an initialization function for talkers
func (t *Talkers) Init() {
	c.Log.Info("Initializing common Responders...")

	c.Cron.AddFunc("0 0 0 1 1 *", t.NewYearMessage2018)
}
