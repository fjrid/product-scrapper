package dwritter

import (
	"github.com/freekup/product-scrapper/app/entity"
	"github.com/freekup/product-scrapper/app/repository"
)

type DBWritter struct {
	repository repository.Repository
}

func NewDBWritter(repository repository.Repository) DWHandlerInterface {
	return &DBWritter{
		repository: repository,
	}
}

func (w *DBWritter) Store(data []entity.Product) error {
	return w.repository.InsertProduct(data)
}
