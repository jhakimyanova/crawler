package scrape

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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
	s := Scraper{URL: ts.URL, AllowedDomain: "127.0.0.1"}
	out := s.ScrapeProducts(ConditionAny)

	// Define the directory path to create
	dirPath := filepath.Join(os.TempDir(), "testdir")

	// Create the temporary directory for storing files created by this test
	err = os.Mkdir(dirPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory: %s", err)
	}

	// Defer the removal of the temporary directory
	defer os.RemoveAll(dirPath)

	SaveProductsData(dirPath, out)

	// Read directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		t.Fatalf("Failed to read the temporary test directory: %s", err)
	}

	// Count files (excluding directories)
	fileCount := 0
	for _, entry := range entries {
		if !entry.IsDir() { // Check if the entry is not a directory
			fileCount++
		}
	}
	if fileCount != 25 {
		t.Errorf("Expected %d product files, got '%d'", 25, fileCount)
	}
}
