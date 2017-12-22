// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package pinner

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"lab.pztrn.name/fat0troll/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
)

func (p *Pinner) execMassMessagePin(update *tgbotapi.Update, groupChats []dbmapping.Chat) string {
	messageToPin := update.Message.CommandArguments()
	for i := range groupChats {
		if groupChats[i].ChatType == "supergroup" {
			message := messageToPin + "\n\n"
			message += "© " + c.Users.GetPrettyName(update.Message.From)

			msg := tgbotapi.NewMessage(groupChats[i].TelegramID, message)
			msg.ParseMode = "Markdown"

			pinnableMessage, err := c.Bot.Send(msg)
			if err != nil {
				c.Log.Error(err.Error())

				message := "*Ваше сообщение не отправлено.*\n\n"
				message += "Обычно это связано с тем, что нарушена разметка Markdown. "
				message += "К примеру, если вы хотели использовать нижнее\\_подчёркивание, то печатать его надо так — \\\\_. То же самое касается всех управляющих разметкой символов в Markdown в случае, если вы их хотите использовать как текст, а не как управляющий символ Markdown."

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				msg.ParseMode = "Markdown"

				c.Bot.Send(msg)

				return "fail"
			}

			pinChatMessageConfig := tgbotapi.PinChatMessageConfig{
				ChatID:              pinnableMessage.Chat.ID,
				MessageID:           pinnableMessage.MessageID,
				DisableNotification: true,
			}

			_, err = c.Bot.PinChatMessage(pinChatMessageConfig)
			if err != nil {
				c.Log.Error(err.Error())
			}
		}
	}

	message := "*Ваше сообщение отправлено и запинено в чаты, где сидит бот.*\n\n"
	message += "Текст отправленного сообщения:\n\n"
	message += messageToPin

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"

	c.Bot.Send(msg)

	return "ok"
}

// PinMessageToAllChats pins message to all groups where bot exist
func (p *Pinner) PinMessageToAllChats(update *tgbotapi.Update) string {
	messageToPin := update.Message.CommandArguments()

	if messageToPin == "" {
		return "fail"
	}

	groupChats, ok := c.Chatter.GetAllGroupChats()
	if !ok {
		return "fail"
	}

	return p.execMassMessagePin(update, groupChats)
}

// PinMessageToSomeChats pins message to selected groups where bot exist
func (p *Pinner) PinMessageToSomeChats(update *tgbotapi.Update) string {
	commandArgs := update.Message.CommandArguments()
	commandArgsList := strings.Split(commandArgs, " ")
	if len(commandArgsList) < 2 {
		return "fail"
	}

	chatsToPin := ""
	messageToPin := ""

	for i := range commandArgsList {
		if i == 0 {
			chatsToPin = commandArgsList[i]
		} else {
			messageToPin += commandArgsList[i] + " "
		}
	}

	if messageToPin == "" {
		return "fail"
	}

	groupChats, ok := c.Chatter.GetGroupChatsByIDs(chatsToPin)
	if !ok {
		return "fail"
	}
	c.Log.Debug("Got " + strconv.Itoa(len(groupChats)) + " group chats...")

	return p.execMassMessagePin(update, groupChats)
}
