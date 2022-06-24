package app

import (
	"net/http"

	"github.com/mmuoDev/price-discount/internal/db"
	"github.com/mmuoDev/price-discount/internal/workflow"
	"github.com/mmuoDev/price-discount/pkg"
)

//ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

//RetrieveProductsHandler handles a http request to retrieve products
func RetrieveProductsHandler(retrieveDBProducts db.RetrieveProductsFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var params pkg.QueryParams
		if err := GetQueryParams(&params, r); err != nil {
			res := ErrorResponse{Error: err.Error()}
			ServeJSON(res, w, http.StatusInternalServerError)
			return
		}
		retrieve := workflow.RetrieveProducts(retrieveDBProducts)
		products, pageInfo, err := retrieve(params)
		if err != nil {
			res := ErrorResponse{Error: err.Error()}
			ServeJSON(res, w, http.StatusInternalServerError)
			return
		}
		res := pkg.ProductsResponse{
			Products: products,
			PageInfo: pageInfo,
		}
		ServeJSON(res, w, http.StatusOK)
	}
}
