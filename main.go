package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"qerq90/yandex/logic/client"
	"qerq90/yandex/logic/db"
	"qerq90/yandex/logic/service"
	"qerq90/yandex/logic/service/bot"
	"qerq90/yandex/logic/service/sender"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func sendNotifications(nc service.NotificationService) {
	for {
		nc.SendNotificationsFromYandexMarket(1)
		fmt.Println("Sleeping for 5 minutes")
		time.Sleep(time.Minute * 5)
	}
}

func makeTelegramApi() (*tgbotapi.BotAPI, error) {
	telegramToken, exists := os.LookupEnv("TELEGRAM_TOKEN")
	if !exists {
		return nil, errors.New("no DB_NAME found in .env file")
	}

	api, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		return nil, err
	}

	return api, nil
}

func main() {
	yandexClient, err := client.MakeYandexMarketClient()
	if err != nil {
		log.Fatal(err)
	}

	vkClient, err := client.MakeVkClient()
	if err != nil {
		log.Fatal(err)
	}

	telegramApi, err := makeTelegramApi()
	if err != nil {
		log.Fatal(err)
	}

	telegramClient := client.MakeTelegramClient(telegramApi)

	dao, err := db.CreateDao()
	if err != nil {
		log.Fatal(err)
	}

	vkSender := sender.MakeVkSender(vkClient, dao)
	telegramSender := sender.MakeTelegramSender(telegramClient, dao)
	ncServiceVk := service.MakeNcService(yandexClient, vkSender)
	ncServiceTelegram := service.MakeNcService(yandexClient, telegramSender)

	telegramBot := bot.MakeTelegramBot(telegramApi, dao)

	go sendNotifications(*ncServiceVk)
	go sendNotifications(*ncServiceTelegram)
	go telegramBot.Run()

	select {}

}
