// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
    // stdlib
	"log"
    "math/rand"
    "time"
    // 3rd party
	"gopkg.in/telegram-bot-api.v4"
)

func DurakMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    reactions := make([]string, 0)
    reactions = append(reactions, "Сам такой!",
        "А ты типа нет?",
        "Фу, как некультурно!",
        "Профессор, если вы такой умный, то почему вы такой бедный? /donate",
        "Попка – не дурак, Попка – самый непадающий бот!")

    // Praise the Random Gods!
    rand.Seed(time.Now().Unix())
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, reactions[rand.Intn(len(reactions))])
    msg.ReplyToMessageID = update.Message.MessageID

    bot.Send(msg)
}

func MatMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    reactions := make([]string, 0)
    reactions = append(reactions, "Фу, как некультурно!",
        "Иди рот с мылом помой",
        "Тшшшш!",
        "Да я твою мамку в кино водил!")

    // Praise the Random Gods!
    rand.Seed(time.Now().Unix())
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, reactions[rand.Intn(len(reactions))])
    msg.ReplyToMessageID = update.Message.MessageID

    bot.Send(msg)
}
