// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package getters

import (
	// stdlib
	"log"
	"time"
	// local
	"../dbmapping"
)

// CreateBroadcastMessage creates broadcast message item in database
func (g *Getters) CreateBroadcastMessage(playerRaw *dbmapping.Player, messageBody string, broadcastType string) (dbmapping.Broadcast, bool) {
	messageRaw := dbmapping.Broadcast{}
	messageRaw.Text = messageBody
	messageRaw.Status = "new"
	messageRaw.BroadcastType = broadcastType
	messageRaw.AuthorID = playerRaw.ID
	messageRaw.CreatedAt = time.Now().UTC()
	_, err := c.Db.NamedExec("INSERT INTO broadcasts VALUES(NULL, :text, :broadcast_type, :status, :author_id, :created_at)", &messageRaw)
	if err != nil {
		log.Printf(err.Error())
		return messageRaw, false
	}
	err2 := c.Db.Get(&messageRaw, c.Db.Rebind("SELECT * FROM broadcasts WHERE author_id=? AND text=?"), messageRaw.AuthorID, messageRaw.Text)
	if err2 != nil {
		log.Println(err2)
		return messageRaw, false
	}

	return messageRaw, true
}

// GetBroadcastMessageByID returns dbmapping.Broadcast instance with given ID.
func (g *Getters) GetBroadcastMessageByID(messageID int) (dbmapping.Broadcast, bool) {
	messageRaw := dbmapping.Broadcast{}
	err := c.Db.Get(&messageRaw, c.Db.Rebind("SELECT * FROM broadcasts WHERE id=?"), messageID)
	if err != nil {
		log.Println(err)
		return messageRaw, false
	}

	return messageRaw, true
}

// UpdateBroadcastMessageStatus updates broadcast message status
func (g *Getters) UpdateBroadcastMessageStatus(messageID int, messageStatus string) (dbmapping.Broadcast, bool) {
	messageRaw := dbmapping.Broadcast{}
	err := c.Db.Get(&messageRaw, c.Db.Rebind("SELECT * FROM broadcasts WHERE id=?"), messageID)
	if err != nil {
		log.Println(err)
		return messageRaw, false
	}
	messageRaw.Status = messageStatus
	_, err = c.Db.NamedExec("UPDATE broadcasts SET status=:status WHERE id=:id", &messageRaw)
	if err != nil {
		log.Printf(err.Error())
		return messageRaw, false
	}
	err = c.Db.Get(&messageRaw, c.Db.Rebind("SELECT * FROM broadcasts WHERE author_id=? AND text=?"), messageRaw.AuthorID, messageRaw.Text)
	if err != nil {
		log.Println(err)
		return messageRaw, false
	}

	return messageRaw, true
}
