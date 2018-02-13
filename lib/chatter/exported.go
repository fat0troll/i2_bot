// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package chatter

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/appcontext"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/chatter/chatterinterface"
)

var (
	c *appcontext.Context
)

// Chatter is a function-handling struct for package chatter.
type Chatter struct{}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	ct := &Chatter{}
	c.RegisterChatterInterface(chatterinterface.ChatterInterface(ct))
}

// Init is a initialization function for package
func (ct *Chatter) Init() {
	c.Log.Info("Initializing Chatter...")
}
