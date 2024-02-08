package scrape

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

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
	_, err = f.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write product data to %s file: %w", fileName, err)
	}
	return nil
}
