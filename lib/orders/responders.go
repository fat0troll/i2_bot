// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package orders

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

// ListAllOrders returns to user all orders in database
func (o *Orders) ListAllOrders(update *tgbotapi.Update) string {
	orders, ok := o.GetAllOrders()
	if !ok {
		return "fail"
	}

	message := "*Приказы на атаку*\n"
	for i := range orders {
		message += "\\[" + strconv.Itoa(orders[i].ID) + "] " + orders[i].TargetSquads + " → "
		if orders[i].Target == "M" {
			message += "🈳 МИСТИКА "
		} else {
			message += "🈵 ОТВАГА "
		}
		if orders[i].Scheduled {
			message += "запланировано на "
			message += orders[i].ScheduledAt.Time.Format("02.01.2006 15:04:05")
		}
		if orders[i].Status == "sent" {
			message += "\nПросмотреть выполнение приказа: /show\\_order" + strconv.Itoa(orders[i].ID)
		} else {
			message += "\nОтправить приказ прямо сейчас: /send\\_order" + strconv.Itoa(orders[i].ID)
		}
		message += "\n"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
