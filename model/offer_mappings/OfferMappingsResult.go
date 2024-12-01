package offermappings

type OfferMappingsResult struct {
	Result Result `json:"result"`
}

type Result struct {
	OfferMappings []OfferMapping `json:"offerMappings"`
}

type OfferMapping struct {
	Offer Offer `json:"offer"`
}

type Offer struct {
	OfferId  string   `json:"offerId"`
	Name     string   `json:"name"`
	Pictures []string `json:"pictures"`
}
