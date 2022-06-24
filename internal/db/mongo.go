package db

import (
	"github.com/mmuoDev/price-discount/internal"
	"github.com/mmuoDev/price-discount/pkg"
)

//RetrieveProducts retri
func RetrieveProducts() RetrieveProductsFunc {
	return func(qp pkg.QueryParams) ([]internal.Product, pkg.PageInfo, error) {
		//Implementation for mongo is specified here!
		//1. The products on mongo are filtered using the query params (qp)
		//2. Pagination is deduced
		//2. Products are returned with pagination info
		return []internal.Product{}, pkg.PageInfo{}, nil 
	}
}
