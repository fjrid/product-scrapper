package entity

import (
	"net/url"
	"strconv"
	"strings"
)

// Used to extract link from URL that used for Tokopedia's promo link
func (val ProductLink) ExtractLink() string {
	if strings.Contains(val.Link, "https://ta.tokopedia.com") {
		u, _ := url.Parse(val.Link)
		return u.Query().Get("r")
	}

	return val.Link
}

func (val Product) FloatPrice() float64 {
	parsedString := strings.ReplaceAll(strings.ReplaceAll(val.Price, ".", ""), "Rp", "")
	intPrice, _ := strconv.Atoi(parsedString)

	return float64(intPrice)
}
