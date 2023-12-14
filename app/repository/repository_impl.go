package repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/freekup/product-scrapper/app/entity"
)

type RepoImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &RepoImpl{
		db: db,
	}
}

func (r *RepoImpl) InsertProduct(products []entity.Product) (err error) {
	query := sq.Insert("products").
		Columns("name", "description", "image_url", "price", "rating", "store_name").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, product := range products {
		query = query.Values(product.Name, product.Description, product.ImageURL, product.FloatPrice(), product.Rating, product.StoreName)
	}

	_, err = query.Exec()
	if err != nil {
		return
	}

	return
}
