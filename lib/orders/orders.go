// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package orders

import (
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
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

	if order.TargetSquads == "all" {
		targetChats = c.DataCache.GetAllSquadsChats()

		// Adding Academy and Bastion chat as they are both the zero chat
		academyGroupID, _ := strconv.ParseInt(c.Cfg.SpecialChats.AcademyID, 10, 64)
		bastionGroupID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
		academyChat := dbmapping.Chat{}
		bastionChat := dbmapping.Chat{}
		err := c.Db.Get(&academyChat, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=?"), academyGroupID)
		if err != nil {
			return "fail"
		}
		err = c.Db.Get(&bastionChat, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=?"), bastionGroupID)
		if err != nil {
			return "fail"
		}

		targetChats = append(targetChats, academyChat)
		targetChats = append(targetChats, bastionChat)
	} else {
		targetIDs := make([]int, 0)
		targetIDsArray := strings.Split(order.TargetSquads, ",")
		for i := range targetIDsArray {
			targetID, _ := strconv.Atoi(targetIDsArray[i])
			targetIDs = append(targetIDs, targetID)
		}
		targetChats = c.DataCache.GetSquadsChatsBySquadsIDs(targetIDs)

		targetChatsIDs := strings.Split(order.TargetSquads, ",")
		for i := range targetChatsIDs {
			if targetChatsIDs[i] == "0" {
				// Adding Academy and Bastion chat as they are both the zero chat
				academyGroupID, _ := strconv.ParseInt(c.Cfg.SpecialChats.AcademyID, 10, 64)
				bastionGroupID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
				academyChat := dbmapping.Chat{}
				bastionChat := dbmapping.Chat{}
				err := c.Db.Get(&academyChat, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=?"), academyGroupID)
				if err != nil {
					return "fail"
				}
				err = c.Db.Get(&bastionChat, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=?"), bastionGroupID)
				if err != nil {
					return "fail"
				}

				targetChats = append(targetChats, academyChat)
				targetChats = append(targetChats, bastionChat)
			}
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
