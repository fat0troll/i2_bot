// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package chatter

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"time"
)

func (ct *Chatter) getAllGroupChatsWithSquads() ([]dbmapping.ChatSquad, bool) {
	chatsSquads := []dbmapping.ChatSquad{}
	groupChats := []dbmapping.Chat{}

	err := c.Db.Select(&groupChats, "SELECT * FROM chats WHERE chat_type IN ('group', 'supergroup')")
	if err != nil {
		c.Log.Error(err)
		return chatsSquads, false
	}

	for i := range groupChats {
		chatSquad := dbmapping.ChatSquad{}
		squad := dbmapping.Squad{}
		err = c.Db.Get(&squad, c.Db.Rebind("SELECT * FROM squads WHERE chat_id=?"), groupChats[i].ID)
		if err != nil {
			c.Log.Debug(err)
		} else {
			chatSquad.ChatRole = "squad"
		}
		err = c.Db.Get(&squad, c.Db.Rebind("SELECT * FROM squads WHERE flood_chat_id=?"), groupChats[i].ID)
		if err != nil {
			c.Log.Debug(err)
		} else {
			chatSquad.ChatRole = "flood"
		}

		chatSquad.Squad = squad
		chatSquad.Chat = groupChats[i]

		chatsSquads = append(chatsSquads, chatSquad)
	}

	return chatsSquads, true
}

// GetChatByID returns dbmapping.Chat instance with given ID.
func (ct *Chatter) GetChatByID(chatID int64) (dbmapping.Chat, bool) {
	chatRaw := dbmapping.Chat{}
	err := c.Db.Get(&chatRaw, c.Db.Rebind("SELECT * FROM chats WHERE id=?"), chatID)
	if err != nil {
		c.Log.Error(err)
		return chatRaw, false
	}

	return chatRaw, true
}

// GetOrCreateChat seeks for chat in database via Telegram update.
// In case, when there is no chat with such ID, new chat will be created.
func (ct *Chatter) GetOrCreateChat(telegramUpdate *tgbotapi.Update) (dbmapping.Chat, bool) {
	chatRaw := dbmapping.Chat{}
	c.Log.Debug("TGID: ", telegramUpdate.Message.Chat.ID)
	err := c.Db.Get(&chatRaw, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=?"), telegramUpdate.Message.Chat.ID)
	if err != nil {
		c.Log.Error("Chat stream not found in database.")
		c.Log.Error(err.Error())

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
			c.Log.Error(err.Error())
			return chatRaw, false
		}
		err2 := c.Db.Get(&chatRaw, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=? AND chat_type=?"), chatRaw.TelegramID, chatRaw.ChatType)
		if err2 != nil {
			c.Log.Error(err2)
			return chatRaw, false
		}
	} else {
		c.Log.Info("Chat stream found in database.")
	}

	return chatRaw, true
}

// GetAllPrivateChats returns all private chats
func (ct *Chatter) GetAllPrivateChats() ([]dbmapping.Chat, bool) {
	privateChats := []dbmapping.Chat{}

	err := c.Db.Select(&privateChats, "SELECT * FROM chats WHERE chat_type='private'")
	if err != nil {
		c.Log.Error(err)
		return privateChats, false
	}

	return privateChats, true
}

// GetLeaguePrivateChats returns all private chats which profiles are in our league
func (ct *Chatter) GetLeaguePrivateChats() ([]dbmapping.Chat, bool) {
	privateChats := []dbmapping.Chat{}

	err := c.Db.Select(&privateChats, "SELECT c.* FROM chats c, players p WHERE c.chat_type='private' AND p.telegram_id = c.telegram_id AND p.league_id = 1")
	if err != nil {
		c.Log.Error(err)
		return privateChats, false
	}

	return privateChats, true
}

// GetAllGroupChats returns all group chats
func (ct *Chatter) GetAllGroupChats() ([]dbmapping.Chat, bool) {
	groupChats := []dbmapping.Chat{}

	err := c.Db.Select(&groupChats, "SELECT * FROM chats WHERE chat_type IN ('group', 'supergroup')")
	if err != nil {
		c.Log.Error(err)
		return groupChats, false
	}

	return groupChats, true
}
