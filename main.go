package main

import (
	"flag"
	"log"
	client "telegram_bot/clients/telegram"
	consumer "telegram_bot/consumer/event-consumer"
	telegram "telegram_bot/events/telegram"
)

const (
	host = "api.telegram.org"
)

func main() {

	tgClient := client.New(host, mustToken())

	eventProcessor := telegram.New(tgClient)
	log.Printf("service start")

	consumer := consumer.New(eventProcessor, eventProcessor, 100)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		" ",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
