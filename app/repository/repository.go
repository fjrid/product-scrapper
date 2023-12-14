package repository

import "github.com/freekup/product-scrapper/app/entity"

type Repository interface {
	InsertProduct(products []entity.Product) (err error)
}
