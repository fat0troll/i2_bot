// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pinner

import (
	"github.com/fat0troll/i2_bot/lib/appcontext"
	"github.com/fat0troll/i2_bot/lib/pinner/pinnerinterface"
)

var (
	c *appcontext.Context
)

// Pinner is a function-handling struct for Pinner
type Pinner struct{}

// New is a appcontext initialization function
func New(ac *appcontext.Context) {
	c = ac
	p := &Pinner{}
	c.RegisterPinnerInterface(pinnerinterface.PinnerInterface(p))
}

// Init is an initialization function for pinner
func (p *Pinner) Init() {
	c.Log.Info("Initializing Pinner...")
}
