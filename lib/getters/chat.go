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
func (g *Getters) GetChatByID(chatID int64) (dbmapping.Chat, bool) {
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
	log.Println("TGID: ", telegramUpdate.Message.Chat.ID)
	err := c.Db.Get(&chatRaw, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=?"), telegramUpdate.Message.Chat.ID)
	if err != nil {
		log.Printf("Chat stream not found in database.")
		log.Printf(err.Error())

		nameOfChat := ""
		if telegramUpdate.Message.Chat.FirstName != "" {
			nameOfChat += telegramUpdate.Message.Chat.FirstName
		}
		if telegramUpdate.Message.Chat.LastName != "" {
			nameOfChat += " " + telegramUpdate.Message.Chat.LastName
		}
		if telegramUpdate.Message.Chat.Title != "" {
			if nameOfChat != "" {
				nameOfChat += " [" + telegramUpdate.Message.Chat.Title + "]"
			} else {
				nameOfChat = telegramUpdate.Message.Chat.Title
			}
		}

		chatRaw.Name = nameOfChat
		chatRaw.ChatType = telegramUpdate.Message.Chat.Type
		chatRaw.TelegramID = telegramUpdate.Message.Chat.ID
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

// UpdateChatTitle updates chat title in database
func (g *Getters) UpdateChatTitle(chatRaw dbmapping.Chat, newTitle string) (dbmapping.Chat, bool) {
	chatRaw.Name = newTitle
	_, err := c.Db.NamedExec("UPDATE chats SET name=:name WHERE id=:id", &chatRaw)
	if err != nil {
		log.Println(err)
		return chatRaw, false
	}

	return chatRaw, true
}
