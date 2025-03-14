package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MyriadFlow/airbot/app/handlers"
	"github.com/bwmarrin/discordgo"
)

func Init() {
	bot_token := os.Getenv("BOT_TOKEN")

	sess, err := discordgo.New("Bot " + bot_token)
	if err != nil {
		log.Fatal(err)
	}
	// Add all handlers
	handlers.AddHandlers(sess)
	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session.
	fmt.Println("shutting down")
	
}
