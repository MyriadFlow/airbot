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
			Name:        "gpt",
			Description: "Generate text with gpt-4",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "prompt",
					Description: "prompt to generate text",
					Required:    true,
				},
			},
		},
		{
			Name:        "upscale",
			Description: "Upscale one of the generated image",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "choice",
					Description: "choice of image to upscale",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "1",
							Value: "1",
						},
						{
							Name:  "2",
							Value: "2",
						},
						{
							Name:  "3",
							Value: "3",
						},
						{
							Name:  "4",
							Value: "4",
						},
					},
				},
			},
		},
	}
	return commands
}
