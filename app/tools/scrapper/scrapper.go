package scrapper

import (
	"context"

	"github.com/chromedp/chromedp"
)

var (
	defaultMaxData   = 100
	defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"
)

type Scrapper struct {
	maxData   int
	userAgent string

	ctxChromedp context.Context
}

func CreateScrapper(maxData int) Scrapper {
	if maxData == 0 {
		maxData = defaultMaxData
	}

	s := Scrapper{
		maxData:   maxData,
		userAgent: defaultUserAgent,
	}

	s.InitChromedpCtx()

	return s
}

func (s *Scrapper) InitChromedpCtx() {
	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(s.userAgent),
	)

	ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)

	s.ctxChromedp, _ = chromedp.NewContext(ctx)
}
