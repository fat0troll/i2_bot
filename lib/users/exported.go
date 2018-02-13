// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package users

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/appcontext"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/users/usersinterface"
)

var (
	c *appcontext.Context
)

// Users is a function-handling struct for users
type Users struct{}

// New is a appcontext initialization function
func New(ac *appcontext.Context) {
	c = ac
	u := &Users{}
	c.RegisterUsersInterface(usersinterface.UsersInterface(u))
}

// Init is an initialization function for users
func (u *Users) Init() {
	c.Log.Info("Initializing Users...")
}
