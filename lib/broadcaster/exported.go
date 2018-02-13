// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package broadcaster

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/appcontext"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/broadcaster/broadcasterinterface"
)

var (
	c *appcontext.Context
)

// Broadcaster is a function-handling struct for broadcaster
type Broadcaster struct{}

// New is a appcontext initialization function
func New(ac *appcontext.Context) {
	c = ac
	b := &Broadcaster{}
	c.RegisterBroadcasterInterface(broadcasterinterface.BroadcasterInterface(b))
}

// Init is an initialization function for talkers
func (b *Broadcaster) Init() {
	c.Log.Info("Initializing Broadcaster...")
}
