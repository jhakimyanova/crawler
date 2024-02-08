package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jhakimyanova/crawler/scrape"
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

	err := scrape.Scrape(ctx, "https://www.ebay.com/sch/garlandcomputer/m.html")
	if err != nil {
		log.Fatal(err)
	}
}
