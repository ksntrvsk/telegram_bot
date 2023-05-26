package telegram

import (
	"strconv"
	"strings"
	events "telegram_bot/events/openAI"
)

func (processor *EventProcessor) processText(text string, chatID int) error {

	lat, lon := location(text)
	processor.tgClient.SendLocation(chatID, lat, lon)

	generationImage, err := events.GenerationImage(text)
	if err != nil {
		return err
	}

	return processor.tgClient.SendPhoto(chatID, generationImage)
}

func location(text string) (float64, float64) {
	location := strings.Split(text, " ")
	lat, _ := strconv.ParseFloat(location[0], 64)
	lon, _ := strconv.ParseFloat(location[1], 64)
	return lat, lon
}
