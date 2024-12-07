package main

import (
	"fmt"
	"log"
	"qerq90/yandex/logic/client"
	"qerq90/yandex/logic/service"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func sendNotifications(nc service.NotificationService) {
	for {
		nc.SendNotificationsFromYandexMarket(51422811)
		fmt.Println("Sleeping for 5 minutes")
		time.Sleep(time.Minute * 5)
	}
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

	ncService := service.MakeNcService(yandexClient, vkClient)

	go sendNotifications(*ncService)

	select {}

}
