package service

import "github.com/architectv/merchantx-task/pkg/repository"

type Offer interface {
}

type Service struct {
	Offer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
