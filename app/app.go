package app

import (
	"fmt"
	"log"
	"os"

	"github.com/MyriadFlow/airbot/app/handlers"
	"github.com/bwmarrin/discordgo"
)

func Init() *discordgo.Session {
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
	defer sess.Close()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	return sess

}
