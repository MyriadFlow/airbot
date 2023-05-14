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
			Name:        "reply",
			Description: "test command",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "option",
					Description: "select the number of image",
					Required:    true,
				},
			},
		},
	}
	return commands
}
