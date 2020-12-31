package repository

import "github.com/jmoiron/sqlx"

type Offer interface {
}

type Repository struct {
	Offer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
