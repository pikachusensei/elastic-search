package product

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
	InStock     bool    `json:"in_stock"`
}
