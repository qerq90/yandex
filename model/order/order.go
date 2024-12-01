package order

// Root represents the top-level structure of the JSON.
type Root struct {
	Pager  Pager   `json:"pager"`
	Orders []Order `json:"orders"`
	Paging Paging  `json:"paging"`
}

// Pager holds pagination information.
type Pager struct {
	Total       int `json:"total"`
	From        int `json:"from"`
	To          int `json:"to"`
	CurrentPage int `json:"currentPage"`
	PagesCount  int `json:"pagesCount"`
	PageSize    int `json:"pageSize"`
}

// Order represents an order structure.
type Order struct {
	ID                            int       `json:"id"`
	Status                        string    `json:"status"`
	Substatus                     string    `json:"substatus"`
	CreationDate                  string    `json:"creationDate"`
	UpdatedAt                     string    `json:"updatedAt"`
	Currency                      string    `json:"currency"`
	ItemsTotal                    float64   `json:"itemsTotal"`
	DeliveryTotal                 float64   `json:"deliveryTotal"`
	BuyerItemsTotal               float64   `json:"buyerItemsTotal"`
	BuyerTotal                    float64   `json:"buyerTotal"`
	BuyerItemsTotalBeforeDiscount float64   `json:"buyerItemsTotalBeforeDiscount"`
	BuyerTotalBeforeDiscount      float64   `json:"buyerTotalBeforeDiscount"`
	PaymentType                   string    `json:"paymentType"`
	PaymentMethod                 string    `json:"paymentMethod"`
	Fake                          bool      `json:"fake"`
	Items                         []Item    `json:"items"`
	Subsidies                     []Subsidy `json:"subsidies"`
	Delivery                      Delivery  `json:"delivery"`
	Notes                         string    `json:"notes"`
	TaxSystem                     string    `json:"taxSystem"`
	CancelRequested               bool      `json:"cancelRequested"`
	ExpiryDate                    string    `json:"expiryDate"`
}

// Item represents an item in the order.
type Item struct {
	ID                       int        `json:"id"`
	OfferID                  string     `json:"offerId"`
	OfferName                string     `json:"offerName"`
	Price                    float64    `json:"price"`
	BuyerPrice               float64    `json:"buyerPrice"`
	BuyerPriceBeforeDiscount float64    `json:"buyerPriceBeforeDiscount"`
	PriceBeforeDiscount      float64    `json:"priceBeforeDiscount"`
	Count                    int        `json:"count"`
	VAT                      string     `json:"vat"`
	ShopSku                  string     `json:"shopSku"`
	Subsidy                  float64    `json:"subsidy"`
	PartnerWarehouseID       string     `json:"partnerWarehouseId"`
	Promos                   []Promo    `json:"promos"`
	Instances                []Instance `json:"instances"`
	Details                  []Detail   `json:"details"`
	Subsidies                []Subsidy  `json:"subsidies"`
	RequiredInstanceTypes    []string   `json:"requiredInstanceTypes"`
	Tags                     []string   `json:"tags"`
}

// Promo represents a promotional discount.
type Promo struct {
	Type          string  `json:"type"`
	Discount      float64 `json:"discount"`
	Subsidy       float64 `json:"subsidy"`
	ShopPromoID   string  `json:"shopPromoId"`
	MarketPromoID string  `json:"marketPromoId"`
}

// Instance represents an instance of an item.
type Instance struct {
	CIS     string `json:"cis"`
	CISFull string `json:"cisFull"`
	UIN     string `json:"uin"`
	RNPT    string `json:"rnpt"`
	GTD     string `json:"gtd"`
}

// Detail represents the details of an item.
type Detail struct {
	ItemCount  int    `json:"itemCount"`
	ItemStatus string `json:"itemStatus"`
	UpdateDate string `json:"updateDate"`
}

// Subsidy represents a financial subsidy.
type Subsidy struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

// Delivery represents the delivery information of an order.
type Delivery struct {
	ID                     string        `json:"id"`
	Type                   string        `json:"type"`
	ServiceName            string        `json:"serviceName"`
	Price                  float64       `json:"price"`
	DeliveryPartnerType    string        `json:"deliveryPartnerType"`
	Courier                Courier       `json:"courier"`
	Dates                  DeliveryDates `json:"dates"`
	Region                 Region        `json:"region"`
	Address                Address       `json:"address"`
	VAT                    string        `json:"vat"`
	DeliveryServiceID      int           `json:"deliveryServiceId"`
	LiftType               string        `json:"liftType"`
	LiftPrice              float64       `json:"liftPrice"`
	OutletCode             string        `json:"outletCode"`
	OutletStorageLimitDate string        `json:"outletStorageLimitDate"`
	DispatchType           string        `json:"dispatchType"`
	Tracks                 []Track       `json:"tracks"`
	Shipments              []Shipment    `json:"shipments"`
	Estimated              bool          `json:"estimated"`
	EACType                string        `json:"eacType"`
	EACCode                string        `json:"eacCode"`
}

// Courier represents courier details.
type Courier struct {
	FullName           string `json:"fullName"`
	Phone              string `json:"phone"`
	PhoneExtension     string `json:"phoneExtension"`
	VehicleNumber      string `json:"vehicleNumber"`
	VehicleDescription string `json:"vehicleDescription"`
}

// DeliveryDates represents the delivery date range.
type DeliveryDates struct {
	FromDate         string `json:"fromDate"`
	ToDate           string `json:"toDate"`
	FromTime         string `json:"fromTime"`
	ToTime           string `json:"toTime"`
	RealDeliveryDate string `json:"realDeliveryDate"`
}

// Region represents the delivery region.
type Region struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Parent   *Region   `json:"parent,omitempty"`
	Children []*Region `json:"children,omitempty"`
}

// Address represents the delivery address.
type Address struct {
	Country    string `json:"country"`
	Postcode   string `json:"postcode"`
	City       string `json:"city"`
	District   string `json:"district"`
	Subway     string `json:"subway"`
	Street     string `json:"street"`
	House      string `json:"house"`
	Block      string `json:"block"`
	Entrance   string `json:"entrance"`
	Entryphone string `json:"entryphone"`
	Floor      string `json:"floor"`
	Apartment  string `json:"apartment"`
	Phone      string `json:"phone"`
	Recipient  string `json:"recipient"`
	GPS        GPS    `json:"gps"`
}

// GPS represents geographical coordinates.
type GPS struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Track represents a tracking information.
type Track struct {
	TrackCode         string `json:"trackCode"`
	DeliveryServiceID int    `json:"deliveryServiceId"`
}

// Shipment represents a shipment in the delivery.
type Shipment struct {
	ID           int     `json:"id"`
	ShipmentDate string  `json:"shipmentDate"`
	ShipmentTime string  `json:"shipmentTime"`
	Tracks       []Track `json:"tracks"`
	Boxes        []Box   `json:"boxes"`
}

// Box represents a box in a shipment.
type Box struct {
	ID           int    `json:"id"`
	FulfilmentID string `json:"fulfilmentId"`
}

// Paging represents pagination information for the next page.
type Paging struct {
	NextPageToken string `json:"nextPageToken"`
}
