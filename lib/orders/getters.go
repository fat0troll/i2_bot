// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package orders

import (
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

// GetAllOrders returns all orders in database
func (o *Orders) GetAllOrders() ([]dbmapping.Order, bool) {
	orders := []dbmapping.Order{}

	err := c.Db.Select(&orders, "SELECT * FROM orders ORDER BY created_at asc")
	if err != nil {
		c.Log.Error(err)
		return orders, false
	}

	return orders, true
}

// GetOrderByID returns single order by ID
func (o *Orders) GetOrderByID(orderID int) (dbmapping.Order, bool) {
	order := dbmapping.Order{}

	err := c.Db.Get(&order, c.Db.Rebind("SELECT * FROM orders WHERE id=?"), orderID)
	if err != nil {
		c.Log.Error(err.Error())
		return order, false
	}

	return order, true
}
