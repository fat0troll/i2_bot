// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package reminder

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

// AlarmsList lists all alarms for user with buttons to enable/disable each of available alarms
func (r *Reminder) AlarmsList(update *tgbotapi.Update, playerRaw *dbmapping.Player) string {
	msg := tgbotapi.MessageConfig{}
	msg.Text = r.formatRemindersMessageText(playerRaw)
	msg.ParseMode = "Markdown"
	msg.ChatID = update.Message.Chat.ID

	remindersMsg, _ := c.Bot.Send(msg)

	keyboard := r.formatRemindersButtons(playerRaw)
	buttonsUpdate := tgbotapi.NewEditMessageReplyMarkup(update.Message.Chat.ID, remindersMsg.MessageID, keyboard)
	c.Bot.Send(buttonsUpdate)

	return "ok"
}
