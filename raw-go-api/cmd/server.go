// Package cmd serve the Backend.
package cmd

import (
	"fmt"
	"net/http"

	"github.com/sefatanam/rga/handlers"
	"github.com/sefatanam/rga/util"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	globalRouter := util.GlobalRouter(mux)

	fmt.Println("Backend running and listening on port 8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Something went wrong!", err)
	}
}
