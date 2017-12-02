// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
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
