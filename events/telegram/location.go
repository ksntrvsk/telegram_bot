package telegram

import (
	"strconv"
	"strings"
)

func (processor *EventProcessor) sendTestLocation(text string, chatID int, userName string) error {

	location := strings.Split(text, " ")

	lat, _ := strconv.ParseFloat(location[0], 64)
	lon, _ := strconv.ParseFloat(location[1], 64)

	processor.tgClient.SendLocation(chatID, lat, lon)
	return nil
}
