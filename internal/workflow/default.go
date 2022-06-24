package workflow

import (
	"log"
	"strings"

	"github.com/mmuoDev/price-discount/internal"
	"github.com/mmuoDev/price-discount/internal/db"
	"github.com/mmuoDev/price-discount/internal/mapping"
	"github.com/mmuoDev/price-discount/pkg"
	"github.com/pkg/errors"
)

//RetrieveProducts retrieves products
func RetrieveProducts(retrieveProducts db.RetrieveProductsFunc) RetrieveProductsFunc {
	return func(qp pkg.QueryParams) ([]pkg.Product, pkg.PageInfo, error) {
		pp := []pkg.Product{}
		log.Printf("Receiving products from DB with params=%+v", qp)
		products, pageInfo, err := retrieveProducts(qp)
		if err != nil {
			return pp, pkg.PageInfo{}, errors.Wrapf(err, "workflow - unable to retrieve products from DB")
		}
		for _, product := range products {
			d := computeDiscount(product)
			p := mapping.ToProducts(product, d)
			pp = append(pp, p)
		}
		return pp, pageInfo, nil
	}
}

//computeDiscount computes discount on a product
func computeDiscount(product internal.Product) internal.Discount {
	newAmount := product.Price
	if strings.ToLower(product.Category) == "boots" {
		discount := (30 * product.Price) / 100
		newAmount = product.Price - discount
		return internal.Discount{
			NewAmount:            newAmount,
			DiscountedPercentage: "30",
		}
	}
	if product.SKU == "000003" {
		discount := (15 * product.Price) / 100
		newAmount = product.Price - discount
		return internal.Discount{
			NewAmount:            newAmount,
			DiscountedPercentage: "15",
		}
	}
	return internal.Discount{
		NewAmount:            newAmount,
		DiscountedPercentage: "",
	}
}
