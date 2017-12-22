// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package reminder

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
	"time"
)

// CreateAlarmSetting creates alarm setting for user
func (r *Reminder) CreateAlarmSetting(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	turnirNumber := strings.TrimPrefix(update.CallbackQuery.Data, "enable_reminder_")
	turnirNumberInt, err := strconv.Atoi(turnirNumber)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	alarm := dbmapping.Alarm{}
	alarm.PlayerID = playerRaw.ID
	alarm.TurnirNumber = turnirNumberInt
	alarm.CreatedAt = time.Now().UTC()

	_, err = c.Db.NamedExec("INSERT INTO `alarms` VALUES(NULL, :player_id, :turnir_number, :created_at)", &alarm)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	keyboard := r.formatRemindersButtons(playerRaw)
	buttonsUpdate := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, keyboard)
	c.Bot.Send(buttonsUpdate)

	return "ok"
}

// DestroyAlarmSetting creates alarm setting for user
func (r *Reminder) DestroyAlarmSetting(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	turnirNumber := strings.TrimPrefix(update.CallbackQuery.Data, "disable_reminder_")
	turnirNumberInt, err := strconv.Atoi(turnirNumber)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	_, err = c.Db.Exec(c.Db.Rebind("DELETE FROM `alarms` WHERE player_id=? AND turnir_number=?"), playerRaw.ID, turnirNumberInt)
	if err != nil {
		c.Log.Error(err.Error())
	}

	keyboard := r.formatRemindersButtons(playerRaw)
	buttonsUpdate := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, keyboard)
	c.Bot.Send(buttonsUpdate)

	return "ok"
}
