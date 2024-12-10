package client

import (
	"errors"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramClient struct {
	api *tgbotapi.BotAPI
}

func MakeTelegramClient() (*TelegramClient, error) {
	telegramToken, exists := os.LookupEnv("TELEGRAM_TOKEN")
	if !exists {
		return nil, errors.New("no DB_NAME found in .env file")
	}

	api, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		return nil, err
	}

	return &TelegramClient{api}, nil
}

func (tg *TelegramClient) SendMessage(telegramId int, message string) {
	msg := tgbotapi.NewMessage(int64(telegramId), message)
	if _, err := tg.api.Send(msg); err != nil {
		log.Default().Println(err)
	}
}
