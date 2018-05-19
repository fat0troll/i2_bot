// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package broadcaster

import (
	"time"

	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

func (b *Broadcaster) createBroadcastMessage(playerRaw *dbmapping.Player, messageBody string, broadcastType string) (dbmapping.Broadcast, bool) {
	messageRaw := dbmapping.Broadcast{}
	messageRaw.Text = messageBody
	messageRaw.Status = "new"
	messageRaw.BroadcastType = broadcastType
	messageRaw.AuthorID = playerRaw.ID
	messageRaw.CreatedAt = time.Now().UTC()
	_, err := c.Db.NamedExec("INSERT INTO broadcasts VALUES(NULL, :text, :broadcast_type, :status, :author_id, :created_at)", &messageRaw)
	if err != nil {
		c.Log.Error(err.Error())
		return messageRaw, false
	}
	err2 := c.Db.Get(&messageRaw, c.Db.Rebind("SELECT * FROM broadcasts WHERE author_id=? AND text=?"), messageRaw.AuthorID, messageRaw.Text)
	if err2 != nil {
		c.Log.Error(err2)
		return messageRaw, false
	}

	return messageRaw, true
}

func (b *Broadcaster) getBroadcastMessageByID(messageID int) (dbmapping.Broadcast, bool) {
	messageRaw := dbmapping.Broadcast{}
	err := c.Db.Get(&messageRaw, c.Db.Rebind("SELECT * FROM broadcasts WHERE id=?"), messageID)
	if err != nil {
		c.Log.Error(err)
		return messageRaw, false
	}

	return messageRaw, true
}

func (b *Broadcaster) updateBroadcastMessageStatus(messageID int, messageStatus string) (dbmapping.Broadcast, bool) {
	messageRaw := dbmapping.Broadcast{}
	err := c.Db.Get(&messageRaw, c.Db.Rebind("SELECT * FROM broadcasts WHERE id=?"), messageID)
	if err != nil {
		c.Log.Error(err.Error())
		return messageRaw, false
	}
	messageRaw.Status = messageStatus
	_, err = c.Db.NamedExec("UPDATE broadcasts SET status=:status WHERE id=:id", &messageRaw)
	if err != nil {
		c.Log.Error(err.Error())
		return messageRaw, false
	}
	err = c.Db.Get(&messageRaw, c.Db.Rebind("SELECT * FROM broadcasts WHERE author_id=? AND text=?"), messageRaw.AuthorID, messageRaw.Text)
	if err != nil {
		c.Log.Error(err.Error())
		return messageRaw, false
	}

	return messageRaw, true
}
