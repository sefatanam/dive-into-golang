package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// Using middleware
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(30 * time.Second))

	// file server
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Using Routes
	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)
	return mux
}
