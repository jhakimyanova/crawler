package scrape

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type Scraper struct {
	AllowedDomain string
	URL           string
}

// ScrapeProducts uses Colly to scrape products from a webpage
func (s Scraper) ScrapeProducts() chan *Product {
	out := make(chan *Product, 1000)

	go func() {
		defer close(out)
		// Create a collector
		c := colly.NewCollector(
		// colly.AllowedDomains(s.AllowedDomain),
		)
		// Random user agent to bypass anti-crawler protection
		extensions.RandomUserAgent(c)

		// Extract data from each page
		c.OnHTML("ul.srp-results li.s-item", ParseProductHTML(out))

		// Handle pagination
		c.OnHTML("a.pagination__next", func(e *colly.HTMLElement) {
			nextPage := e.Attr("href")
			if nextPage != "" {
				c.Visit(e.Request.AbsoluteURL(nextPage))
			}
		})

		// Handle errors
		c.OnError(func(r *colly.Response, err error) {
			log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})

		c.OnRequest(func(r *colly.Request) {
			log.Println("Visiting", r.URL)
		})

		// Start scraping
		err := c.Visit(s.URL)
		if err != nil {
			log.Printf("ERROR: failed visiting %s: %v", s.URL, err)
		}
	}()
	return out
}
