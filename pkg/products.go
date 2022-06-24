package pkg

//QueryParams represents parameters needed to filter products
type QueryParams struct {
	Category      string `json:"category"`
	PriceLessThan string `json:"priceLessThan"`
	Limit         int64  `json:"limit"`
	Page          int64  `json:"page"`
}

//Price represents a price
type Price struct {
	Original           int64  `json:"original"`
	Final              int64  `json:"final"`
	DiscountPercentage string `json:"discount_percentage"`
	Currency           string `json:"currency"`
}

//Product represents a product
type Product struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price  `json:"price"`
	SKU      string `json:"sku"`
}

// PageInfo returns pagination info
type PageInfo struct {
	CurrentPage int64 `json:"currentPage"`
	LastPage    int64 `json:"lastPage"`
}

//ProductsResponse represents an API retrieve products response
type ProductsResponse struct {
	Products []Product `json:"products"`
	PageInfo PageInfo  `json:"pageInfo"`
}
