package bot

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func BotInit() *discordgo.Session {
	bot_token := os.Getenv("BOT_TOKEN")

	sess, err := discordgo.New("Bot " + bot_token)
	if err != nil {
		log.Fatal(err)
	}
	return sess
}
