package scrape

import (
	"log"
	"net/url"
	"path"

	"github.com/gocolly/colly"
)

// ParseProductHTML extracts product details
func ParseProductHTML(out chan *Product) func(e *colly.HTMLElement) {

	return func(e *colly.HTMLElement) {
		titleDOM := e.DOM.Find("div.s-item__title span")
		titleDOM.Find("span").Remove()
		title := titleDOM.Text()

		price := e.ChildText(".s-item__price")
		href, exists := e.DOM.Find(".s-item__link").Attr("href")
		if !exists {
			log.Printf("ERROR: failed to find product URL for %s", title)
			return
		}
		id, err := getLastSegment(href)
		if err != nil {
			log.Printf("ERROR: failed to parse product URL %s: %v", href, err)
			return
		}
		condition := e.ChildText(".SECONDARY_INFO")
		p := &Product{
			ID:        id,
			Title:     title,
			Condition: condition,
			Price:     price,
			URL:       href,
		}
		out <- p
	}
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
