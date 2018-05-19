// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package reminder

import (
	"github.com/fat0troll/i2_bot/lib/appcontext"
	"github.com/fat0troll/i2_bot/lib/reminder/reminderinterface"
)

var (
	c *appcontext.Context
)

// Reminder is a function-handling struct for Reminder
type Reminder struct{}

// New is a appcontext initialization function
func New(ac *appcontext.Context) {
	c = ac
	r := &Reminder{}
	c.RegisterReminderInterface(reminderinterface.ReminderInterface(r))
}

// Init is an initialization function for reminder
func (r *Reminder) Init() {
	c.Log.Info("Initializing Reminder...")

	c.Cron.AddFunc("0 55 0-23/2 * * *", r.SendReminders)
}
