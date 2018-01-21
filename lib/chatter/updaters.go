// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package chatter

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/dbmapping"
)

// UpdateChatTitle updates chat title in database
func (ct *Chatter) UpdateChatTitle(chatRaw *dbmapping.Chat, newTitle string) (*dbmapping.Chat, bool) {
	chatRaw.Name = newTitle
	_, err := c.Db.NamedExec("UPDATE chats SET name=:name WHERE id=:id", &chatRaw)
	if err != nil {
		c.Log.Error(err)
		return chatRaw, false
	}

	return chatRaw, true
}

// UpdateChatTelegramID updates chat's TelegramID when it converts to supergroup
func (ct *Chatter) UpdateChatTelegramID(update *tgbotapi.Update) (*dbmapping.Chat, bool) {
	c.Log.Debug("Updating existing Telegram chat ID...")
	chatRaw := dbmapping.Chat{}
	err := c.Db.Get(&chatRaw, c.Db.Rebind("SELECT * FROM chats WHERE telegram_id=?"), update.Message.MigrateFromChatID)
	if err != nil {
		c.Log.Error(err.Error())
		return &chatRaw, false
	}
	if update.Message.SuperGroupChatCreated {
		chatRaw.ChatType = "supergroup"
	}
	chatRaw.TelegramID = update.Message.MigrateToChatID
	_, err = c.Db.NamedExec("UPDATE chats SET chat_type=:chat_type, telegram_id=:telegram_id WHERE id=:id", &chatRaw)
	if err != nil {
		c.Log.Error(err.Error())
		return &chatRaw, false
	}

	return &chatRaw, true
}
