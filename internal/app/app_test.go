package app_test

import (
	"fmt"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/mmuoDev/price-discount/internal"
	"github.com/mmuoDev/price-discount/internal/app"
	"github.com/mmuoDev/price-discount/pkg"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveProductsWorksAsExpected(t *testing.T) {
	retrieveFromDBIsInvoked := false

	mockProductsRetrieval := func(o *app.OptionalArgs) {
		o.RetrieveProducts = func(qp pkg.QueryParams) ([]internal.Product, pkg.PageInfo, error) {
			retrieveFromDBIsInvoked = true
			var dbRes []internal.Product
			app.FileToStruct(filepath.Join("testdata", "retrieve_products_db_response.json"), &dbRes)
			return dbRes, pkg.PageInfo{
				CurrentPage: 1,
				LastPage:    1,
			}, nil
		}
	}

	opts := []app.Options{
		mockProductsRetrieval,
	}

	ap := app.New(opts...)
	serverURL, cleanUpServer := app.NewTestServer(ap.Handler())
	defer cleanUpServer()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/products", serverURL), nil)

	client := &http.Client{}
	res, _ := client.Do(req)

	t.Run("Http Status Code is 200", func(t *testing.T) {
		assert.Equal(t, res.StatusCode, http.StatusOK)
	})
	t.Run("DB is invoked", func(t *testing.T) {
		assert.True(t, retrieveFromDBIsInvoked)
	})
	t.Run("Response is as expected", func(t *testing.T) {
		app.AssertResBodyEqual(t, filepath.Join("testdata", "retrieve_products_response.json"), res)
	})
}
