// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
	"strconv"
	"time"
)

// LongMessage is an easter egg
func (t *Talkers) LongMessage(update *tgbotapi.Update) string {
	message := "Я ТЕБЕ НЕ ЗЕЛЁНЫЙ АКСОЛОТЛЬ! А ТЫ - НЕ ЦИФЕРКА!"
	trigger := rand.Intn(5)
	if trigger > 3 {
		if update.Message.From.ID == 203319120 || update.Message.From.ID == 372137239 {
			message = "НЕ РАКОНЬ!"
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID

	c.Bot.Send(msg)

	return "ok"
}

// DurakMessage is an easter egg
func (t *Talkers) DurakMessage(update *tgbotapi.Update) string {
	reactions := make([]string, 0)
	reactions = append(reactions, "Сам такой!",
		"А ты типа нет?",
		"Фу, как некультурно!",
		"Попка – не дурак, Попка – самый непадающий бот!")

	// Praise the Random Gods!
	rand.Seed(time.Now().Unix())
	message := reactions[rand.Intn(len(reactions))]
	if update.Message.From.ID == 324205150 {
		message = "Молодец, Яру. Возьми с полки пирожок."
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID

	c.Bot.Send(msg)

	return "ok"
}

// MatMessage is an easter rgg
func (t *Talkers) MatMessage(update *tgbotapi.Update) string {
	reactions := make([]string, 0)
	reactions = append(reactions, "Фу, как некультурно!",
		"Иди рот с мылом помой",
		"Тшшшш!",
		"Да я твою мамку в кино водил!",
		"Приятно пообщаться с умным собеседником. К тебе это не относится.")

	// Praise the Random Gods!
	rand.Seed(time.Now().Unix())
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reactions[rand.Intn(len(reactions))])
	msg.ReplyToMessageID = update.Message.MessageID

	c.Bot.Send(msg)

	return "ok"
}

// NewYearMessage2018 pins new year 2018 message to bastion, default and academy chats.
func (t *Talkers) NewYearMessage2018() {
	academyChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.AcademyID, 10, 64)
	bastionChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.BastionID, 10, 64)
	defaultChatID, _ := strconv.ParseInt(c.Cfg.SpecialChats.DefaultID, 10, 64)

	message := "*Совет лиги Инстинкт поздравляет вас, дорогие игроки, с Новым 2018 Годом!*\n"
	message += "*Важное сообщение от* Совета лиги Инстинкт (коллективное сознательное)\n\n"
	message += "*Д*орогие бойцы Инстинкта!\n"
	message += "*А* как насчет новогоднего поздравления?\n"
	message += "*С*корее наполните бокалы!\n"
	message += "*К*онечно, хочется пожелать отличного кача...\n"
	message += "*А*, и про удачный дроп элиток не забыть!\n"
	message += "*Л*иберах побольше в охотах.\n"
	message += "*Л*овите отвагу с мистик в БО, спуску не давайте)\n"
	message += "*О*днажды мы станем топ не только по очкам, но и по атаке;\n"
	message += "*Х*уй им, в общем, всем — с бантиком!"

	msg := tgbotapi.NewMessage(defaultChatID, message)
	msg.ParseMode = "Markdown"

	pinnableMessage, _ := c.Bot.Send(msg)

	pinChatMessageConfig := tgbotapi.PinChatMessageConfig{
		ChatID:              pinnableMessage.Chat.ID,
		MessageID:           pinnableMessage.MessageID,
		DisableNotification: false,
	}

	_, err := c.Bot.PinChatMessage(pinChatMessageConfig)
	if err != nil {
		c.Log.Error(err.Error())
	}

	msg = tgbotapi.NewMessage(bastionChatID, message)
	msg.ParseMode = "Markdown"

	pinnableMessage, _ = c.Bot.Send(msg)

	pinChatMessageConfig = tgbotapi.PinChatMessageConfig{
		ChatID:              pinnableMessage.Chat.ID,
		MessageID:           pinnableMessage.MessageID,
		DisableNotification: false,
	}

	_, err = c.Bot.PinChatMessage(pinChatMessageConfig)
	if err != nil {
		c.Log.Error(err.Error())
	}

	msg = tgbotapi.NewMessage(academyChatID, message)
	msg.ParseMode = "Markdown"

	pinnableMessage, _ = c.Bot.Send(msg)

	pinChatMessageConfig = tgbotapi.PinChatMessageConfig{
		ChatID:              pinnableMessage.Chat.ID,
		MessageID:           pinnableMessage.MessageID,
		DisableNotification: false,
	}

	_, err = c.Bot.PinChatMessage(pinChatMessageConfig)
	if err != nil {
		c.Log.Error(err.Error())
	}
}
