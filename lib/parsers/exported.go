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

type Parsers struct{}

func New(ac *appcontext.Context) {
	c = ac
	p := &Parsers{}
	c.RegisterParsersInterface(parsersinterface.ParsersInterface(p))
}
