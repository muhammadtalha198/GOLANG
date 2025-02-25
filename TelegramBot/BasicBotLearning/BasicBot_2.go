package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	// Load bot token from environment variable
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Panic("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	log.Println("botToken: ", botToken)

	// fmt.Println("botToken: ", botToken)

	// Create bot instance
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// Retrieve bot username
	botUsername := bot.Self.UserName // Get the bot username

	// Log bot username
	log.Println("Bot started as:", botUsername)

	// Listen for messages
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // Check if it's a message
			handlecommands(bot, update.Message)
		}
	}
}

func handlecommands(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	if message.Text == "/start" {

		msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to the bot! Use /help to see available commands.")
		_, err := bot.Send(msg) // Send the response

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}

	} else if message.Text == "/help" {

		msg := tgbotapi.NewMessage(message.Chat.ID, "Available commands:\n/start - Welcome message\n/help - Show this message\n/about - About the bot")
		_, err := bot.Send(msg) // Send the response

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}

	} else if message.Text == "/about" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "This is a simple Telegram bot written in Go!")
		_, err := bot.Send(msg) // Send the response

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}

	} else if message.Text == "/about" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "This is a simple Telegram bot written in Go!")
		_, err := bot.Send(msg) // Send the response

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}

	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, "+message.From.FirstName+"!")
		_, err := bot.Send(msg) // Send the response
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}
