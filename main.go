package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	ID        string
	Title     string
	Condition string
	Price     string
	URL       string
}

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
	doc.Find("ul.srp-results li.s-item").Each(func(i int, s *goquery.Selection) {
		p, _ := ParseProductHTML(s)
		fmt.Printf("Product %d: %s\nCondition %s\nPrice %s\nURL: %s\nID: %s\n", i+1, p.Title, p.Condition, p.Price, p.URL, p.ID)
	})
}

func ParseProductHTML(s *goquery.Selection) (*Product, error) {
	// Extracting the title and href attribute from a link
	title := s.Find(".s-item__title").Text()
	price := s.Find(".s-item__price").Text()
	href, exists := s.Find(".s-item__link").Attr("href")
	if !exists {
		return nil, fmt.Errorf("failed to find product URL for %s", title)
	}
	condition := s.Find(".SECONDARY_INFO").Text()

	id, err := getLastSegment(href)
	if err != nil {
		return nil, fmt.Errorf("failed to parse product URL %s: %v", href, title)
	}
	return &Product{
		ID:        id,
		Title:     title,
		Condition: condition,
		Price:     price,
		URL:       href,
	}, nil
}

// getLastSegment returns the last segment of a URL's path, excluding any query parameters.
func getLastSegment(urlStr string) (string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err // Return the error if the URL cannot be parsed
	}

	// Get the path and trim the trailing slash if present
	trimmedPath := path.Clean(parsedURL.Path)

	// Get the last segment of the path
	lastSegment := path.Base(trimmedPath)

	return lastSegment, nil
}

func main() {
	Scrape()
}
