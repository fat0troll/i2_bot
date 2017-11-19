// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package squader

import (
	"lab.pztrn.name/fat0troll/i2_bot/lib/appcontext"
	"lab.pztrn.name/fat0troll/i2_bot/lib/squader/squaderinterface"
)

var (
	c *appcontext.Context
)

// Squader is a function-handling struct for package squader.
type Squader struct{}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	s := &Squader{}
	c.RegisterSquaderInterface(squaderinterface.SquaderInterface(s))
}

// Init is a initialization function for package
func (s *Squader) Init() {
	c.Log.Info("Initializing Squader...")
}
