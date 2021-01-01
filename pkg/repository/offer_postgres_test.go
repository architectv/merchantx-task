package repository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestOfferPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewOfferPostgres(db)

	type args struct {
		offer *model.Offer
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				offer: &model.Offer{
					SellerId:  1,
					OfferId:   1,
					Name:      "test name",
					Price:     1000,
					Quantity:  15,
					Available: true,
				},
			},
			mock: func(args args) {
				input := args.offer
				mock.ExpectExec(fmt.Sprintf("INSERT INTO %s", offersTable)).
					WithArgs(input.SellerId, input.OfferId, input.Name,
						input.Price, input.Quantity, input.Available).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Wrong data",
			input: args{
				offer: &model.Offer{
					SellerId:  -1,
					OfferId:   1,
					Name:      "test name",
					Price:     1000,
					Quantity:  15,
					Available: true,
				},
			},
			mock: func(args args) {
				input := args.offer
				mock.ExpectExec(fmt.Sprintf("INSERT INTO %s", offersTable)).
					WithArgs(input.SellerId, input.OfferId, input.Name,
						input.Price, input.Quantity, input.Available).
					WillReturnError(fmt.Errorf("check constraints"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			err := r.Create(tt.input.offer)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOfferPostgres_GetByTuple(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewOfferPostgres(db)

	type args struct {
		sellerId, offerId int
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    model.Offer
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				sellerId: 1,
				offerId:  1,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"seller_id", "offer_id", "name", "price", "quantity", "available"}).
					AddRow(args.sellerId, args.offerId, "test name", 1000, 100, true)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", offersTable)).
					WithArgs(args.sellerId, args.offerId).WillReturnRows(rows)
			},
			want:    model.Offer{1, 1, "test name", 1000, 100, true},
			wantErr: false,
		},
		{
			name: "Not Found",
			input: args{
				sellerId: 1,
				offerId:  1,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"seller_id", "offer_id", "name", "price", "quantity", "available"})
				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", offersTable)).
					WithArgs(args.sellerId, args.offerId).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			got, err := r.GetByTuple(tt.input.sellerId, tt.input.offerId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestOfferPostgres_GetAllByParams(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewOfferPostgres(db)

	type args struct {
		sellerId int
		offerId  int
		substr   string
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    []*model.Offer
		wantErr bool
	}{
		{
			name: "Ok (sellerId, offerId, substr)",
			input: args{
				sellerId: 1,
				offerId:  1,
				substr:   "test",
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"seller_id", "offer_id", "name", "price", "quantity", "available"}).
					AddRow(args.sellerId, args.offerId, "test name", 1000, 100, true)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", offersTable)).
					WithArgs(args.sellerId, args.offerId, args.substr).WillReturnRows(rows)
			},
			want: []*model.Offer{
				{1, 1, "test name", 1000, 100, true},
			},
			wantErr: false,
		},
		{
			name: "Ok (sellerId, offerId)",
			input: args{
				sellerId: 1,
				offerId:  1,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"seller_id", "offer_id", "name", "price", "quantity", "available"}).
					AddRow(args.sellerId, args.offerId, "test name", 1000, 100, true)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", offersTable)).
					WithArgs(args.sellerId, args.offerId).WillReturnRows(rows)
			},
			want: []*model.Offer{
				{1, 1, "test name", 1000, 100, true},
			},
			wantErr: false,
		},
		{
			name: "Ok (offerId, substr)",
			input: args{
				offerId: 1,
				substr:  "test",
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"seller_id", "offer_id", "name", "price", "quantity", "available"}).
					AddRow(1, args.offerId, "test name", 1000, 100, true)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", offersTable)).
					WithArgs(args.offerId, args.substr).WillReturnRows(rows)
			},
			want: []*model.Offer{
				{1, 1, "test name", 1000, 100, true},
			},
			wantErr: false,
		},
		{
			name: "Ok (sellerId, substr)",
			input: args{
				sellerId: 1,
				substr:   "test",
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"seller_id", "offer_id", "name", "price", "quantity", "available"}).
					AddRow(args.sellerId, 1, "test name", 1000, 100, true)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", offersTable)).
					WithArgs(args.sellerId, args.substr).WillReturnRows(rows)
			},
			want: []*model.Offer{
				{1, 1, "test name", 1000, 100, true},
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			input: args{
				sellerId: 1,
				offerId:  1,
				substr:   "test",
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"seller_id", "offer_id", "name", "price", "quantity", "available"})
				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", offersTable)).
					WithArgs(args.sellerId, args.offerId, args.substr).WillReturnRows(rows)
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Internal Error",
			input: args{
				sellerId: 1,
				offerId:  1,
				substr:   "test",
			},
			mock: func(args args) {
				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", offersTable)).
					WithArgs(args.sellerId, args.offerId, args.substr).WillReturnError(fmt.Errorf("some error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			got, err := r.GetAllByParams(tt.input.sellerId, tt.input.offerId, tt.input.substr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestOfferPostgres_Update(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewOfferPostgres(db)

	type args struct {
		offer *model.Offer
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		wantErr bool
	}{
		{
			name: "Ok (name, price, quantity)",
			input: args{
				offer: &model.Offer{
					SellerId:  1,
					OfferId:   1,
					Name:      "test name",
					Price:     1000,
					Quantity:  15,
					Available: true,
				},
			},
			mock: func(args args) {
				input := args.offer
				mock.ExpectExec(fmt.Sprintf("UPDATE %s SET (.+) WHERE (.+)", offersTable)).
					WithArgs(input.Name, input.Price, input.Quantity, input.SellerId, input.OfferId).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Ok (name, price)",
			input: args{
				offer: &model.Offer{
					SellerId:  1,
					OfferId:   1,
					Name:      "test name",
					Price:     1000,
					Quantity:  0,
					Available: true,
				},
			},
			mock: func(args args) {
				input := args.offer
				mock.ExpectExec(fmt.Sprintf("UPDATE %s SET (.+) WHERE (.+)", offersTable)).
					WithArgs(input.Name, input.Price, input.SellerId, input.OfferId).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Ok (name, quantity)",
			input: args{
				offer: &model.Offer{
					SellerId:  1,
					OfferId:   1,
					Name:      "test name",
					Price:     0,
					Quantity:  15,
					Available: true,
				},
			},
			mock: func(args args) {
				input := args.offer
				mock.ExpectExec(fmt.Sprintf("UPDATE %s SET (.+) WHERE (.+)", offersTable)).
					WithArgs(input.Name, input.Quantity, input.SellerId, input.OfferId).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Ok (price, quantity)",
			input: args{
				offer: &model.Offer{
					SellerId:  1,
					OfferId:   1,
					Name:      "",
					Price:     1000,
					Quantity:  15,
					Available: true,
				},
			},
			mock: func(args args) {
				input := args.offer
				mock.ExpectExec(fmt.Sprintf("UPDATE %s SET (.+) WHERE (.+)", offersTable)).
					WithArgs(input.Price, input.Quantity, input.SellerId, input.OfferId).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Ok (empty input)",
			input: args{
				offer: &model.Offer{
					SellerId:  1,
					OfferId:   1,
					Name:      "",
					Price:     0,
					Quantity:  0,
					Available: true,
				},
			},
			mock: func(args args) {
				input := args.offer
				mock.ExpectExec(fmt.Sprintf("UPDATE %s SET (.+) WHERE (.+)", offersTable)).
					WithArgs(input.SellerId, input.OfferId).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			err := r.Update(tt.input.offer)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOfferPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewOfferPostgres(db)

	type args struct {
		sellerId, offerId int
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				sellerId: 1,
				offerId:  1,
			},
			mock: func(args args) {
				mock.ExpectExec(fmt.Sprintf("UPDATE %s SET (.+) WHERE (.+)", offersTable)).
					WithArgs(args.sellerId, args.offerId).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			input: args{
				sellerId: 1,
				offerId:  1,
			},
			mock: func(args args) {
				mock.ExpectExec(fmt.Sprintf("UPDATE %s SET (.+) WHERE (.+)", offersTable)).
					WithArgs(args.sellerId, args.offerId).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			err := r.Delete(tt.input.sellerId, tt.input.offerId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
