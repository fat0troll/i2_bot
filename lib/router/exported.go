// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"lab.pztrn.name/fat0troll/i2_bot/lib/appcontext"
)

var (
	c *appcontext.Context
)

// Router is a function-handling struct for router
type Router struct{}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	r := &Router{}
	c.RegisterRouterInterface(r)
}

// Init is an initialization function for package router
func (r *Router) Init() {
	c.Log.Info("Initialized request router...")
}
