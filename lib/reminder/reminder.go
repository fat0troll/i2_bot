// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package reminder

import (
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

func (r *Reminder) getRemindersForUser(playerRaw *dbmapping.Player) ([]dbmapping.Alarm, bool) {
	alarmsList := []dbmapping.Alarm{}

	err := c.Db.Select(&alarmsList, "SELECT * FROM alarms WHERE player_id=?", playerRaw.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return alarmsList, false
	}

	return alarmsList, true
}

func (r *Reminder) formatRemindersButtons(playerRaw *dbmapping.Player) tgbotapi.InlineKeyboardMarkup {
	currentAlarms, _ := r.getRemindersForUser(playerRaw)

	alarmExist := make(map[string]string)
	for i := range currentAlarms {
		alarmExist[strconv.Itoa(currentAlarms[i].TurnirNumber)] = "enabled"
	}

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	rows := make(map[int][]tgbotapi.InlineKeyboardButton)
	rows[0] = []tgbotapi.InlineKeyboardButton{}
	rows[1] = []tgbotapi.InlineKeyboardButton{}
	rows[2] = []tgbotapi.InlineKeyboardButton{}
	for i := 1; i <= 12; i++ {
		hours := 2 * (i - 1)
		if alarmExist[strconv.Itoa(i)] != "" {
			hoursStr := "✅ "
			hoursStr += strconv.Itoa(hours) + ":55"
			btn := tgbotapi.NewInlineKeyboardButtonData(hoursStr, "disable_reminder_"+strconv.Itoa(i))
			rows[(i-1)/4] = append(rows[(i-1)/4], btn)
		} else {
			hoursStr := "🚫 "
			hoursStr += strconv.Itoa(hours) + ":55"
			btn := tgbotapi.NewInlineKeyboardButtonData(hoursStr, "enable_reminder_"+strconv.Itoa(i))
			rows[(i-1)/4] = append(rows[(i-1)/4], btn)
		}
	}

	for i := 0; i <= 2; i++ {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rows[i])
	}

	return keyboard
}

func (r *Reminder) formatRemindersMessageText(playerRaw *dbmapping.Player) string {
	message := "*Ваши напоминания о битвах:*\n"
	message += "За пять минут до битвы бот может присылать вам в личные сообщения напоминание о том, "
	message += "что битва скоро состоится, и стоит встать на атаку.\n"
	message += "Кнопками ниже вы можете настроить, к каким из битв вас оповещать. Время московское.\n\n"

	currentAlarms, ok := r.getRemindersForUser(playerRaw)
	if !ok {
		message += "Не удалось получить настройки оповещений из базы. Ошибка."
	} else {
		message += "Установлено оповещений: " + strconv.Itoa(len(currentAlarms))
	}

	return message
}
