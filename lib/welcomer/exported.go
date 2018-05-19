// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomer

import (
	"github.com/fat0troll/i2_bot/lib/appcontext"
	"github.com/fat0troll/i2_bot/lib/welcomer/welcomerinterface"
)

var (
	c *appcontext.Context
)

// Welcomer is a function-handling struct for welcomer
type Welcomer struct{}

// New is a appcontext initialization function
func New(ac *appcontext.Context) {
	c = ac
	m := &Welcomer{}
	c.RegisterWelcomerInterface(welcomerinterface.WelcomerInterface(m))
}

// Init is an initialization function for welcomer
func (w *Welcomer) Init() {
	c.Log.Info("Initializing Welcomer...")
}
