package app

import (
	"github.com/freekup/product-scrapper/app/repository"
)

var App = struct {
	Repository repository.Repository
}{}
