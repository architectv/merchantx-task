package model

type Offer struct {
	ProductId int
	SellerId  int
	OfferId   int
	Price     int
	Quantity  int
	Available bool
}
