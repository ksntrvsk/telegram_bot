package eventconsumer

import (
	"log"
	"telegram_bot/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (consumer Consumer) Start() error {

	for {
		gotEvents, err := consumer.fetcher.Fetch(consumer.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := consumer.HandleEvents(gotEvents); err != nil {
			log.Printf(err.Error())

			continue
		}
	}
}

func (consumer *Consumer) HandleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := consumer.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
