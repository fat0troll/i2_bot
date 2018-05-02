// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package router

import (
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// RouteInline routes inline requests to bot
func (r *Router) RouteInline(update tgbotapi.Update) string {
	playerRaw, err := c.DataCache.GetOrCreatePlayerByTelegramID(update.InlineQuery.From.ID)
	if err != nil {
		c.Log.Error(err.Error())
		return "fail"
	}

	results := make([]interface{}, 0)

	if (playerRaw.LeagueID != 1) || (playerRaw.Status == "banned") {
		article := tgbotapi.NewInlineQueryResultArticle("0", "Команда боту @PokememBroBot:", "/me")
		article.Description = "Получить статистику"

		results = append(results, article)
	} else {
		orderNumber, _ := strconv.Atoi(update.InlineQuery.Query)
		if orderNumber != 0 {
			order, ok := c.Orders.GetOrderByID(orderNumber)
			if !ok {
				return "fail"
			}

			attackTarget := ""
			if order.Target == "M" {
				attackTarget = "⚔ 🈳 МИСТИКА"
			} else {
				attackTarget = "⚔ 🈵 ОТВАГА"
			}

			article := tgbotapi.NewInlineQueryResultArticle(strconv.Itoa(orderNumber), "Выполнить приказ отряда:", attackTarget)
			article.Description = attackTarget

			results = append(results, article)
		} else {
			if update.InlineQuery.Query != "Статы" {
				availableCommands := make(map[string]string)
				availableCommands["10"] = "🌲Лес"
				availableCommands["11"] = "⛰Горы"
				availableCommands["12"] = "🚣Озеро"
				availableCommands["13"] = "🏙Город"
				availableCommands["14"] = "🏛Катакомбы"
				availableCommands["15"] = "⛪️Кладбище"
				outputCommands := make(map[string]string)
				for i, value := range availableCommands {
					if strings.Contains(value, update.InlineQuery.Query) {
						outputCommands[i] = value
					}
				}

				for i, value := range outputCommands {
					article := tgbotapi.NewInlineQueryResultArticle(i, "Команда боту @PokememBroBot:", value)
					article.Description = value

					results = append(results, article)
				}
			} else {
				article := tgbotapi.NewInlineQueryResultArticle("0", "Команда боту @PokememBroBot:", "/me")
				article.Description = "Получить статистику"

				results = append(results, article)
			}
		}
	}

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       results,
	}

	_, err = c.Bot.AnswerInlineQuery(inlineConf)
	if err != nil {
		c.Log.Error(err.Error())
	}

	return "ok"
}
