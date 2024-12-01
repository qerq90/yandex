package warehouse

type WarehouseAnswer struct {
	Result Result `json:"result"`
}

type Result struct {
	Warehouses []Warehouse `json:"warehouses"`
}

type Warehouse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
