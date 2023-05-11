package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func AddHandlers(sess *discordgo.Session) {

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		url := "https://discord.com/api/v9/interactions"
		server_id := os.Getenv("SERVER_ID")
		user_token := os.Getenv("USER_TOKEN")
		channel_id := os.Getenv("CHANNEL_ID")

		if m.Author.ID == s.State.User.ID {
			return
		}
		prefix := "!mj"
		args := strings.Split(m.Content, " ")

		if args[0] != prefix {
			return
		}
		if args[1] == "imagine" {
			prompt := "Batman"
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
			reqbody, _ := io.ReadAll(req.Body)

			fmt.Println("request Body:", reqbody)
			fmt.Println("response Status:", resp.Status)
			fmt.Println("response Headers:", resp.Header)
			body, _ := io.ReadAll(resp.Body)
			fmt.Println("response Body:", string(body))
		}
	})
}
