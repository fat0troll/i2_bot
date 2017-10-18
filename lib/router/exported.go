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

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	rh := RouterHandler{}
	c.RegisterRouterInterface(rh)
}

// Init is an initialization function for package router
func (r *Router) Init() {
	log.Printf("Initialized request router...")
}
