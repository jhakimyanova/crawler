package scrape

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Scraper struct {
	URL       string
	PageSize  int
	UserAgent string
	Client    *http.Client
}

func (s Scraper) Scrape(parCtx context.Context) error {
	page := 1
	parsedURL, err := url.Parse(s.URL)
	if err != nil {
		return fmt.Errorf("failed to parse url %s: %w", s.URL, err)
	}
	queryParams := parsedURL.Query()
	out := make(chan *Product, 1000)
	for {
		queryParams.Set("_pgn", strconv.Itoa(page))
		queryParams.Set("_ipg", strconv.Itoa(s.PageSize))
		parsedURL.RawQuery = queryParams.Encode()
		log.Printf("Scrapping %s", parsedURL)
		ctx, cancel := context.WithTimeout(parCtx, 5*time.Second)
		defer cancel()
		count, err := s.scrapePage(ctx, parsedURL.String(), out)
		if err != nil {
			log.Printf("failed to scrape %d page: %v", page, err)
		}
		fmt.Println(count)
		if count == 0 {
			break
		}
		page++
	}
	return nil
}

func (s Scraper) scrapePage(ctx context.Context, pageURL string, out chan *Product) (int, error) {
	// Create the HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", pageURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create HTTP GET request: %w", err)
	}
	// Set a custom User-Agent and other necessary headers for crawling
	req.Header.Set("User-Agent", s.UserAgent)
	// Add other headers as needed

	// Send the request
	resp, err := s.Client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to free resources

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	return ParseDocument(resp.Body, out)
}
