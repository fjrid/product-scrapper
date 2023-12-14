package dwritter

import (
	"github.com/freekup/product-scrapper/app/entity"
)

type DWHandlerInterface interface {
	Store(data []entity.Product) error
}
