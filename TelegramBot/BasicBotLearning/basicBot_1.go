package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Create bot instance
	bot, err := tgbotapi.NewBotAPI("7916383261:AAHPQLeStspOG6rSHpjXmjAMzjSDSOC1LvA")
	if err != nil {
		log.Panic(err)
	}

	// Log bot username
	log.Printf("Bot started as: %s", bot.Self.UserName)

	// Listen for messages
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // Check if it's a message
			if update.Message.Text == "/exit" { // Check for the exit command
				log.Println("Exiting bot...")
				os.Exit(0) // Exit the program
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.FirstName+"!")
			bot.Send(msg) // Send the response
		}
	}
}
