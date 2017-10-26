// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package getters

import (
	// stdlib
	"log"
	"time"
	// 3rd-party
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// local
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
)

// GetChatByID returns dbmapping.Chat instance with given ID.
func (g *Getters) GetChatByID(chatID int) (dbmapping.Chat, bool) {
	chatRaw := dbmapping.Chat{}
	err := c.Db.Get(&chatRaw, c.Db.Rebind("SELECT * FROM chats WHERE id=?"), chatID)
	if err != nil {
		log.Println(err)
		return chatRaw, false
	}

	return chatRaw, true
}

// GetOrCreateChat seeks for chat in database via Telegram update.
// In case, when there is no chat with such ID, new chat will be created.
func (g *Getters) GetOrCreateChat(telegramUpdate *tgbotapi.Update) (dbmapping.Chat, bool) {
	chatRaw := dbmapping.Chat{}
	err := c.Db.Get(&chatRaw, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=?"), telegramUpdate.Message.Chat.ID)
	if err != nil {
		log.Printf("Chat stream not found in database.")
		log.Printf(err.Error())

		chatRaw.Name = telegramUpdate.Message.Chat.FirstName + " " + telegramUpdate.Message.Chat.LastName
		chatRaw.ChatType = telegramUpdate.Message.Chat.Type
		chatRaw.TelegramID = int(telegramUpdate.Message.Chat.ID)
		chatRaw.CreatedAt = time.Now().UTC()
		_, err = c.Db.NamedExec("INSERT INTO chats VALUES(NULL, :name, :chat_type, :telegram_id, :created_at)", &chatRaw)
		if err != nil {
			log.Printf(err.Error())
			return chatRaw, false
		}
		err2 := c.Db.Get(&chatRaw, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=? AND chat_type=?"), chatRaw.TelegramID, chatRaw.ChatType)
		if err2 != nil {
			log.Println(err2)
			return chatRaw, false
		}
	} else {
		log.Printf("Chat stream found in database.")
	}

	return chatRaw, true
}

// GetAllPrivateChats returns all private chats
func (g *Getters) GetAllPrivateChats() ([]dbmapping.Chat, bool) {
	privateChats := []dbmapping.Chat{}

	err := c.Db.Select(&privateChats, "SELECT * FROM chats WHERE chat_type='private'")
	if err != nil {
		log.Println(err)
		return privateChats, false
	}

	return privateChats, true
}
