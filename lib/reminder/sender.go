// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package reminder

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"time"
)

// SendReminders sends reminders for users about coming league tournament
func (r *Reminder) SendReminders() {
	currentHour := time.Now().Hour()
	nextTournamentID := (currentHour / 2) + 1

	playersRaw := []dbmapping.Player{}
	err := c.Db.Select(&playersRaw, "SELECT p.* FROM players p, alarms a WHERE a.turnir_number=? AND a.player_id = p.id GROUP BY p.id", nextTournamentID)
	if err != nil {
		c.Log.Error(err.Error())
	}

	for i := range playersRaw {
		message := "*Турнир Лиг покемемов состоится через пять минут!*\n"
		message += "Вперёд, за опытом и деньгами! Выстави атаку и жди результата!"

		msg := tgbotapi.NewMessage(int64(playersRaw[i].TelegramID), message)
		msg.ParseMode = "Markdown"

		c.Bot.Send(msg)
	}
}
