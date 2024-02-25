package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/MyriadFlow/airbot/app/commands"
	"github.com/MyriadFlow/airbot/utils/chatgpt"
	"github.com/bwmarrin/discordgo"
)

func AddHandlers(sess *discordgo.Session) {
	commands := commands.RegisterCommands()
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	application_id := os.Getenv("APPLICATION_ID")
	for i, v := range commands {
		cmd, err := sess.ApplicationCommandCreate(application_id, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
		fmt.Println("command registered: ", cmd.Name)
	}
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"generate": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			margs := make([]string, 0, len(options))
			msgformat := "Take a look at the value(s) you entered:\n"

			if option, ok := optionMap["prompt"]; ok {
				margs = append(margs, option.StringValue())
				prompt := strings.Join(margs[:], " ")
				pprmptTrimmed := strings.ReplaceAll(prompt, "\n", " ")
				sess_id := s.State.SessionID
				nonce := fmt.Sprint(rand.Int())
				Generate(pprmptTrimmed, sess_id, nonce)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: msgformat + prompt,
					},
				})

			}
		},
		"gpt": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			margs := make([]string, 0, len(options))
			msgformat := "Take a look at your response:\n"

			if option, ok := optionMap["prompt"]; ok {
				margs = append(margs, option.StringValue())
				prompt := strings.Join(margs[:], " ")
				res, err := chatgpt.GetChatGPTResponse(prompt)
				if err != nil {
					fmt.Println("error in generating response:", err.Error())
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "error generating response",
						},
					})
				}
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: msgformat + res,
					},
				})
			}
		},
	}
	const prefix = "!airbot"
	sess.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		args := strings.Split(m.Content, " ")

		if args[0] != prefix {
			return
		}
		if args[1] == "upscale" {
			if m.MessageReference != nil {
				repliedMessageID := m.MessageReference.MessageID
				imageURL, _, err := getImageFromMessageID(s, os.Getenv("CHANNEL_ID"), repliedMessageID)
				imageID := getImageId(imageURL)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				number, _ := strconv.Atoi(args[2])
				sess_id := s.State.SessionID
				nonce := fmt.Sprint(rand.Int())
				Upscale(number, repliedMessageID, imageID, sess_id, nonce)
			}
		}
		if args[1] == "variation" {
			if m.MessageReference != nil {
				repliedMessageID := m.MessageReference.MessageID
				imageURL, _, err := getImageFromMessageID(s, os.Getenv("CHANNEL_ID"), repliedMessageID)
				imageID := getImageId(imageURL)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				number, _ := strconv.Atoi(args[2])
				sess_id := s.State.SessionID
				nonce := fmt.Sprint(rand.Int())
				Variation(number, repliedMessageID, imageID, sess_id, nonce)
			}
		}
		if args[1] == "subtle" {
			if m.MessageReference != nil {
				repliedMessageID := m.MessageReference.MessageID
				imageURL, _, err := getImageFromMessageID(s, os.Getenv("CHANNEL_ID"), repliedMessageID)
				imageID := getImageId(imageURL)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				number := 1
				sess_id := s.State.SessionID
				nonce := fmt.Sprint(rand.Int())
				VarySubtle(number, repliedMessageID, imageID, sess_id, nonce)
			}
		}
		if args[1] == "region" {
			if m.MessageReference != nil {
				repliedMessageID := m.MessageReference.MessageID
				imageURL, _, err := getImageFromMessageID(s, os.Getenv("CHANNEL_ID"), repliedMessageID)
				imageID := getImageId(imageURL)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				number := 1
				sess_id := s.State.SessionID
				nonce := fmt.Sprint(rand.Int())
				VaryRegion(number, repliedMessageID, imageID, sess_id, nonce)
			}
		}

		if args[1] == "strong" {
			if m.MessageReference != nil {
				repliedMessageID := m.MessageReference.MessageID
				imageURL, _, err := getImageFromMessageID(s, os.Getenv("CHANNEL_ID"), repliedMessageID)
				imageID := getImageId(imageURL)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				number := 1
				sess_id := s.State.SessionID
				nonce := fmt.Sprint(rand.Int())
				VaryStrong(number, repliedMessageID, imageID, sess_id, nonce)
			}
		}
		if args[1] == "upscale2x" {
			if m.MessageReference != nil {
				repliedMessageID := m.MessageReference.MessageID
				imageURL, _, err := getImageFromMessageID(s, os.Getenv("CHANNEL_ID"), repliedMessageID)
				imageID := getImageId(imageURL)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				number := 1
				sess_id := s.State.SessionID
				nonce := fmt.Sprint(rand.Int())
				Upscale2x(number, repliedMessageID, imageID, sess_id, nonce)
			}
		}
		if args[1] == "upscale4x" {
			if m.MessageReference != nil {
				repliedMessageID := m.MessageReference.MessageID
				imageURL, _, err := getImageFromMessageID(s, os.Getenv("CHANNEL_ID"), repliedMessageID)
				imageID := getImageId(imageURL)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				number := 1
				sess_id := s.State.SessionID
				nonce := fmt.Sprint(rand.Int())
				Upscale4x(number, repliedMessageID, imageID, sess_id, nonce)
			}
		}

		if args[1] == "maxupscale" {
			if m.MessageReference != nil {
				repliedMessageID := m.MessageReference.MessageID
				imageURL, _, err := getImageFromMessageID(s, os.Getenv("CHANNEL_ID"), repliedMessageID)
				imageID := getImageId(imageURL)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				sess_id := s.State.SessionID
				UpscaleMax(repliedMessageID, imageID, sess_id)
			}
		}

	})
}
