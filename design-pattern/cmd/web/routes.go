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

	// display test pages

	mux.Get("/test-patterns", app.TestPatterns)

	// factory routes
	mux.Get("/api/dog-from-factory", app.CreateDogFromFactory)
	mux.Get("/api/cat-from-factory", app.CreateCatFromFactory)
	// abstract factory routes
	mux.Get("/api/dog-from-abstract-factory", app.CreateDogFromAbstractFactory)
	mux.Get("/api/cat-from-abstract-factory", app.CreateCatFromAbstractFactory)

	// Using Routes
	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)

	// builder routes
	mux.Get("/api/dog-from-builder", app.CreateDogWithBuilder)
	mux.Get("/api/cat-from-builder", app.CreateCatWithBuilder)

	// apis
	mux.Get("/api/dog-breeds", app.GetAllDogBreedsJSON)

	return mux
}
