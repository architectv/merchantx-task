package service

import (
	"errors"
	"testing"

	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/architectv/merchantx-task/pkg/repository"
	mock_repository "github.com/architectv/merchantx-task/pkg/repository/mocks"
	"github.com/architectv/merchantx-task/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestOfferService_GetAllByParams(t *testing.T) {
	type args struct {
		sellerId int
		offerId  int
		substr   string
	}
	type mockBehavior func(r *mock_repository.MockOffer, args args)

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
				substr:   "test",
			},
			mock: func(r *mock_repository.MockOffer, args args) {
				r.EXPECT().GetAllByParams(args.sellerId, args.offerId, args.substr).
					Return(nil, nil)
			},
			wantErr: false,
		},
		{
			name: "Wrong Data",
			input: args{
				sellerId: -1,
				offerId:  1,
				substr:   "test",
			},
			mock: func(r *mock_repository.MockOffer, args args) {
				r.EXPECT().GetAllByParams(args.sellerId, args.offerId, args.substr).
					Return(nil, errors.New("some error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockOffer(c)
			tt.mock(repo, tt.input)
			s := &OfferService{repo: repo}

			_, err := s.GetAllByParams(tt.input.sellerId, tt.input.offerId, tt.input.substr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOfferService_PutWithFile(t *testing.T) {
	const prefix = "../../"
	db, err := test.PrepareTestDatabase(prefix)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer db.Close()

	type args struct {
		sellerId int
		filename string
	}

	tests := []struct {
		name    string
		input   args
		want    *model.Statistics
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				sellerId: 1,
				filename: prefix + test.TestDir + "ok.xlsx",
			},
			want: &model.Statistics{
				CreateCount: 5,
				UpdateCount: 1,
				DeleteCount: 1,
				ErrorCount:  5,
			},
			wantErr: false,
		},
		{
			name: "Bad File",
			input: args{
				sellerId: 1,
				filename: prefix + test.TestDir + "bad_file",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewOfferPostgres(db)
			s := &OfferService{repo: repo}

			got, err := s.PutWithFile(tt.input.sellerId, tt.input.filename)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
