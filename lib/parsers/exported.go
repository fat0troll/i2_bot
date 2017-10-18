// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package parsers

import (
	// local
	"../appcontext"
	"../parsers/parsersinterface"
)

var (
	c *appcontext.Context
)

// Parsers is a function-handling struct for package parsers
type Parsers struct{}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	p := &Parsers{}
	c.RegisterParsersInterface(parsersinterface.ParsersInterface(p))
}
