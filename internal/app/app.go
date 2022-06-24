package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mmuoDev/price-discount/internal/db"
)

//App has handlers for this app
type App struct {
	RetrieveProductsHandler http.HandlerFunc
}

//Handler returns the http handler for this app
func (a App) Handler() http.HandlerFunc {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/products", a.RetrieveProductsHandler)
	return http.HandlerFunc(router.ServeHTTP)
}

//OptionalArgs optional arguments for this application
type OptionalArgs struct {
	RetrieveProducts db.RetrieveProductsFunc
}

// Options is a type for application options to modify the app
type Options func(o *OptionalArgs)

//New creates a new instance of the App
func New(options ...Options) App {
	o := OptionalArgs{
		RetrieveProducts: db.RetrieveProducts(),
	}
	for _, option := range options {
		option(&o)
	}
	retrieveProducts := RetrieveProductsHandler(o.RetrieveProducts)
	return App{
		RetrieveProductsHandler: retrieveProducts,
	}
}
