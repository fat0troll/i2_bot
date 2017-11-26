// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

// RouteInline routes inline requests to bot
func (r *Router) RouteInline(update *tgbotapi.Update) string {
	availableCommands := make(map[string]string)
	availableCommands["0"] = "🌲Лес"
	availableCommands["1"] = "⛰Горы"
	availableCommands["2"] = "🚣Озеро"
	availableCommands["3"] = "🏙Город"
	availableCommands["4"] = "🏛Катакомбы"
	availableCommands["5"] = "⛪️Кладбище"
	outputCommands := make(map[string]string)
	for i, value := range availableCommands {
		if strings.Contains(value, update.InlineQuery.Query) {
			outputCommands[i] = value
		}
	}

	results := make([]interface{}, 0)
	for i, value := range outputCommands {
		article := tgbotapi.NewInlineQueryResultArticle(i, "Команда боту @PokememBroBot:", value)
		article.Description = value

		results = append(results, article)
	}

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       results,
	}

	_, err := c.Bot.AnswerInlineQuery(inlineConf)
	if err != nil {
		c.Log.Error(err.Error())
	}

	return "fail"
}
