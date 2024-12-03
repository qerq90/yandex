package client

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"qerq90/yandex/model/offer"
	offermappings "qerq90/yandex/model/offer_mappings"
	"qerq90/yandex/model/order"
	"qerq90/yandex/model/product"
	"qerq90/yandex/model/stock"
	"qerq90/yandex/model/warehouse"
	"strconv"
	"strings"
	"time"
)

type YandexMarketClient struct {
	apiKey     string
	campaignId string
}

type warehouseId = int
type warehouseName = string
type offerID = string
type offerName = string

const (
	YANDEX_FORMAT = "02-01-2006"
)

var (
	offerInfoById = make(map[offerID]product.Product)
)

func MakeYandexMarketClient() (*YandexMarketClient, error) {
	apiKey, nonEmpty := os.LookupEnv("YANDEX_API_KEY")
	if !nonEmpty {
		return nil, errors.New("no YANDEX_API_KEY in ENV")
	}

	return &YandexMarketClient{apiKey: apiKey, campaignId: "56473059"}, nil
}

func (c *YandexMarketClient) GetOrders() []product.OfferProducts {
	orders := c.getOrders()

	offers := []product.OfferProducts{}

	for _, order := range orders {
		newOfferProducts := product.OfferProducts{Id: strconv.Itoa(order.ID)}
		offerProducts := []product.Product{}

		for _, item := range order.Items {
			for i := 0; i < item.Count; i++ {
				offerProduct := product.Product{Id: item.OfferID, Name: item.OfferName, Img: "blablabla", Status: order.Status}
				offerProducts = append(offerProducts, offerProduct)
			}
		}

		newOfferProducts.Products = offerProducts
		offers = append(offers, newOfferProducts)
	}

	return offers
}

func (c *YandexMarketClient) getOrders() []order.Order {
	resp, err := c.makeRequestWithAuth(http.MethodGet, "https://api.partner.market.yandex.ru/campaigns/"+c.campaignId+"/orders?fromDate="+time.Now().Format(YANDEX_FORMAT)+"&toDate="+time.Now().Format(YANDEX_FORMAT), nil)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ordersResponse = order.Root{}
	err = json.Unmarshal(data, &ordersResponse)
	if err != nil {
		log.Fatal(err)
	}

	return ordersResponse.Orders
}

func (c *YandexMarketClient) makeRequestWithAuth(method string, url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Api-Key", c.apiKey)
	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *YandexMarketClient) GetWarehouseMapping() map[warehouseId]warehouseName {
	resp, err := c.makeRequestWithAuth(http.MethodGet, "https://api.partner.market.yandex.ru/warehouses", nil)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	warehouseAnswer := warehouse.WarehouseAnswer{}
	err = json.Unmarshal(data, &warehouseAnswer)
	if err != nil {
		log.Fatal(err)
	}

	warehouseMapping := make(map[warehouseId]warehouseName)

	for i := 0; i < len(warehouseAnswer.Result.Warehouses); i++ {
		warehouse := warehouseAnswer.Result.Warehouses[i]
		warehouseMapping[warehouse.Id] = warehouse.Name
	}

	return warehouseMapping
}

func (c *YandexMarketClient) GetOfferMapping() map[offerID]offerName {
	resp, err := c.makeRequestWithAuth(http.MethodPost, "https://api.partner.market.yandex.ru/businesses/58195411/offer-mappings", nil)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := offermappings.OfferMappingsResult{}
	json.Unmarshal(data, &response)

	offerMappings := response.Result.OfferMappings
	offerMapping := make(map[offerID]offerName)

	for i := 0; i < len(offerMappings); i++ {
		offerId := offerMappings[i].Offer.OfferId
		offerName := offerMappings[i].Offer.Name

		offerMapping[offerId] = offerName
	}

	return offerMapping
}

func (c *YandexMarketClient) GetWarehouses() []offer.Warehouse {
	resp, err := c.makeRequestWithAuth(http.MethodPost, "https://api.partner.market.yandex.ru/businesses/58195411/offer-mappings", nil)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := offermappings.OfferMappingsResult{}
	json.Unmarshal(data, &response)

	offerMappings := response.Result.OfferMappings
	offerIds := make([]string, 0)
	offerMapping := make(map[offerID]offerName)

	for i := 0; i < len(offerMappings); i++ {
		offerId := offerMappings[i].Offer.OfferId
		offerName := offerMappings[i].Offer.Name

		offerIds = append(offerIds, offerId)
		offerMapping[offerId] = offerName
	}

	stockRequestBody := stock.StockRequest{WithTurnover: true, OfferIds: offerIds}
	marshalled, err := json.Marshal(stockRequestBody)
	if err != nil {
		log.Fatal(err)
	}
	jsonBody := io.NopCloser(strings.NewReader(string(marshalled)))

	resp, err = c.makeRequestWithAuth(http.MethodPost, "https://api.partner.market.yandex.ru/campaigns/"+c.campaignId+"/offers/stocks", jsonBody)
	if err != nil {
		log.Fatal(err)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := offer.WarehouseStocks{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Fatal(err)
	}

	return res.Result.Warehouses
}
