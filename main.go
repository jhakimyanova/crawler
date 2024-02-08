package main

import (
	"log"

	"github.com/jhakimyanova/crawler/scrape"
)

func main() {
	err := scrape.Scrape("https://www.ebay.com/sch/garlandcomputer/m.html")
	if err != nil {
		log.Fatal(err)
	}
}
