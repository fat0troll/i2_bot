// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package getters

import (
	// stdlib
	"log"
	// local
	"../appcontext"
	"../getters/gettersinterface"
)

var (
	c *appcontext.Context
)

type Getters struct{}

func New(ac *appcontext.Context) {
	c = ac
	g := &Getters{}
	c.RegisterGettersInterface(gettersinterface.GettersInterface(g))
}

func (g *Getters) Init() {
	log.Printf("Initializing getters...")
}
