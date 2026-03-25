package cmd

import (
	"net/http"

	"github.com/sefatanam/rga/handlers"
	"github.com/sefatanam/rga/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middleware.After))
}
