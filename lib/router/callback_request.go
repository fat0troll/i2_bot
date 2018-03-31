// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"regexp"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// RouteCallback routes inline requests to bot
func (r *Router) RouteCallback(update tgbotapi.Update) string {
	playerRaw, err := c.DataCache.GetOrCreatePlayerByTelegramID(update.CallbackQuery.From.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	if playerRaw.Status == "banned" {
		return c.Talkers.BanError(&update)
	}

	var enableAlarmCallback = regexp.MustCompile("enable_reminder_(\\d+)\\z")
	var disableAlarmCallback = regexp.MustCompile("disable_reminder_(\\d+)\\z")

	switch {
	case enableAlarmCallback.MatchString(update.CallbackQuery.Data):
		return c.Reminder.CreateAlarmSetting(&update, playerRaw)
	case disableAlarmCallback.MatchString(update.CallbackQuery.Data):
		return c.Reminder.DestroyAlarmSetting(&update, playerRaw)
	}

	return "ok"
}
