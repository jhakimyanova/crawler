package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Scrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.ebay.com/sch/garlandcomputer/m.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find and iterate through each product
	doc.Find(".s-item").Each(func(i int, s *goquery.Selection) {
		// Extracting the title and href attribute from a link
		title := s.Find(".s-item__title").Text()
		href, _ := s.Find(".s-item__link").Attr("href")

		fmt.Printf("Product %d: %s\nURL: %s\n", i+1, title, href)
	})
}

func main() {
	Scrape()
}
