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

// func getYandex(w http.ResponseWriter, _ *http.Request) {
// 	yandexClient, err := client.MakeYandexMarketClient()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	warehouses := yandexClient.GetWarehouses()
// 	offerMapping := yandexClient.GetOfferMapping()
// 	warehouseMapping := yandexClient.GetWarehouseMapping()

// 	warehousesWithName := &offer.Warehouses{}
// 	for i := 0; i < len(warehouses); i++ {
// 		warehouse := warehouses[i]

// 		for i := 0; i < len(warehouse.Offers); i++ {
// 			offer := &warehouse.Offers[i]
// 			offer.OfferId = offerMapping[offer.OfferId]
// 		}
// 		warehouseWithName := offer.WarehouseWithName{}
// 		warehouseWithName.Offers = warehouse.Offers
// 		warehouseWithName.WarehouseName = warehouseMapping[warehouse.WarehouseId]

// 		warehousesWithName.Warehouses = append(warehousesWithName.Warehouses, warehouseWithName)
// 	}

// 	result, err := json.Marshal(warehousesWithName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	io.Writer.Write(w, result)
// }

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

	for {
		ncService.SendNotificationsFromYandexMarket()
		fmt.Println("Sleeping for 5 minutes")
		time.Sleep(time.Minute * 5)
	}

}
