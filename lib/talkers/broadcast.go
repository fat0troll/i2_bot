// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import ( // stdlib
	// 3rd party
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// AdminBroadcastMessage sends message to all private chats with bot
func (t *Talkers) AdminBroadcastMessage(update tgbotapi.Update) string {
	broadcastingMessageBody := strings.Replace(update.Message.Text, "/send_all", "", 1)

	privateChats, ok := c.Getters.GetAllPrivateChats()
	if !ok {
		return "fail"
	}

	for i := range privateChats {
		chat := privateChats[i]
		broadcastingMessage := "*Привет, " + chat.Name + "!*\n\n"
		broadcastingMessage += "*Важное сообщение от администратора " + update.Message.From.FirstName + " " + update.Message.From.LastName + "* (@" + update.Message.From.UserName + ")\n\n"
		broadcastingMessage += broadcastingMessageBody

		msg := tgbotapi.NewMessage(int64(chat.TelegramID), broadcastingMessage)
		msg.ParseMode = "Markdown"
		c.Bot.Send(msg)
	}

	message := "Сообщение всем отправлено. Надеюсь, пользователи бота за него тебя не убьют.\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}
