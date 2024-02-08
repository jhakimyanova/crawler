package scrape

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const PAGE_SIZE = 60

func Scrape(parCtx context.Context, urlStr string) error {
	page := 1
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("failed to parse url %s: %w", urlStr, err)
	}
	queryParams := parsedURL.Query()
	for {
		queryParams.Set("_pgn", strconv.Itoa(page))
		parsedURL.RawQuery = queryParams.Encode()
		log.Printf("Scrapping %s", parsedURL)
		ctx, cancel := context.WithTimeout(parCtx, 5*time.Second)
		defer cancel()
		count, err := ScrapePage(ctx, parsedURL.String())
		if err != nil {
			log.Printf("failed to scrape %d page: %v", page, err)
		}
		fmt.Println(count)
		if count < PAGE_SIZE {
			break
		}
		page++
	}
	return nil
}

func ScrapePage(ctx context.Context, url string) (int, error) {
	// Create the HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create HTTP GET request: %w", err)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to free resources

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to create document from HTTP response body: %w", err)
	}

	// Find and iterate through each product
	items := doc.Find("ul.srp-results li.s-item").Each(func(i int, s *goquery.Selection) {
		p, _ := ParseProductHTML(s)
		err := CreateProductFile("data", p)
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
