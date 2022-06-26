package mapping

import (
	"github.com/mmuoDev/price-discount/internal"
	"github.com/mmuoDev/price-discount/pkg"
)

const (
	CURRENCY = "EUR"
)

//ToProducts maps from DB product response to pkg Product
func ToProducts(p internal.Product, d internal.Discount) pkg.Product {
	percentDiscount := ""
	if d.DiscountedPercentage != "" {
		percentDiscount = d.DiscountedPercentage + "%"
	}
	return pkg.Product{
		Name:     p.Name,
		SKU:      p.SKU,
		Category: p.Category,
		Price: pkg.Price{
			Currency:           CURRENCY,
			Original:           p.Price/100,
			Final:              d.NewAmount/100,
			DiscountPercentage: percentDiscount,
		},
	}
}
