package repository

import (
	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Offer interface {
	Create(input *model.Offer) error
	GetByTuple(sellerId, offerId int) (model.Offer, error)
	GetAllByParams(sellerId, offerId int, substr string) ([]*model.Offer, error)
	Update(input *model.Offer) error
	Delete(sellerId, offerId int) error
}

type Repository struct {
	Offer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Offer: NewOfferPostgres(db),
	}
}
