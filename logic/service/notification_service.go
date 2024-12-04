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

	err := nc.vkClient.SendMessage(message, vkId, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getStatus(status string) string {
	if status == "CANCELLED" {
		return "Отмена"
	} else {
		return "Создание"
	}
}
