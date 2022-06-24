package internal

//Product represents a product entity
type Product struct {
	Name     string `bson:"name"`
	SKU      string `bson:"sku"`
	Category string `bson:"category"`
	Price    int64  `bson:"price"`
}

//Discount represents disount info
type Discount struct {
	NewAmount            int64
	DiscountedPercentage string
}
