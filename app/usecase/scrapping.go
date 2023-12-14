package usecase

import (
	"fmt"
	"log"
	"sync"

	"github.com/freekup/product-scrapper/app/entity"
	"github.com/freekup/product-scrapper/app/tools/dwritter"
	"github.com/freekup/product-scrapper/app/tools/scrapper"
)

var defaultMaxThread = 10

type ScrappingUsecase struct {
	saveTo    []string
	ds        scrapper.Scrapper
	dw        dwritter.DataWritter
	maxThread int
}

func NewScrappingUsecase(saveTo []string, dataScrapper scrapper.Scrapper, dataWritter dwritter.DataWritter) *ScrappingUsecase {
	su := ScrappingUsecase{
		saveTo:    saveTo,
		ds:        dataScrapper,
		dw:        dataWritter,
		maxThread: defaultMaxThread,
	}

	return &su
}

func (u *ScrappingUsecase) ProcessScrap() {
	var productLinks []entity.ProductLink

	productLinks, err := u.ds.GetProductList()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Total data:", len(productLinks), "data")

	products := u.getDataProducts(productLinks)

	log.Println("Total scrapped data:", len(products), "data")

	u.dw.Write(u.saveTo, products)

	fmt.Println("Finish storing", len(products), "products")
}

func (u *ScrappingUsecase) getDataProducts(productLinks []entity.ProductLink) (products []entity.Product) {
	var (
		wg sync.WaitGroup

		// Use to handle maximum thread
		ch = make(chan bool, u.maxThread)
	)

	defer close(ch)

	products = make([]entity.Product, 0)

	for i, productLink := range productLinks {
		wg.Add(1)
		ch <- true

		go func(idx int, productLink entity.ProductLink) {
			log.Println(idx, "Fetching", productLink.ExtractLink())

			defer wg.Done()

			product, err := u.ds.GetProductDetail(productLink)
			if err != nil {
				log.Fatal(err)
			}

			products = append(products, product)

			<-ch
		}(i, productLink)
	}

	wg.Wait()

	return
}
