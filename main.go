package main

import (
	"flag"
	"log"
	"telegram-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
	//processor.New()
	// fetcher.New()
	//Consumer(fetcher, processor)
}

func mustToken() string {
	token := flag.String("t", "", "token for using telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("telegram token isnt specified")
	}

	return *token
}
