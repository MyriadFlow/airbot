package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/MyriadFlow/airbot/app"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

func main() {
	godotenv.Load()
	sess := app.Init()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	sess.Close()
}
