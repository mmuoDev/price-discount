package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mmuoDev/price-discount/internal/app"
)

func main() {

	a := app.New()
	log.Println(fmt.Sprintf("Starting server on port:%s", os.Getenv("APP_PORT")))
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), a.Handler()))
}
