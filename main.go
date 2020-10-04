package main

import (
	"os"

	"github.com/Acksell/among-us-queue/bot"
)

var token string

func init() {
	token = os.Getenv("BOT_TOKEN")
}

func main() {
	bot.MakeAndListen(token)
}
