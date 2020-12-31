package repository

type Offer interface {
}

type Repository struct {
	Offer
}

func NewRepository() *Repository {
	return &Repository{}
}
