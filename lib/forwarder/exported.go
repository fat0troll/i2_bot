// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package forwarder

import (
	"lab.pztrn.name/fat0troll/i2_bot/lib/appcontext"
	"lab.pztrn.name/fat0troll/i2_bot/lib/forwarder/forwarderinterface"
)

var (
	c *appcontext.Context
)

// Forwarder is a function-handling struct for package forwarder.
type Forwarder struct{}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	f := &Forwarder{}
	c.RegisterForwarderInterface(forwarderinterface.ForwarderInterface(f))
}

// Init is a initialization function for package
func (f *Forwarder) Init() {
	c.Log.Info("Initializing forwarder...")
}
