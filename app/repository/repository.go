package repository

import "github.com/freekup/product-scrapper/app/entity"

type Repository interface {
	CreateTable() (err error)
	Droptable() (err error)

	InsertProduct(products []entity.Product) (err error)
}
