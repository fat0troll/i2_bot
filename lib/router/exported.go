// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	// stdlib
	"log"
	// local
	"../appcontext"
)

var (
	c *appcontext.Context
	r *Router
)

func New(ac *appcontext.Context) {
	c = ac
	rh := RouterHandler{}
	c.RegisterRouterInterface(rh)
}

func (r *Router) Init() {
	log.Printf("Initialized request router...")
}
