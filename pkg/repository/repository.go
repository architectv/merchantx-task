package repository

import (
	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Offer interface {
	Get(sellerId, offerId int, substr string) ([]model.Offer, error)
	GetByTuple(sellerId, offerId int) (model.Offer, error)
	Create(input *model.Offer) error
	Delete(sellerId, offerId int) error
	Update(input *model.Offer) error
}

type Repository struct {
	Offer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Offer: NewOfferPostgres(db),
	}
}
