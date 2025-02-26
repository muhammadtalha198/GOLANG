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
			handleMessage(bot, update.Message)
		}
	}
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if message.Text == "/start" {
		startBot(bot, message)
		handlecommands(bot, message)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Please start the bot with this cmd! /start")
		_, err := bot.Send(msg) // Send the response

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}

func startBot(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome! Use /help to see available commands.")
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}

func handlecommands(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	switch message.Text {

	case "hello":

		msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, "+message.From.FirstName+"!")
		_, err := bot.Send(msg) // Send the response
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}

	case "/help":

		msg := tgbotapi.NewMessage(message.Chat.ID, "Available commands:\n/exit - exit the bot\n/about - About the bot\n/hello - say hello to you\n")
		_, err := bot.Send(msg) // Send the response

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}

	default:

		msg := tgbotapi.NewMessage(message.Chat.ID, "This is a simple Telegram bot written in Go!")
		_, err := bot.Send(msg) // Send the response

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}

}
