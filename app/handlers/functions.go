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

func Generate(prompt string, sess_id string, nonce string) {
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 2,
		"application_id": "936929561302675456",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"session_id": "` + sess_id + `",
		"data": {
			"version": "1166847114203123795",
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
				"type":1,
				"application_id": "936929561302675456",
				"version": "1166847114203123795",
				"name": "imagine",
				"description": "Create images with Midjourney",
				"options": [{
					"type": 3,
					"name": "prompt",
					"description": "The prompt to imagine",
					"required": true,
					"description_localized": "The prompt to imagine",
					"name_localized": "prompt"		  
				}],
				"integration_types": [
			    	0
		      ],
			  "description_localized": "Create images with Midjourney",
			  "name_localized": "imagine"		
			},
			"attachments": []
		},
		"nonce": "` + nonce + `",
		"analytics_location": "slash_ui"
	}`
	fmt.Println("request json:", jsonStr)
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

func Upscale(number int, messageid string, imageid string, sess_id string, nonce string) error {
	numberString := strconv.Itoa(number)
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"nonce": "` + nonce + `",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "` + sess_id + `",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::JOB::upsample::` + numberString + `::` + imageid + `"
		}
	}`
	fmt.Println("request json:", jsonStr)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("authorization", user_token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil

}
func Variation(number int, messageid string, imageid string, sess_id string, nonce string) {
	numberString := strconv.Itoa(number)
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"nonce": "` + nonce + `",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "` + sess_id + `",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::JOB::variation::` + numberString + `::` + imageid + `"
		}
	}`
	fmt.Println("request json:", jsonStr)
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

func VarySubtle(number int, messageid string, imageid string, sess_id string, nonce string) {
	numberString := strconv.Itoa(number)
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"nonce": "` + nonce + `",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "` + sess_id + `",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::JOB::low_variation::` + numberString + `::` + imageid + `::SOLO"
		}
	}`
	fmt.Println("request json:", jsonStr)
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

func VaryStrong(number int, messageid string, imageid string, sess_id string, nonce string) {
	numberString := strconv.Itoa(number)
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"nonce": "` + nonce + `",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "` + sess_id + `",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::JOB::high_variation::` + numberString + `::` + imageid + `::SOLO"
		}
	}`
	fmt.Println("request json:", jsonStr)
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
func VaryRegion(number int, messageid string, imageid string, sess_id string, nonce string) {
	numberString := strconv.Itoa(number)
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"nonce": "` + nonce + `",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "` + sess_id + `",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::Inpaint::` + numberString + `::` + imageid + `::SOLO"
		}
	}`
	fmt.Println("request json:", jsonStr)
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

func UpscaleSubtle(number int, messageid string, imageid string, sess_id string, nonce string) {
	numberString := strconv.Itoa(number)
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"nonce": "` + nonce + `",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "` + sess_id + `",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::JOB::upsample_v6_2x_subtle::` + numberString + `::` + imageid + `::SOLO"
		}
	}`
	fmt.Println("request json:", jsonStr)
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

func UpscaleCreative(number int, messageid string, imageid string, sess_id string, nonce string) {
	numberString := strconv.Itoa(number)
	url := "https://discord.com/api/v9/interactions"
	server_id := os.Getenv("SERVER_ID")
	user_token := os.Getenv("USER_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")
	jsonStr := `{
		"type": 3,
		"nonce": "` + nonce + `",
		"guild_id": "` + server_id + `",
		"channel_id": "` + channel_id + `",
		"message_flags": 0,
		"message_id": "` + messageid + `",
		"application_id": "936929561302675456",
		"session_id": "` + sess_id + `",
		"data": {
			"component_type": 2,
			"custom_id": "MJ::JOB::upsample_v6_2x_creative::` + numberString + `::` + imageid + `::SOLO"
		}
	}`
	fmt.Println("request json:", jsonStr)
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
