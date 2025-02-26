package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func initializeBot() (*tgbotapi.BotAPI, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Load bot token from environment variable
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	// Create bot instance
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func main() {

	bot, err := initializeBot()
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}

	// Retrieve bot username
	botUsername := bot.Self.UserName // Get the bot username
	log.Println("Bot started as:", botUsername)

	// Listen for messages
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // Check if it's a message
			askName(bot, update.Message)
		}
	}

}

var userState = make(map[int64]string) // Store user state by chat ID

func askName(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	if message.Text == "/askname" {

		userState[message.Chat.ID] = "waiting_for_name"
		log.Println("message.Chat.ID: ", message.Chat.ID) // Log chat ID

		msg := tgbotapi.NewMessage(message.Chat.ID, "What's your name?")
		bot.Send(msg)

	} else {
		state, ok := userState[message.Chat.ID] // Check if the user has a state

		if ok { // If the user exists in the map
			if state == "waiting_for_name" { // If the user is in "waiting_for_name" state
				msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, "+message.Text+"!")
				bot.Send(msg)
				log.Printf("User %s sent: %s", message.From.UserName, message.Text)
				delete(userState, message.Chat.ID) // Clear the user's state
			}
		}
	}
}
