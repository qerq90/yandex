package product

type OfferProducts struct {
	Id       string
	Products []Product
	IsLocal  bool
}

type Product struct {
	Id     string
	Name   string
	Img    string
	Status string
}
