package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jhakimyanova/crawler/scrape"
)

const (
	PAGE_SIZE  = 60
	URL        = "https://www.ebay.com/sch/i.html?_ssn=garlandcomputer"
	USER_AGENT = "CustomCrawler/1.0"
)

func main() {
	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())

	// Listen for interrupt signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Block until a signal is received
		sig := <-sigs
		fmt.Println("Received signal:", sig)

		// Cancel the context
		cancel()
	}()

	scrapper := scrape.Scraper{
		URL:       URL,
		PageSize:  PAGE_SIZE,
		UserAgent: USER_AGENT,
		Client:    createHTTPClient(),
	}
	err := scrapper.Scrape(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func createHTTPClient() *http.Client {
	// Create a custom HTTP transport for the client
	transport := &http.Transport{
		// DialContext controls the dialer used to create TCP connections.
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,  // Timeout for establishing a new connection
			KeepAlive: 30 * time.Second, // Keep-alive period for an active network connection
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second, // Timeout for TLS handshake
		MaxIdleConns:        100,              // Max idle connections to keep per-host
		IdleConnTimeout:     90 * time.Second, // Max amount of time an idle connection will remain idle before closing
		MaxConnsPerHost:     10,               // Max number of connections per host
	}

	// Create an HTTP client with the transport
	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second, // Total timeout for requests, including Dial / DialContext, TLS handshake, and reading the response body
	}
}
