package telegram

import (
	"errors"
	"fmt"
	"strings"
	"telegram_bot/clients/telegram"
	"telegram_bot/events"
)

type EventProcessor struct {
	tgClient *telegram.Client
	offset   int
}

type Meta struct {
	ChatID   int
	Username string
}

func New(client *telegram.Client) *EventProcessor {
	return &EventProcessor{
		tgClient: client,
	}
}

func (eventProcessor *EventProcessor) Fetch(limit int) ([]events.Event, error) {

	updates, err := eventProcessor.tgClient.Updates(eventProcessor.offset, limit)
	if err != nil {
		return nil, myErr("can't get updates", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	result := make([]events.Event, 0, len(updates))
	for _, update := range updates {
		result = append(result, event(update))
	}

	eventProcessor.offset = updates[len(updates)-1].ID + 1

	return result, nil
}

func (eventProcessor *EventProcessor) Process(event events.Event) error {

	switch event.Type {
	case events.Message:
		return eventProcessor.processMessage(event)
	default:
		return fmt.Errorf("unknown event type")
	}
}

func (eventProcessor *EventProcessor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return myErr("can't process message", err)
	}

	if strings.Contains(event.Text, "/") {
		if err := eventProcessor.doCmd(event.Text, meta.ChatID, meta.Username); err != nil {
			return myErr("can't process message: %w", err)
		}
	} else {
		if err := eventProcessor.sendTestLocation(event.Text, meta.ChatID, meta.Username); err != nil {
			return myErr("can't process message: %w", err)
		}
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	result, ok := event.Meta.(Meta)
	if !ok {
		err := errors.New("unknown meta type")
		return Meta{}, myErr("can't get meta", err)
	}

	return result, nil
}

func event(update telegram.Update) events.Event {

	updateType := fetchType(update)
	updateText := fetchText(update)

	result := events.Event{
		Type: updateType,
		Text: updateText,
	}

	if updateType == events.Message {
		result.Meta = Meta{
			ChatID:   update.Message.Chat.ID,
			Username: update.Message.From.UserName,
		}
	}
	return result
}

func fetchType(update telegram.Update) events.TypeEvent {
	if update.Message == nil {
		return events.Unknown
	}
	return events.Message
}

func fetchText(update telegram.Update) string {
	if update.Message == nil {
		return ""
	}
	return update.Message.Text
}

func myErr(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}
