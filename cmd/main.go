package main

import (
	bot "bot/bot"
)

func main() {
	bot, err := bot.New("bot.yml")
	if err != nil {
		panic(err)
	}

	bot.Run()
}
