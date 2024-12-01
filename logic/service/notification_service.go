package service

import (
	"fmt"
	"qerq90/yandex/logic/client"
	"slices"
)

var (
	alreadyProcessed []string
)

type NotificationService struct {
	yandexClient *client.YandexMarketClient
	vkClient     *client.VkClient
}

func MakeNcService(yandexClient *client.YandexMarketClient, vkClient *client.VkClient) *NotificationService {
	return &NotificationService{yandexClient: yandexClient, vkClient: vkClient}
}

func (nc *NotificationService) SendNotificationsFromYandexMarket() {
	orders := nc.yandexClient.GetOrders()

	message := ""
	for _, order := range orders {
		if slices.Index(alreadyProcessed, order.Id) == -1 {
			message += "New Order №" + order.Id + "\n"
			for _, item := range order.Products {
				message += fmt.Sprintf("%s \n\n", item.Name)
			}

			alreadyProcessed = append(alreadyProcessed, order.Id)
		}
	}

	nc.vkClient.SendMessage(message, 51422811, nil)
}
