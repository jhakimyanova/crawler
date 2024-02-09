package scrape

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"sync/atomic"
)

// SaveProductsData reads product data from the out channel and writes it to the result files
func SaveProductsData(dir string, out chan *Product) {
	var wg sync.WaitGroup
	var count atomic.Int32
	for p := range out {
		p := p
		wg.Add(1)
		go func() {
			defer wg.Done()
			CreateProductFile(dir, p)
			count.Add(1)
		}()
	}
	wg.Wait()
	log.Printf("DEBUG: Written %d files", count.Load())
}

// CreateProductFile reads product data from the out channel and writes it to the result file in the specied directory
func CreateProductFile(dir string, p *Product) error {
	fileName := path.Join(dir, fmt.Sprintf("%s.json", p.ID))
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create product file: %w", err)
	}
	defer f.Close()
	jsonData, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return fmt.Errorf("failed marshal JSON for %v product: %w", p, err)
	}
	log.Printf("DEBUG: writing file %s", fileName)
	_, err = f.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write product data to %s file: %w", fileName, err)
	}

	return nil
}
