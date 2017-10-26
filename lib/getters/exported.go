// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package getters

import (
	// stdlib
	"log"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/appcontext"
	"lab.pztrn.name/fat0troll/i2_bot/lib/getters/gettersinterface"
)

var (
	c *appcontext.Context
)

// Getters is a function-handling struct for package getters.
type Getters struct{}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	g := &Getters{}
	c.RegisterGettersInterface(gettersinterface.GettersInterface(g))
}

// Init is a initialization function for package
func (g *Getters) Init() {
	log.Printf("Initializing getters...")
}
