package main

import (
	"flag"
	"log"
)

func main() {

	t := mustToken()

	//tgClient = telegram.New(token) - через клиента происходит общение с телеграмм

	//fetcher = fetcher.New() - получает события от клиента и передает их процессору

	//processor = processor.New() - обрабатывает событие и формирует ответ

	// consumer.Start(fetcher, processor) - получает и обрабатывает события
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token

}
