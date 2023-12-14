package scrapper

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"github.com/freekup/product-scrapper/app/entity"
)

func (s *Scrapper) GetProductList() (results []entity.ProductLink, err error) {
	var (
		currentPage = 0
	)

	for len(results) < s.maxData {
		currentPage++

		var nodes []*cdp.Node
		nodes, err = s.fetchProductList(currentPage)
		if err != nil {
			return
		}

		wg := sync.WaitGroup{}

		for _, node := range nodes {
			wg.Add(1)

			go func(node *cdp.Node) {
				defer wg.Done()

				product, err := s.extractProductList(node)
				if err != nil {
					log.Fatalf("Product link is empty: %+v", err)
				}

				results = append(results, product)
			}(node)
		}

		wg.Wait()
	}

	results = results[:s.maxData]

	return
}

func (s *Scrapper) fetchProductList(page int) (nodes []*cdp.Node, err error) {
	link := fmt.Sprintf("https://www.tokopedia.com/p/handphone-tablet/handphone?page=%v", page)
	nodes = make([]*cdp.Node, 0)

	log.Println("Scapping List:", link)

	err = chromedp.Run(s.ctxChromedp,
		chromedp.Navigate(link),

		chromedp.Sleep(1000*time.Millisecond),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1000*time.Millisecond),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(1000*time.Millisecond),

		chromedp.Nodes("div.css-bk6tzz.e1nlzfl2", &nodes, chromedp.ByQueryAll),
	)
	if err != nil {
		return
	}

	return
}

func (s *Scrapper) extractProductList(node *cdp.Node) (product entity.ProductLink, err error) {
	chromedp.Run(
		s.ctxChromedp,
		chromedp.AttributeValue("a.css-54k5sq", "href", &product.Link, nil, chromedp.ByQuery, chromedp.FromNode(node)),
	)

	if product.Link == "" {
		err = errors.New("product link is empty")
		return
	}

	return
}
