// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package orders

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
)

// Internal functions

func (o *Orders) getOrderByID(orderID int) (dbmapping.Order, bool) {
	order := dbmapping.Order{}

	err := c.Db.Get(&order, c.Db.Rebind("SELECT * FROM orders WHERE id=?"), orderID)
	if err != nil {
		c.Log.Error(err.Error())
		return order, false
	}

	return order, true
}

func (o *Orders) sendOrder(order *dbmapping.Order) string {
	targetChats := []dbmapping.Chat{}
	ok := false

	if order.TargetSquads == "all" {
		targetChats, ok = c.Squader.GetAllSquadChats()
		if !ok {
			return "fail"
		}
	} else {
		targetChats, ok = c.Squader.GetSquadChatsBySquadsIDs(order.TargetSquads)
		if !ok {
			return "fail"
		}
	}

	for i := range targetChats {
		message := "Поступил приказ:"

		msg := tgbotapi.NewMessage(targetChats[i].TelegramID, message)
		keyboard := tgbotapi.InlineKeyboardMarkup{}
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonSwitch("В атаку!", strconv.Itoa(order.ID))
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

		msg.ReplyMarkup = keyboard
		msg.ParseMode = "Markdown"

		pinnableMessage, err := c.Bot.Send(msg)
		if err != nil {
			c.Log.Error(err.Error())
		} else {
			pinChatMessageConfig := tgbotapi.PinChatMessageConfig{
				ChatID:              pinnableMessage.Chat.ID,
				MessageID:           pinnableMessage.MessageID,
				DisableNotification: true,
			}

			_, err = c.Bot.PinChatMessage(pinChatMessageConfig)
			if err != nil {
				c.Log.Error(err.Error())
			}
		}
	}

	return "ok"
}

// External functions

// SendOrder sends order to selected or all squads
func (o *Orders) SendOrder(update *tgbotapi.Update) string {
	command := update.Message.Command()
	orderNumber := strings.TrimPrefix(command, "send_order")
	orderID, _ := strconv.Atoi(orderNumber)

	if orderID == 0 {
		return "fail"
	}

	order, ok := o.getOrderByID(orderID)
	if !ok {
		return "fail"
	}

	return o.sendOrder(&order)
}
