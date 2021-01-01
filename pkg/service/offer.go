package service

import (
	"log"

	"github.com/architectv/merchantx-task/pkg/model"
	"github.com/architectv/merchantx-task/pkg/repository"
	"github.com/tealeg/xlsx"
)

type OfferService struct {
	repo repository.Offer
}

func NewOfferService(repo repository.Offer) *OfferService {
	return &OfferService{
		repo: repo,
	}
}

// TODO: return slice of pointers
func (s *OfferService) Get(sellerId, offerId int, substr string) ([]model.Offer, error) {
	offers, err := s.repo.Get(sellerId, offerId, substr)
	return offers, err
}

const (
	OfferIdCol   = 0
	NameCol      = 1
	PriceCol     = 2
	QuantityCol  = 3
	AvailableCol = 4
)

func parseXlsx(sellerId int, filename string, stat *model.Statistics) ([]model.Offer, error) {
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	var offers []model.Offer

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
			log.Println(offer)

			offers = append(offers, offer)
		}
	}

	return offers, nil
}

func (s *OfferService) Put(sellerId int, filename string) (*model.Statistics, error) {
	stat := new(model.Statistics)
	offers, err := parseXlsx(sellerId, filename, stat)
	if err != nil {
		return nil, err
	}
	log.Println(len(offers))
	for _, v := range offers {
		s.process(&v, stat)
	}

	return stat, nil
}

func (s *OfferService) process(v *model.Offer, stat *model.Statistics) {
	var err error
	if s.isExist(v.SellerId, v.OfferId) {
		if v.Available {
			log.Println("UPDATE")
			err = s.repo.Update(v)
			if err != nil {
				log.Println(err.Error())
			}
			stat.UpdateCount++
		} else {
			log.Println("DELETE")
			err = s.repo.Delete(v.SellerId, v.OfferId)
			if err != nil {
				log.Println(err.Error())
			}
			stat.DeleteCount++
		}
	} else if v.Available {
		log.Println("CREATE")
		err = s.repo.Create(v)
		if err != nil {
			log.Println(err.Error())
		}
		stat.CreateCount++
	} else {
		stat.ErrorCount++
	}
}

func (s *OfferService) isExist(sellerId, offerId int) bool {
	_, err := s.repo.GetByTuple(sellerId, offerId)
	return err == nil
}
