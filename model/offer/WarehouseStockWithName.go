package offer

type Warehouses struct {
	Warehouses []WarehouseWithName `json:"warehouses"`
}

type WarehouseWithName struct {
	WarehouseName string  `json:"warehouseId"`
	Offers        []Offer `json:"offers"`
}
