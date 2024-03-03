package commands

import "github.com/bwmarrin/discordgo"

func RegisterCommands() []*discordgo.ApplicationCommand {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "generate",
			Description: "Generate Image with given prompt",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "prompt",
					Description: "prompt to generate image",
					Required:    true,
				},
			},
		},
		{
			Name:        "help",
			Description: "Display available commands",
		},
	}
	return commands
}
