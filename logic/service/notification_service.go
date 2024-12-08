package service

import (
	"fmt"
	"qerq90/yandex/logic/client"
	"qerq90/yandex/logic/service/sender"
	"slices"
)

var (
	alreadyProcessed []string //TODO change to db(maybe something like redis)
)

type NotificationService struct {
	yandexClient *client.YandexMarketClient
	sender       sender.Sender
}

func MakeNcService(yandexClient *client.YandexMarketClient, sender sender.Sender) *NotificationService {
	return &NotificationService{yandexClient: yandexClient, sender: sender}
}

func (nc *NotificationService) SendNotificationsFromYandexMarket(vkId int) {
	orders := nc.yandexClient.GetOrders()

	message := ""
	for _, order := range orders {
		if slices.Index(alreadyProcessed, order.Id) == -1 {
			message += "Заказ №" + order.Id
			message += fmt.Sprintf(" [%s]\n", getStatus(order.Products[0].Status))
			message += "Товары:\n"

			for _, item := range order.Products {
				message += fmt.Sprintf("%s \n\n", item.Name)
			}

			alreadyProcessed = append(alreadyProcessed, order.Id)
		}
	}

	nc.sender.Send(vkId, message)
}

func getStatus(status string) string {
	if status == "CANCELLED" {
		return "Отмена"
	} else {
		return "Создание"
	}
}
