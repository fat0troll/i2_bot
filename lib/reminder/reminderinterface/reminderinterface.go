// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package reminderinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
)

// ReminderInterface implements Reminder for importing via appcontext
type ReminderInterface interface {
	Init()

	AlarmsList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	CreateAlarmSetting(update *tgbotapi.Update, playerRaw *dbmapping.Player) string
	DestroyAlarmSetting(update *tgbotapi.Update, playerRaw *dbmapping.Player) string

	SendReminders()
}
