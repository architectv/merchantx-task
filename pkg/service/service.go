package service

import (
	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/architectv/merchantx-task/pkg/repository"
)

type Offer interface {
	GetAllByParams(sellerId, offerId int, substr string) ([]*model.Offer, error)
	PutWithFile(sellerId, statId int, filename string) error
	GetStat(id int) (*model.Statistics, error)
	CreateStat() (int, error)
	ErrorStat(id int, status string) error
}

type Service struct {
	Offer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Offer: NewOfferService(repos.Offer),
	}
}
