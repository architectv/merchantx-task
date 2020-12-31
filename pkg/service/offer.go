package service

import (
	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/architectv/merchantx-task/pkg/repository"
)

type OfferService struct {
	repo repository.Offer
}

func NewOfferService(repo repository.Offer) *OfferService {
	return &OfferService{
		repo: repo,
	}
}

func (s *OfferService) Get(sellerId, offerId int, substr string) ([]model.Offer, error) {
	offers, err := s.repo.Get(sellerId, offerId, substr)
	return offers, err
}
