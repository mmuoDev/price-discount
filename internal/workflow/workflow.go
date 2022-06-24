package workflow

import "github.com/mmuoDev/price-discount/pkg"

//RetrieveProductsFunc returns functionality to retrieve products
type RetrieveProductsFunc func (qp pkg.QueryParams) ([]pkg.Product, pkg.PageInfo, error) 