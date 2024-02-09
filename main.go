package main

import (
	"log"
	"time"

	"github.com/jhakimyanova/crawler/scrape"
)

const (
	ALLOWED_DOMAIN    = "www.ebay.com"
	URL               = "https://www.ebay.com/sch/i.html?_ssn=garlandcomputer"
	PRODUCT_FILES_DIR = "data"
)

func main() {
	startTime := time.Now()
	s := scrape.Scraper{URL: URL, AllowedDomain: ALLOWED_DOMAIN}
	out := s.ScrapeProducts()
	scrape.SaveProductsData(PRODUCT_FILES_DIR, out)
	log.Printf("DEBUG: Elapsed scraping time: %s", time.Since(startTime))
}
