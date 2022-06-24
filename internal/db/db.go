package db

import (
	"github.com/mmuoDev/price-discount/internal"
	"github.com/mmuoDev/price-discount/pkg"
)

const (
	PULL_LIMIT = 5 //default number of pages if none is specified. This can also be specified on OPEN API SPEC
)

//RetrieveProductsFunc returns functionality to retrieve products
type RetrieveProductsFunc func(qp pkg.QueryParams) ([]internal.Product, pkg.PageInfo, error)
