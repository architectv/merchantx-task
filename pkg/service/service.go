package service

import (
	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/architectv/merchantx-task/pkg/repository"
)

type Offer interface {
	GetAllByParams(sellerId, offerId int, substr string) ([]*model.Offer, error)
	PutWithFile(sellerId int, filename string) (*model.Statistics, error)
}

type Service struct {
	Offer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Offer: NewOfferService(repos.Offer),
	}
}
