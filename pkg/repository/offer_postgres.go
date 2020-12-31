package repository

import (
	"fmt"
	"strings"

	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/jmoiron/sqlx"
)

type OfferPostgres struct {
	db *sqlx.DB
}

func NewOfferPostgres(db *sqlx.DB) *OfferPostgres {
	return &OfferPostgres{db: db}
}

func (r *OfferPostgres) Get(sellerId, offerId int, substr string) ([]model.Offer, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if sellerId != 0 {
		setValues = append(setValues, fmt.Sprintf("s.id=$%d", argId))
		args = append(args, sellerId)
		argId++
	}

	if offerId != 0 {
		setValues = append(setValues, fmt.Sprintf("o.offer_id=$%d", argId))
		args = append(args, offerId)
		argId++
	}

	if substr != "" {
		setValues = append(setValues, fmt.Sprintf("p.name ILIKE '%%$%d%%'", argId))
		args = append(args, substr)
		argId++
	}

	setValues = append(setValues, fmt.Sprintf("o.available = true"))
	setQuery := "WHERE " + strings.Join(setValues, " AND ")

	values := "seller_id, offer_id, name, price, quantity"
	query := fmt.Sprintf(`SELECT %s FROM %s s JOIN %s o ON s.id = o.seller_id
		JOIN %s p ON p.id = o.product_id %s;`,
		values, sellersTable, offersTable, productsTable, setQuery)

	var offers []model.Offer
	err := r.db.Select(&offers, query, args...)
	if err != nil {
		return nil, err
	}

	return offers, nil
}
