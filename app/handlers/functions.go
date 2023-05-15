package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Generate(prompt string) {
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 2,
		"application_id": "936929561302675456",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"session_id": "2fb980f65e5c9a77c96ca01f2c242cf6",
		"data": {
			"version": "1077969938624553050",
			"id": "938956540159881230",
			"name": "imagine",
			"type": 1,
			"options": [{
				"type": 3,
				"name": "prompt",
				"value": "` + prompt + `"
			}],
			"application_command": {
				"id": "938956540159881230",
				"application_id": "936929561302675456",
				"version": "1077969938624553050",
				"default_permission": true,
				"default_member_permissions": null,
				"type": 1,
				"nsfw": false,
				"name": "imagine",
				"description": "Create images with Midjourney",
				"dm_permission": true,
				"options": [{
					"type": 3,
					"name": "prompt",
					"description": "The prompt to imagine",
					"required": true
				}]
			},
			"attachments": []
		}
	}`
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authorization", user_token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func Upscale(number int, messageid string, imageid string) {
	numberString := strconv.Itoa(number)
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"nonce": "1106909041919000576",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "937a1c8132cd7ce3940aa8f59dedf961",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::JOB::upsample::` + numberString + `::` + imageid + `"
		}
	}`
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authorization", user_token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func UpscaleMax(messageid string, imageid string) {
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "1f3dbdf09efdf93d81a3a6420882c92c",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::JOB::upsample_max::1::` + imageid + `::SOLO"
		}
	}`
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authorization", user_token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func Variation(number int, messageid string, imageid string) {
	numberString := strconv.Itoa(number)
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"nonce": "1106909401836421120",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "937a1c8132cd7ce3940aa8f59dedf961",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::JOB::variation::` + numberString + `::` + imageid + `"
		}
	}`
	fmt.Println(jsonStr)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authorization", user_token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func getImageFromMessageID(session *discordgo.Session, channelID, messageID string) (string, string, error) {
	message, err := session.ChannelMessage(channelID, messageID)
	if err != nil {
		return "", "", fmt.Errorf("error retrieving message: %s", err)
	}

	var imageURL string
	var imageID string

	// Check if there are any attachments in the message
	if len(message.Attachments) > 0 {
		attachment := message.Attachments[0]
		imageURL = attachment.URL
		imageID = attachment.ID
	} else {
		// Check if there are any embeds in the message
		if len(message.Embeds) > 0 {
			embed := message.Embeds[0]

			// Check if the embed contains an image
			if len(embed.Image.URL) > 0 {
				imageURL = embed.Image.URL
			}

			// Check if the embed contains an image ID
			if len(embed.Image.ProxyURL) > 0 {
				imageID = embed.Image.ProxyURL
			}
		}
	}

	return imageURL, imageID, nil
}

func getImageId(url string) string {
	arr := strings.Split(url, "_")
	png := arr[len(arr)-1]
	imageId := strings.Split(png, ".")[0]
	return imageId
}
