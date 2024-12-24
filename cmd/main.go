package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shayan0v0n/telegram-boilterplate-go.git/pkg/telegram"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	t := os.Getenv("TELEGRAM_APITOKEN")

	client := telegram.NewClient(t)

	for {
		updates, err := client.GetUpdates()
		if err != nil {
			log.Printf("Error fetching updates: %v", err)
			continue
		}

		for _, update := range updates {
			if update.Message != nil {
				log.Printf("Received message: %s", update.Message.Text)
				err = client.SendMessage(update.Message.Chat.ID, "Hello! You sent: "+update.Message.Text)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}

			}
		}
	}
}
