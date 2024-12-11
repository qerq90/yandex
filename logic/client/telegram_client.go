package client

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramClient struct {
	api *tgbotapi.BotAPI
}

func MakeTelegramClient(api *tgbotapi.BotAPI) *TelegramClient {
	return &TelegramClient{api}
}

func (tg *TelegramClient) SendMessage(telegramId int, message string) {
	msg := tgbotapi.NewMessage(int64(telegramId), message)
	if _, err := tg.api.Send(msg); err != nil {
		log.Default().Println(err)
	}
}
