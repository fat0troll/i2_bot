// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package welcomer

import (
	// stdlib
	"log"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/appcontext"
	"lab.pztrn.name/fat0troll/i2_bot/lib/welcomer/welcomerinterface"
)

var (
	c *appcontext.Context
)

// Welcomer is a function-handling struct for talkers
type Welcomer struct{}

// New is a appcontext initialization function
func New(ac *appcontext.Context) {
	c = ac
	m := &Welcomer{}
	c.RegisterWelcomerInterface(welcomerinterface.WelcomerInterface(m))
}

// Init is an initialization function for talkers
func (w *Welcomer) Init() {
	log.Printf("Initializing Welcomer...")
}
