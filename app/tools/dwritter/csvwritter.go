package dwritter

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/freekup/product-scrapper/app/entity"
)

type CSVWritter struct {
}

func NewCSVWritter() DWHandlerInterface {
	return &CSVWritter{}
}

func (w *CSVWritter) Store(products []entity.Product) error {
	csvFile, err := os.Create("products.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Write([]string{"Name of Product", "Description", "Image Link", "Price", "Rating", "Store Name"})

	for _, product := range products {
		csvwriter.Write([]string{product.Name, product.Description, product.ImageURL, product.Price, product.Rating, product.StoreName})
	}

	csvwriter.Flush()
	csvFile.Close()

	return nil
}
