package main

import (
	"flag"
	"log"
	"time"

	"github.com/jhakimyanova/crawler/scrape"
)

const (
	AllowedDomain   = "www.ebay.com"
	URL             = "https://www.ebay.com/sch/i.html?_ssn=garlandcomputer"
	ProductFilesDir = "data"
)

func main() {
	var condition scrape.Condition
	flag.Var(&condition, "condition", "Specifies the condition of the product.\nPossible values: any, new, used, unknown. Default: any")
	flag.Parse()

	startTime := time.Now()
	s := scrape.Scraper{URL: URL, AllowedDomain: AllowedDomain}
	out := s.ScrapeProducts(condition)
	scrape.SaveProductsData(ProductFilesDir, out)
	log.Printf("DEBUG: Elapsed scraping time: %s", time.Since(startTime))
}
