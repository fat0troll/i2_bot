// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package appcontext

var (
	a *Context
)

// New is a Context creation function
func New() *Context {
	c := &Context{}
	return c
}
