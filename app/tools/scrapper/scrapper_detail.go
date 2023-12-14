package scrapper

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/freekup/product-scrapper/app/entity"
)

func (s *Scrapper) GetProductDetail(productLink entity.ProductLink) (result entity.Product, errx error) {
	result = entity.Product{}

	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(s.userAgent),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(productLink.ExtractLink()),
		chromedp.Sleep(1000*time.Millisecond),

		chromedp.Text("h1.css-1os9jjn", &result.Name, chromedp.ByQuery),
		chromedp.Text("div[data-testid=lblPDPDescriptionProduk]", &result.Description, chromedp.ByQuery),
		chromedp.AttributeValue("img.css-1c345mg", "src", &result.ImageURL, nil, chromedp.ByQuery),
		chromedp.Text("div[data-testid=lblPDPDetailProductPrice].price", &result.Price, chromedp.ByQuery),
		chromedp.Text("span[data-testid=lblPDPDetailProductRatingNumber].main", &result.Rating, chromedp.ByQuery),
		chromedp.Text("h2.css-1wdzqxj-unf-heading.e1qvo2ff2", &result.StoreName, chromedp.ByQuery),
	)
	if err != nil {
		log.Println("Error while performing the automation logic:", err)
	}

	return
}
