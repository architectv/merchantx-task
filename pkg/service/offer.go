package service

import (
	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/architectv/merchantx-task/pkg/repository"
	"github.com/tealeg/xlsx"
)

const (
	OfferIdCol   = 0
	NameCol      = 1
	PriceCol     = 2
	QuantityCol  = 3
	AvailableCol = 4
)

type OfferService struct {
	repo repository.Offer
}

func NewOfferService(repo repository.Offer) *OfferService {
	return &OfferService{repo: repo}
}

func (s *OfferService) GetAllByParams(sellerId, offerId int, substr string) ([]*model.Offer, error) {
	offers, err := s.repo.GetAllByParams(sellerId, offerId, substr)
	return offers, err
}

func (s *OfferService) PutWithFile(sellerId int, filename string) (*model.Statistics, error) {
	stat := new(model.Statistics)
	if err := s.processXlsx(sellerId, filename, stat); err != nil {
		return nil, err
	}

	return stat, nil
}

func (s *OfferService) isExist(sellerId, offerId int) bool {
	_, err := s.repo.GetByTuple(sellerId, offerId)
	return err == nil
}

func (s *OfferService) process(v *model.Offer, stat *model.Statistics) error {
	var err error
	if s.isExist(v.SellerId, v.OfferId) {
		if v.Available {
			err = s.repo.Update(v)
			if err != nil {
				return err
			}
			stat.UpdateCount++
		} else {
			err = s.repo.Delete(v.SellerId, v.OfferId)
			if err != nil {
				return err
			}
			stat.DeleteCount++
		}
	} else if v.Available {
		err = s.repo.Create(v)
		if err != nil {
			return err
		}
		stat.CreateCount++
	} else {
		stat.ErrorCount++
	}

	return nil
}

func (s *OfferService) processXlsx(sellerId int, filename string, stat *model.Statistics) error {
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}

	for _, sh := range wb.Sheets {
		for j := 0; j < sh.MaxRow; j++ {
			offerId, err := sh.Cell(j, OfferIdCol).Int()
			if err != nil || offerId <= 0 {
				stat.ErrorCount++
				continue
			}

			name := sh.Cell(j, NameCol).String()
			if "" == name {
				stat.ErrorCount++
				continue
			}

			price, err := sh.Cell(j, PriceCol).Int()
			if err != nil || price <= 0 {
				stat.ErrorCount++
				continue
			}

			quantity, err := sh.Cell(j, QuantityCol).Int()
			if err != nil || quantity <= 0 {
				stat.ErrorCount++
				continue
			}

			available := sh.Cell(j, AvailableCol).Bool()

			offer := model.Offer{
				SellerId:  sellerId,
				OfferId:   offerId,
				Name:      name,
				Price:     price,
				Quantity:  quantity,
				Available: available,
			}

			if err := s.process(&offer, stat); err != nil {
				return err
			}
		}
	}

	return nil
}
