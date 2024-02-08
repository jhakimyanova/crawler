package scrape

type Product struct {
	ID        string `json:"-"`
	Title     string `json:"title"`
	Condition string `json:"condition"`
	Price     string `json:"price"`
	URL       string `json:"product_url"`
}
