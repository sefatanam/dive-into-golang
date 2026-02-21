// Package cmd serve the Backend.
package cmd

import (
	"fmt"
	"net/http"

	"github.com/sefatanam/rga/middleware"
	"github.com/sefatanam/rga/util"
)

func Serve() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()
	manager.Use(middleware.Emoji, middleware.Logger)

	initRoutes(mux, manager)
	globalRouter := util.GlobalRouter(mux)

	fmt.Println("Backend running and listening on port 8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Something went wrong!", err)
	}
}
