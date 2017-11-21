// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package statistics

import (
	"strconv"
	"strings"
)

// GetPoints returns points to use in database
func (s *Statistics) GetPoints(pointsStr string) int {
	value := 0
	if strings.HasSuffix(pointsStr, "K") {
		valueNumber := strings.Replace(pointsStr, "K", "", 1)
		valueFloat, _ := strconv.ParseFloat(valueNumber, 64)
		value = int(valueFloat * 1000)
	} else if strings.HasSuffix(pointsStr, "M") {
		valueNumber := strings.Replace(pointsStr, "M", "", 1)
		valueFloat, _ := strconv.ParseFloat(valueNumber, 64)
		value = int(valueFloat * 1000000)
	} else {
		value, _ = strconv.Atoi(pointsStr)
	}
	return value
}

// GetPrintablePoints returns to output points (ht, attack, mp...) formatted
// like in PokememBroBot itself.
func (s *Statistics) GetPrintablePoints(points int) string {
	if points < 1000 {
		return strconv.Itoa(points)
	} else if points < 1000000 {
		floatNum := float64(points) / 1000.0
		return strconv.FormatFloat(floatNum, 'f', -1, 64) + "K"
	} else {
		floatNum := float64(points) / 1000000.0
		return strconv.FormatFloat(floatNum, 'f', -1, 64) + "M"
	}
}
