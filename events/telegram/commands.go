package telegram

import (
	"log"
	"strings"
)

const (
	StartCmd = "/start"
	HelpCmd  = "/help"
)

func (processor *EventProcessor) doCmd(text string, chatID int, userName string) error {

	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, userName)

	switch text {
	case HelpCmd:
		processor.sendHelp(chatID)
	case StartCmd:
		processor.sendHello(chatID)
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
