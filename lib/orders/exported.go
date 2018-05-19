// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package orders

import (
	"github.com/fat0troll/i2_bot/lib/appcontext"
	"github.com/fat0troll/i2_bot/lib/orders/ordersinterface"
)

var (
	c *appcontext.Context
)

// Orders is a function-handling struct for package orders.
type Orders struct{}

// New is an initialization function for appcontext
func New(ac *appcontext.Context) {
	c = ac
	o := &Orders{}
	c.RegisterOrdersInterface(ordersinterface.OrdersInterface(o))
}

// Init is a initialization function for package
func (o *Orders) Init() {
	c.Log.Info("Initializing Orders...")
}
