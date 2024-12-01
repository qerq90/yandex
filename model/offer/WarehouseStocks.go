package offer

import "qerq90/yandex/model/util"

type WarehouseStocks struct {
	Result WarehouseResult `json:"result"`
}

type WarehouseResult struct {
	Warehouses []Warehouse `json:"warehouses"`
}

type Warehouse struct {
	WarehouseId int     `json:"warehouseId"`
	Offers      []Offer `json:"offers"`
}

type Offer struct {
	OfferId         string          `json:"offerId"`
	TurnOverSummary TurnOverSummary `json:"turnoverSummary"`
	Stocks          []Stock         `json:"stocks"`
}

type TurnOverSummary struct {
	Turnover     string           `json:"turnover"`
	TurnoverDays util.JsonFloat64 `json:"turnoverDays"`
}

type Stock struct {
	StockType string `json:"type"`
	Count     int    `json:"count"`
}
