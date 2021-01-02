package model

type Offer struct {
	SellerId  int    `json:"id" db:"seller_id"`
	OfferId   int    `json:"offer_id" db:"offer_id"`
	Name      string `json:"name" db:"name"`
	Price     int    `json:"price" db:"price"`
	Quantity  int    `json:"quantity" db:"quantity"`
	Available bool   `json:"-" db:"available"`
}
