package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Product is the representaion of the EBAY product
type Product struct {
	ID        string `json:"-"`
	Title     string `json:"title"`
	Condition string `json:"condition"`
	Price     string `json:"price"`
	URL       string `json:"product_url"`
}

// ParseProductHTML extracts product details
func ParseProductHTML(e *colly.HTMLElement) (*Product, error) {
	titleDOM := e.DOM.Find("div.s-item__title span")
	titleDOM.Find("span").Remove()
	title := titleDOM.Text()

	price := e.ChildText(".s-item__price")
	href, exists := e.DOM.Find(".s-item__link").Attr("href")
	if !exists {
		return nil, fmt.Errorf("ERROR: failed to find product URL for %s", title)
	}
	id, err := getLastSegment(href)
	if err != nil {
		return nil, fmt.Errorf("ERROR: failed to parse product URL %s: %w", href, err)
	}
	condition := e.ChildText(".SECONDARY_INFO")
	return &Product{
		ID:        id,
		Title:     title,
		Condition: condition,
		Price:     price,
		URL:       href,
	}, nil
}
