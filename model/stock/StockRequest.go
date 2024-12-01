package stock

type StockRequest struct {
	WithTurnover bool     `json:"withTurnover,omitempty"`
	Archived     bool     `json:"archived,omitempty"`
	OfferIds     []string `json:"offerIds"`
}
