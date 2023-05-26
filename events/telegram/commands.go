package telegram

import (
	"log"
	"strings"
	"telegram_bot/clients/advice"
	"telegram_bot/clients/unsplash"
)

const (
	StartCmd           = "/start"
	HelpCmd            = "/help"
	AdviceCmd          = "/advice"
	ImageCmd           = "/image"
	LocationCmd        = "/location"
	GenerationImageCmd = "/generation"
)

func (processor *EventProcessor) doCmd(text string, chatID int, userName string) error {

	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, userName)

	switch text {
	case HelpCmd:
		processor.sendHelp(chatID)
	case StartCmd:
		processor.sendHello(chatID)
	case AdviceCmd:
		processor.sendAdvice(chatID)
	case ImageCmd:
		processor.sendImage(chatID)
	case LocationCmd:
		processor.sendLocation(chatID)
	case GenerationImageCmd:
		processor.sendGenerationImage(chatID)
	default:
		processor.tgClient.SendMessage(chatID, msgUnknownCommand)
	}

	return nil
}

func (processor *EventProcessor) sendHelp(chatID int) error {
	return processor.tgClient.SendMessage(chatID, msgHelp)
}

func (processor *EventProcessor) sendHello(chatID int) error {
	return processor.tgClient.SendMessage(chatID, msgHello)
}

func (processor *EventProcessor) sendAdvice(chatID int) error {
	adviceClient := advice.New("api.adviceslip.com")

	advice, err := adviceClient.Advice()
	if err != nil {
		return err
	}
	return processor.tgClient.SendMessage(chatID, advice)
}

func (processor *EventProcessor) sendImage(chatID int) error {
	unsplashClient := unsplash.New("api.unsplash.com")

	image, err := unsplashClient.Image("")
	if err != nil {
		return err
	}

	return processor.tgClient.SendPhoto(chatID, image)
}

func (processor *EventProcessor) sendLocation(chatID int) error {
	return processor.tgClient.SendMessage(chatID, msgLocation)
}

func (processor *EventProcessor) sendGenerationImage(chatID int) error {
	return processor.tgClient.SendMessage(chatID, msgGenerationImage)
}
