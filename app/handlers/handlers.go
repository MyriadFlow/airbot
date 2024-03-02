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

	defer sess.Close()
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
		if args[1] == "upscaleSubtle" {
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
				UpscaleSubtle(number, repliedMessageID, imageID, sess_id, nonce)
			}
		}
		if args[1] == "upscaleCreative" {
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
				UpscaleCreative(number, repliedMessageID, imageID, sess_id, nonce)
			}
		}
		if args[1] == "gpt" {
			parts := strings.SplitN(m.Content, " ", 3)
			if len(parts) < 3 {
				s.ChannelMessageSend(m.ChannelID, "Invalid format. Usage: !airbot gpt <prompt>")
				return
			}
			prompt := parts[2]

			res, err := chatgpt.GetChatGPTResponse(prompt)
			if err != nil {
				fmt.Println("Error generating response:", err.Error())
				s.ChannelMessageSend(m.ChannelID, "Error generating response.")
				return
			}
			fmt.Println("res", res)
			// Truncate the response if it exceeds Discord's maximum message length
			if len(res) > 2000 {
				res = res[:2000]
			}
			reply := &discordgo.MessageReference{
				MessageID: m.ID,
			}
			_, err = s.ChannelMessageSendReply(m.ChannelID, res, reply)
			if err != nil {
				fmt.Println("Error sending message reply:", err.Error())
				return
			}
		}

		if args[1] == "help" {
			const helpMessage = "Available commands:\n" +
				"1. /generate <prompt>: Generates text based on the provided prompt.\n" +
				"2. /gpt <prompt>: Generates a response using GPT based on the provided prompt.\n" +
				"3. !airbot upscale <number> (in reply to an image): Upscales the replied image by the specified factor.\n" +
				"4. !airbot variation <number> (in reply to an image): Creates variations of the replied image.\n" +
				"5. !airbot subtle (in reply to an image): Creates a subtly varied version of the replied image.\n" +
				"6. !airbot region (in reply to an image): Creates a regionally varied version of the replied image.\n" +
				"7. !airbot strong (in reply to an image): Creates a strongly varied version of the replied image.\n" +
				"8. !airbot upscaleSubtle (in reply to an image): Upscales the replied image subtly.\n" +
				"9. !airbot upscaleCreative (in reply to an image): Upscales the replied image creatively."

			reply := &discordgo.MessageReference{
				MessageID: m.ID,
			}
			s.ChannelMessageSendReply(m.ChannelID, helpMessage, reply)
		}
	})
}
