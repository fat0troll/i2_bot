// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"strconv"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
)

func (dc *DataCache) initChats() {
	c.Log.Info("Initializing Chats storage...")
	dc.chats = make(map[int]*dbmapping.Chat)
}

func (dc *DataCache) loadChats() {
	c.Log.Info("Load current Chats data from database to DataCache...")
	chats := []dbmapping.Chat{}
	err := c.Db.Select(&chats, "SELECT * FROM chats")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}

	dc.chatsMutex.Lock()
	for i := range chats {
		dc.chats[chats[i].ID] = &chats[i]
	}
	c.Log.Info("Loaded chats in DataCache: " + strconv.Itoa(len(dc.chats)))
	dc.chatsMutex.Unlock()
}

// External function

// GetAllGroupChats returns all bot's group chats
func (dc *DataCache) GetAllGroupChats() []dbmapping.Chat {
	chats := []dbmapping.Chat{}

	dc.chatsMutex.Lock()
	for i := range dc.chats {
		if dc.chats[i].ChatType == "group" || dc.chats[i].ChatType == "supergroup" {
			chats = append(chats, *dc.chats[i])
		}
	}
	dc.chatsMutex.Unlock()

	return chats
}

// GetAllPrivateChats returns all bot's private chats
func (dc *DataCache) GetAllPrivateChats() []dbmapping.Chat {
	chats := []dbmapping.Chat{}

	dc.chatsMutex.Lock()
	for i := range dc.chats {
		if dc.chats[i].ChatType == "private" {
			chats = append(chats, *dc.chats[i])
		}
	}
	dc.chatsMutex.Unlock()

	return chats
}

// GetChatByID returns Chat by it's ID
func (dc *DataCache) GetChatByID(chatID int) (*dbmapping.Chat, error) {
	if dc.chats[chatID] != nil {
		return dc.chats[chatID], nil
	}

	return nil, errors.New("There is no chat with ID=" + strconv.Itoa(chatID))
}

// GetGroupChatsByIDs returns bot's group chats with given IDs
func (dc *DataCache) GetGroupChatsByIDs(chatIDs []int) []dbmapping.Chat {
	chats := []dbmapping.Chat{}

	dc.chatsMutex.Lock()
	for i := range dc.chats {
		if dc.chats[i].ChatType == "group" || dc.chats[i].ChatType == "supergroup" {
			for j := range chatIDs {
				if dc.chats[i].ID == j {
					chats = append(chats, *dc.chats[i])
				}
			}
		}
	}
	dc.chatsMutex.Unlock()

	return chats
}

// GetLeaguePrivateChats returns private chats for all league members
func (dc *DataCache) GetLeaguePrivateChats() []dbmapping.Chat {
	dc.playersMutex.Lock()
	dc.chatsMutex.Lock()

	chats := []dbmapping.Chat{}

	for i := range dc.players {
		if dc.players[i].Status != "banned" && dc.players[i].Status != "spy" && dc.players[i].Status != "league_changed" && dc.players[i].LeagueID == 1 {
			for ii := range dc.chats {
				if int(dc.chats[ii].TelegramID) == int(dc.players[i].TelegramID) {
					chats = append(chats, *dc.chats[ii])
				}
			}
		}
	}

	dc.playersMutex.Unlock()
	dc.chatsMutex.Unlock()

	return chats
}

// GetOrCreateChat returns current or new Chat object by Telegram update
func (dc *DataCache) GetOrCreateChat(update *tgbotapi.Update) (*dbmapping.Chat, error) {
	telegramID := update.Message.Chat.ID
	chatRaw := dbmapping.Chat{}
	c.Log.Info("DataCache: Getting chat with Telegram ID=", telegramID)

	dc.chatsMutex.Lock()
	for i := range dc.chats {
		if dc.chats[i].TelegramID == int64(telegramID) {
			c.Log.Debug("Chat stream found in DataCache")
			dc.chatsMutex.Unlock()
			return dc.chats[i], nil
		}
	}
	dc.chatsMutex.Unlock()

	// If we're here: there is no chat with given Telegram ID
	c.Log.Error("Chat stream not found in DataCache. Adding chat...")

	nameOfChat := ""
	if update.Message.Chat.FirstName != "" {
		nameOfChat += update.Message.Chat.FirstName
	}
	if update.Message.Chat.LastName != "" {
		nameOfChat += " " + update.Message.Chat.LastName
	}
	if update.Message.Chat.Title != "" {
		if nameOfChat != "" {
			nameOfChat += " [" + update.Message.Chat.Title + "]"
		} else {
			nameOfChat = update.Message.Chat.Title
		}
	}

	chatRaw.Name = nameOfChat
	chatRaw.ChatType = update.Message.Chat.Type
	chatRaw.TelegramID = update.Message.Chat.ID
	chatRaw.CreatedAt = time.Now().UTC()
	_, err := c.Db.NamedExec("INSERT INTO chats VALUES(NULL, :name, :chat_type, :telegram_id, :created_at)", &chatRaw)
	if err != nil {
		c.Log.Error(err.Error())
		return nil, err
	}
	err = c.Db.Get(&chatRaw, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=? AND chat_type=?"), chatRaw.TelegramID, chatRaw.ChatType)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	dc.chatsMutex.Lock()
	dc.chats[chatRaw.ID] = &chatRaw
	dc.chatsMutex.Unlock()

	return &chatRaw, nil
}

// UpdateChatTitle updates chat title with new one
func (dc *DataCache) UpdateChatTitle(chatID int, newTitle string) (*dbmapping.Chat, error) {
	chatRaw, err := c.DataCache.GetChatByID(chatID)
	if err != nil {
		return nil, err
	}
	chatRaw.Name = newTitle
	_, err = c.Db.NamedExec("UPDATE chats SET name=:name WHERE id=:id", &chatRaw)
	if err != nil {
		return nil, err
	}

	dc.chatsMutex.Lock()
	dc.chats[chatRaw.ID] = chatRaw
	dc.chatsMutex.Unlock()

	return dc.chats[chatRaw.ID], nil
}
