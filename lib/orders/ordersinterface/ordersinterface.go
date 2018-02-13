// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package ordersinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

// OrdersInterface implements Orders for importing via appcontext.
type OrdersInterface interface {
	Init()

	GetAllOrders() ([]dbmapping.Order, bool)
	GetOrderByID(orderID int) (dbmapping.Order, bool)

	ListAllOrders(update *tgbotapi.Update) string

	SendOrder(update *tgbotapi.Update) string
}
