package scrape

import (
	"net/url"
	"path"
)

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
