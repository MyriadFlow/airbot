package app

import (
	"github.com/MyriadFlow/airbot/app/bot"
	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load()

	bot.BotInit()

}
