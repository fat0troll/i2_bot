// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package sender

import (
	"github.com/fat0troll/i2_bot/lib/appcontext"
	"github.com/fat0troll/i2_bot/lib/sender/senderinterface"
)

var (
	c *appcontext.Context
)

const (
	parseMode = "Markdown"
)

// Sender is a function-handling struct for sender
type Sender struct{}

// New is a appcontext initialization function
func New(ac *appcontext.Context) {
	c = ac
	s := &Sender{}
	c.RegisterSenderInterface(senderinterface.SenderInterface(s))
}

// Init is an initialization function for sender
func (s *Sender) Init() {
	c.Log.Info("Initializing Sender...")
}
