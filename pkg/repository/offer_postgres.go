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

func (r *OfferPostgres) Create(input *model.Offer) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (seller_id, offer_id, name, price, quantity, available)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (seller_id, offer_id)
		DO UPDATE SET name=$3, price=$4, quantity=$5, available=$6`, offersTable)
	_, err := r.db.Exec(query, input.SellerId, input.OfferId, input.Name,
		input.Price, input.Quantity, input.Available)

	return err
}

func (r *OfferPostgres) GetByTuple(sellerId, offerId int) (model.Offer, error) {
	var offer model.Offer
	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE seller_id=$1 AND offer_id=$2 AND available=true`,
		offersTable)
	err := r.db.Get(&offer, query, sellerId, offerId)

	return offer, err
}

func (r *OfferPostgres) GetAllByParams(sellerId, offerId int, substr string) ([]*model.Offer, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if sellerId != 0 {
		setValues = append(setValues, fmt.Sprintf("seller_id=$%d", argId))
		args = append(args, sellerId)
		argId++
	}

	if offerId != 0 {
		setValues = append(setValues, fmt.Sprintf("offer_id=$%d", argId))
		args = append(args, offerId)
		argId++
	}

	if substr != "" {
		setValues = append(setValues, fmt.Sprintf("name ILIKE $%d", argId))
		substr = "%" + substr + "%"
		args = append(args, substr)
		argId++
	}

	setValues = append(setValues, fmt.Sprintf("available = true"))
	setQuery := "WHERE " + strings.Join(setValues, " AND ")

	values := "seller_id, offer_id, name, price, quantity"
	query := fmt.Sprintf(`SELECT %s FROM %s %s`, values, offersTable, setQuery)

	var offers []*model.Offer
	err := r.db.Select(&offers, query, args...)
	if err != nil {
		return nil, err
	}

	return offers, nil
}

func (r *OfferPostgres) Update(input *model.Offer) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Price != 0 {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, input.Price)
		argId++
	}

	if input.Quantity != 0 {
		setValues = append(setValues, fmt.Sprintf("quantity=$%d", argId))
		args = append(args, input.Quantity)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE seller_id=$%d AND offer_id=$%d`,
		offersTable, setQuery, argId, argId+1)

	args = append(args, input.SellerId, input.OfferId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *OfferPostgres) Delete(sellerId, offerId int) error {
	query := fmt.Sprintf(
		`UPDATE %s SET available=false WHERE seller_id=$1 AND offer_id=$2`,
		offersTable)
	_, err := r.db.Exec(query, sellerId, offerId)

	return err
}
