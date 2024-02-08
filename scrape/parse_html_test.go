package scrape

import (
	"os"
	"testing"
)

func TestParseProductHTML(t *testing.T) {
	f, err := os.Open("test_ebay.html")
	if err != nil {
		t.Fatalf("failed to open file with test data: %v", err)
	}
	defer f.Close()

	// Test ParseDocument with html response saved to a test file
	out := make(chan *Product, 1000)
	n, err := ParseDocument(f, out)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if n != 60 {
		t.Errorf("Expected %d Product items, got '%d'", 60, n)
	}
}
