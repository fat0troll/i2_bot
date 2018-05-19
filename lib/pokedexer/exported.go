// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pokedexer

import (
	// local
	"github.com/fat0troll/i2_bot/lib/appcontext"
	"github.com/fat0troll/i2_bot/lib/pokedexer/pokedexerinterface"
)

var (
	c *appcontext.Context
)

// Pokedexer is a function-handling struct for package pokedexer
type Pokedexer struct{}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	p := &Pokedexer{}
	c.RegisterPokedexerInterface(pokedexerinterface.PokedexerInterface(p))
}
