package scrape

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"path"

	"github.com/PuerkitoBio/goquery"
)

func ParseDocument(r io.Reader, out chan *Product) (int, error) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return 0, fmt.Errorf("failed to create document from HTTP response body: %w", err)
	}

	// Find and iterate through each product
	items := doc.Find("ul.srp-results li.s-item").Each(func(i int, s *goquery.Selection) {
		p, _ := ParseProductHTML(s)
		out <- p
		// err := CreateProductFile("data", p)
		if err != nil {
			log.Printf("failed to create product file: %s", err)
		}
	})
	return len(items.Nodes), nil
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
		return nil, fmt.Errorf("failed to parse product URL %s: %w", href, err)
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
