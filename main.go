package main

import (
	"github.com/jhakimyanova/crawler/scrape"
)

const (
	ALLOWED_DOMAIN    = "www.ebay.com"
	URL               = "https://www.ebay.com/sch/i.html?_ssn=garlandcomputer"
	PRODUCT_FILES_DIR = "data"
)

func main() {
	s := scrape.Scraper{URL: URL, AllowedDomain: ALLOWED_DOMAIN}
	out := s.ScrapeProducts()
	scrape.SaveProductsData(PRODUCT_FILES_DIR, out)
}
