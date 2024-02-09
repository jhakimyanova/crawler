package scrape

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestParseProductHTML(t *testing.T) {
	// Read mock data from file
	mockData, err := os.ReadFile("test_ebay.html")
	if err != nil {
		t.Fatalf("scrapeTitle() returned an error: %v", err)
	}
	// Setup a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(mockData)
	}))
	defer ts.Close()

	// Test scrapeTitle with the URL of the test server
	fmt.Println(ts.URL)
	s := Scraper{URL: ts.URL, AllowedDomain: "127.0.0.1"}
	out := s.ScrapeProducts(ConditionAny)
	count := 0
	for range out {
		count++
	}
	if count != 25 {
		t.Errorf("Expected %d Product items, got '%d'", 25, count)
	}
}
