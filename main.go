package main

import (
	"github.com/MyriadFlow/airbot/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app.Init()
}
