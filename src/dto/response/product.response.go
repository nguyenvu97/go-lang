package response

import "time"

type ProductDto struct {
	ID            int32     `json:"id"`
	ProductURL    string    `json:"product_url"`
	ProductCategory string  `json:"product_category"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	Status        int32     `json:"status"`
	Quantity      int     `json:"quantity"`
	ProductName   string    `json:"product_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	}
